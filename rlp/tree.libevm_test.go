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

package rlp

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTreeRoundTripOnUpstream(t *testing.T) {
	for _, tt := range encTests {
		t.Run("", func(t *testing.T) {
			if tt.error != "" {
				return
			}
			switch tt.val.(type) {
			// These return arbitrary byte slices, not real RLP.
			case *testEncoder:
				return
			case struct{ E testEncoderValueMethod }:
				return
			case undecodableEncoder:
				return
			}

			buf, err := hex.DecodeString(tt.output)
			require.NoErrorf(t, err, "hex.DecodeString(%T.output)", tt)

			node, err := ParseTree(buf)
			require.NoErrorf(t, err, "ParseTree(%T => %#x)", tt.val, buf)

			var got bytes.Buffer
			require.NoError(t, node.EncodeRLP(&got))
			assert.Equal(t, buf, got.Bytes())
		})
	}
}

func TestTreeRoundTripFuzzed(t *testing.T) {
	t.Parallel()

	type test struct {
		fuzzed   bool
		fuzzSeed int64
		value    any
		node     ItemNode
	}
	tests := []test{{
		value: []byte(nil),
		node:  StringNode{},
	}}

	for i := 0; i < 3e3; i++ {
		seed := int64(i)
		rng := rand.New(rand.NewSource(seed)) //nolint:gosec // Not security; reproducible (deterministic) fuzzing is useful
		val, node := randomItemNode(rng)
		tests = append(tests, test{
			fuzzed:   true,
			fuzzSeed: seed,
			value:    val,
			node:     node,
		})
	}

	throttle := make(chan struct{}, runtime.GOMAXPROCS(0))
	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			t.Parallel()
			throttle <- struct{}{}
			defer func() { <-throttle }()

			if tt.fuzzed {
				t.Logf("Fuzzing seed: %d", tt.fuzzSeed)
			}

			rlp, err := EncodeToBytes(tt.value)
			require.NoError(t, err, "EncodeToBytes([concrete value])")
			t.Logf("RLP encoding of %T: %#x\nValue: %x", tt.value, rlp, tt.value)

			t.Run("ParseTree", func(t *testing.T) {
				got, err := ParseTree(rlp)
				require.NoError(t, err, "ParseTree()")
				if diff := cmp.Diff(tt.node, got, cmpopts.EquateEmpty()); diff != "" {
					t.Errorf("ParseTree() diff (-want +got): \n%s", diff)
				}
			})

			t.Run("ItemNode.EncodeRLP()", func(t *testing.T) {
				got, err := EncodeToBytes(tt.node)
				require.NoErrorf(t, err, "EncodeToBytes(..., %T)", tt.node)
				if diff := cmp.Diff(rlp, got); diff != "" {
					t.Errorf("diff -Encode([concrete values]) +Encode(%T):\n%s", tt.node, diff)
				}
			})
		})
	}
}

func randomItemNode(rng *rand.Rand) (any, ItemNode) {
	switch rng.Intn(3) {
	case 0:
		v := byte(rng.Intn(128))
		if v == 0 {
			// https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp/#:~:text=thus%20making%20the%20integer%20value%20zero%20equivalent%20to%20the%20empty%20byte%20array
			return v, StringNode{}
		}
		return v, ByteNode(v)

	case 1:
		v := make([]byte, rng.Intn(110)) // >55 is a "long" string so ~half of each
		rng.Read(v)                      //nolint:gosec // Documented as returning nil error
		v = bytes.TrimLeft(v, "\x00")

		var node ItemNode
		switch n := len(v); {
		case n == 1 && v[0] < 128:
			node = ByteNode(v[0])
		default:
			node = StringNode(v)
		}
		return v, node

	default:
		var (
			vals []any
			list ListNode
		)
		for i, n := 0, rng.Intn(5); i < n; i++ {
			v, item := randomItemNode(rng)
			vals = append(vals, v)
			list = append(list, item)
		}
		return vals, list
	}
}

func TestEncodeItemNote(t *testing.T) {
	encodeToHex := func(v any) string {
		b, err := EncodeToBytes(v)
		require.NoError(t, err)
		return hex.EncodeToString(b)
	}
	makeBytes := func(n int) []byte {
		return append([]byte{1}, make([]byte, n-1)...)
	}

	tests := []struct {
		node    ItemNode
		wantHex string
	}{
		{ByteNode(0), "00"},
		{ByteNode(1), "01"},
		{ByteNode(127), "7f"},
		{ByteNode(128), "8180"},
		{ByteNode(255), "81ff"},
		{StringNode{}, "80"},
		{StringNode{0}, "8100"},
		{StringNode{1, 2, 3, 4, 5, 6, 7, 8, 9}, "89010203040506070809"},
		{
			StringNode(makeBytes(55)),
			encodeToHex(makeBytes(55)),
		},
		{
			StringNode(makeBytes(56)),
			encodeToHex(makeBytes(56)),
		},
	}

	for _, tt := range tests {
		var got bytes.Buffer
		require.NoError(t, tt.node.EncodeRLP(&got))
		assert.Equal(t, unhex(tt.wantHex), got.Bytes())
	}
}

func TestByteLength(t *testing.T) {
	tests := []struct{ n, want uint64 }{
		{0, 0},
		{1<<8 - 1, 1},
		{1 << 8, 2},
		{1<<16 - 1, 2},
		{1 << 16, 3},
		{1<<24 - 1, 3},
		{1 << 24, 4},
	}
	for _, tt := range tests {
		assert.Equalf(t, tt.want, byteLength(tt.n), "byteLength(%d)", tt.n)
	}
}

func TestBigEndian(t *testing.T) {
	assert.Empty(t, bigEndian(0), "bigEndian(0)")

	rng := rand.New(rand.NewSource(42)) //nolint:gosec // Reproducible fuzzing required

	for i := 1; i <= 8; i++ {
		prefix := make([]byte, 8-i)
		for j := 0; j < 100; j++ {
			suffix := make([]byte, i)
			for suffix[0] == 0 {
				rng.Read(suffix) //nolint:gosec // Documented as always returning nil error
			}

			n := binary.BigEndian.Uint64(append(prefix, suffix...))
			assert.Equalf(t, suffix, bigEndian(n), "bigEndian(%d)", n)
		}
	}

}
