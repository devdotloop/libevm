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

package snapshot

import "github.com/ava-labs/libevm/common"

// A LibEVMOption configures default behaviour of this package.
type LibEVMOption interface {
	apply(*libevmConfig)
}

type libevmConfig struct {
	ancestorRoot                 *common.Hash
	preserveDescendantsOnCapZero bool
}

func asLibEVMConfig(opts []LibEVMOption) *libevmConfig {
	c := new(libevmConfig)
	for _, o := range opts {
		o.apply(c)
	}
	return c
}

type libevmFuncOpt func(*libevmConfig)

func (f libevmFuncOpt) apply(c *libevmConfig) { f(c) }

// WithNextNearestAncestor specifies the nearest snapshot ancestor not already
// provided as another argument. This is the grandparent when calling
// [Tree.Update] and the parent when calling methods that do not already accept
// the parent, such as [Tree.Snapshot].
//
// Although modelled as a tree, there may be multiple paths that result in the
// same state root, such that a node (a layer) identified by its root alone may
// be ambiguous. Specifying a parent is enough to remove this ambiguity. In the
// case of [Tree.Update] we wish to disambiguate the new layer's parent, hence
// specification of a grandparent.
//
// This option is ignored if used when selecting the disk layer as it is already
// unambiguous and its ancestors are unknown.
func WithNextNearestAncestor(root common.Hash) LibEVMOption {
	return libevmFuncOpt(func(c *libevmConfig) {
		c.ancestorRoot = &root
	})
}

// recordAncestryWhenLocked assumes that a write lock is already held and then
// stores the snapshot keyed by both its and its parent's roots. See
// [WithNextNearestAncestor] for rationale.
func (t *Tree) recordAncestryWhenLocked(snap snapshot) {
	if t.ancestry == nil {
		t.ancestry = make(map[common.Hash]map[common.Hash]snapshot)
	}
	if r := snap.Root(); t.ancestry[r] == nil {
		t.ancestry[r] = make(map[common.Hash]snapshot)
	}
	t.ancestry[snap.Root()][snap.Parent().Root()] = snap
}

// layerWhenLocked assumes that a read lock is already held and then returns the
// requested snapshot. If no ancestral root hash is provided then the most
// recent snapshot with the given root is returned, in keeping with default geth
// behaviour.
func (t *Tree) layerWhenLocked(root common.Hash, opts ...LibEVMOption) snapshot {
	if disk, ok := t.layers[root].(*diskLayer); ok {
		return disk
	}
	if a := asLibEVMConfig(opts).ancestorRoot; a != nil {
		return t.ancestry[root][*a]
	}
	return t.layers[root]
}

// PreserveDescendantsOnCapZero signals to [Tree.Cap], if capping to zero layers
// (i.e. flattening), that descendants of the flattened layer must be kept.
// Without this option, the entire tree is replaced with the new base (disk)
// layer.
func PreserveDescendantsOnCapZero() LibEVMOption {
	return libevmFuncOpt(func(c *libevmConfig) {
		c.preserveDescendantsOnCapZero = true
	})
}

func (t *Tree) updateLayersAfterCapZero(base *diskLayer, flattened *diffLayer, opts []LibEVMOption) {
	// Original geth behaviour
	t.layers = map[common.Hash]snapshot{base.root: base}
	t.ancestry = nil
	if !asLibEVMConfig(opts).preserveDescendantsOnCapZero {
		return
	}

	// TODO(arr4n): this isn't the most efficient algorithm as it requires
	// every descendant to traverse up to `flattened`.
	var descendants []*diffLayer
	for _, snaps := range t.ancestry {
		for _, s := range snaps {
			if isDescendendantOf(s, flattened) {
				descendants = append(descendants, s.(*diffLayer)) //nolint:forcetypeassert // Guaranteed by being a descendant of something
			}
		}
	}

	for _, snap := range descendants {
		snap.origin = base
		if areSameSnapshot(snap.Parent(), flattened) {
			snap.parent = base
		}

		t.layers[snap.Root()] = snap
		t.recordAncestryWhenLocked(snap)
	}
}

func isDescendendantOf(snap, ancestor Snapshot) bool {
	for {
		if _, isDisk := snap.(*diskLayer); isDisk || snap == nil || ancestor == nil {
			return false
		}
		snap = snap.(snapshot).Parent()
		if areSameSnapshot(snap, ancestor) {
			return true
		}
	}
}

func areSameSnapshot(a, b Snapshot) bool {
	switch a := a.(type) {
	case *diskLayer:
		b, ok := b.(*diskLayer)
		return ok && a == b

	case *diffLayer:
		b, ok := b.(*diffLayer)
		return ok && a == b
	}
	return false
}
