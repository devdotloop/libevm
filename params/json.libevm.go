// Copyright 2024 the libevm authors.
//
// The libevm additions to go-ethereum are free software: you can redistribute
// them and/or modify them under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The libevm additions are distributed in the hope that they will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see
// <http://www.gnu.org/licenses/>.

package params

import (
	"encoding/json"
	"fmt"
)

var _ interface {
	json.Marshaler
	json.Unmarshaler
} = (*ChainConfig)(nil)

// UnmarshalJSON implements the [json.Unmarshaler] interface.
func (c *ChainConfig) UnmarshalJSON(data []byte) error {
	r := registeredExtras
	if !r.Registered() {
		return unmarshalChainConfigJSONDirectly(data, c)
	}

	extra := r.Get().newChainConfig()
	if err := UnmarshalChainConfigJSON(data, c, extra, r.Get().reuseJSONRoot); err != nil {
		return err
	}
	c.extra = extra
	return nil
}

// UnmarshalChainConfigJSON unmarshals JSON into a coupled ChainConfig and
// arbitrary extra payload, as if the payload's type had been registered as
// [Extras]. This is equivalent to regular JSON handling but without the need
// for explicit registration with [RegisterExtras]; e.g. when unmarshalling
// configs from chains that use different payload types.
//
// `extra` MUST NOT be nil. Such circumstances are equivalent to regular JSON
// handling, in which case [ChainConfig.UnmarshalJSON] or [json.Unmarshal]
// should be used directly.
func UnmarshalChainConfigJSON[C any](data []byte, c *ChainConfig, extra *C, reuseJSONRoot bool) error {
	if !reuseJSONRoot {
		return c.unmarshalJSONWithExtra(data, extra)
	}
	if err := json.Unmarshal(data, extra); err != nil {
		return fmt.Errorf("unmarshal JSON into %T payload %T: %v", c, extra, err)
	}
	return unmarshalChainConfigJSONDirectly(data, c)
}

// chainConfigWithoutMethods avoids infinite recursion into
// [ChainConfig.UnmarshalJSON].
type chainConfigWithoutMethods ChainConfig

func unmarshalChainConfigJSONDirectly(data []byte, c *ChainConfig) error {
	if err := json.Unmarshal(data, (*chainConfigWithoutMethods)(c)); err != nil {
		return fmt.Errorf("unmarshal JSON into %T: %v", c, err)
	}
	return nil
}

// chainConfigWithExportedExtra supports JSON (un)marshalling of a [ChainConfig]
// while exposing the `extra` field as the "extra" JSON key.
type chainConfigWithExportedExtra struct {
	*chainConfigWithoutMethods     // embedded to achieve regular JSON unmarshalling
	Extra                      any `json:"extra"` // `c.extra` is otherwise unexported
}

// unmarshalJSONWithExtra unmarshals JSON under the assumption that the [Extras]
// payload is in the JSON "extra" key. All other unmarshalling is performed as
// if no [Extras] were present.
func (c *ChainConfig) unmarshalJSONWithExtra(data []byte, extra any) error {
	cc := &chainConfigWithExportedExtra{
		chainConfigWithoutMethods: (*chainConfigWithoutMethods)(c),
		Extra:                     extra,
	}
	return json.Unmarshal(data, cc)
}

// MarshalJSON implements the [json.Marshaler] interface.
func (c *ChainConfig) MarshalJSON() ([]byte, error) {
	switch reg := registeredExtras; {
	case !reg.Registered():
		return json.Marshal((*chainConfigWithoutMethods)(c))

	case !reg.Get().reuseJSONRoot:
		return c.marshalJSONWithExtra()

	default: // reg.reuseJSONRoot == true
		// The inverse of reusing the JSON root is merging two JSON buffers,
		// which isn't supported by the native package. So we use
		// map[string]json.RawMessage intermediates.
		geth, err := toJSONRawMessages((*chainConfigWithoutMethods)(c))
		if err != nil {
			return nil, err
		}
		extra, err := toJSONRawMessages(c.extra)
		if err != nil {
			return nil, err
		}

		for k, v := range extra {
			if _, ok := geth[k]; ok {
				return nil, fmt.Errorf("duplicate JSON key %q in both %T and registered extra", k, c)
			}
			geth[k] = v
		}
		return json.Marshal(geth)
	}
}

// marshalJSONWithExtra is the inverse of unmarshalJSONWithExtra().
func (c *ChainConfig) marshalJSONWithExtra() ([]byte, error) {
	cc := &chainConfigWithExportedExtra{
		chainConfigWithoutMethods: (*chainConfigWithoutMethods)(c),
		Extra:                     c.extra,
	}
	return json.Marshal(cc)
}

func toJSONRawMessages(v any) (map[string]json.RawMessage, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	msgs := make(map[string]json.RawMessage)
	if err := json.Unmarshal(buf, &msgs); err != nil {
		return nil, err
	}
	return msgs, nil
}
