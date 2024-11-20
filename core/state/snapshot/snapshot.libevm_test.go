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

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb/memorydb"
)

type hashLike interface {
	common.Hash | byte | int | string
}

func asHash[T hashLike](root T) common.Hash {
	switch r := any(root).(type) {
	case common.Hash:
		return r
	case byte:
		return common.Hash{r}
	case int:
		// It doesn't matter if this truncates as we only expect to use it with
		// single-byte hex literals.
		return common.Hash{byte(r)}
	case string:
		var h common.Hash
		copy(h[:], []byte(r))
		return h
	default:
		return common.Hash{}
	}
}

// update calls [Tree.Update] to create a layer with a single account, derived
// from `label`, which MUST be distinct and SHOULD be human-readable.
func update[T hashLike](tb testing.TB, tree *Tree, root byte, parent T, label string, opts ...LibEVMOption) {
	tb.Helper()

	accounts := map[common.Hash][]byte{
		asHash(label): []byte(label),
	}
	require.NoErrorf(
		tb,
		tree.Update(asHash(root), asHash(parent), nil, accounts, nil, opts...),
		"%T.Update(..., [%q], ...)", tree, label,
	)
}

type labelSet map[string]struct{}

// labelsOf assumes that the snapshot was created with [update] and returns all
// of the labels in the layer. If no flattening has been performed then only a
// single label is expected.
func labelsOf(tb testing.TB, snap Snapshot) labelSet {
	tb.Helper()
	require.NotNil(tb, snap, "Snapshot for label extraction")

	it := snap.(snapshot).AccountIterator(common.Hash{}) //nolint:forcetypeassert // Known invariant
	defer it.Release()

	labels := make(labelSet)
	for it.Next() {
		if err := it.Error(); err != nil {
			tb.Fatalf("%T.Error() = %v", it, err)
		}
		labels[string(it.Account())] = struct{}{}
	}
	return labels
}

// labelsOfPath is the multi-layer equivalent of [labelsOf].
func labelsOfPath(tb testing.TB, layers []Snapshot) []labelSet {
	tb.Helper()

	labels := make([]labelSet, len(layers))
	for i, l := range layers {
		labels[i] = labelsOf(tb, l)
	}
	return labels
}

// asLabelSets converts the slices of labels into sets; this is useful for
// defining test expectations.
func asLabelSets(labels ...[]string) []labelSet {
	sets := make([]labelSet, len(labels))
	for i, lbls := range labels {
		sets[i] = make(labelSet)
		for _, l := range lbls {
			sets[i][l] = struct{}{}
		}
	}
	return sets
}

func newEmptyTree(tb testing.TB) *Tree {
	tb.Helper()
	disk := memorydb.New()
	tree, err := New(
		Config{
			CacheSize: 16, // Mb (arbitrary but non-zero)
		},
		disk,
		nil, // triedb unnecessary as we're not rebuilding
		types.EmptyRootHash,
	)
	require.NoError(tb, err, "New()")
	return tree
}

// Labels used by [ambiguousTree].
const (
	A       = "A"
	B0      = "B0"
	B1      = "B1"
	CFromB0 = "C from B0"
	CFromB1 = "C from B1"
	DViaB0  = "D0 (via B0)"
	DViaB1  = "D1 (via B1)"
	EFromD0 = "E from D0"
	EFromD1 = "E from D1"
)

// ambiguousTree returns a tree with ambiguous roots 0xC and 0xE. They are,
// however, created via [update] with distinct labels so they can be
// differentiated with deeper inspection.
//
//	A
//	├── B0
//	│   └── C*
//	│       └── D0
//	│           └── E**
//	└── B1
//	    └── C*
//	        └── D1
//	            └── E**
func ambiguousTree(tb testing.TB) *Tree {
	tb.Helper()
	tree := newEmptyTree(tb)
	update(tb, tree, 0xA, tree.DiskRoot(), A)
	update[byte](tb, tree, 0xB0, 0xA, B0)
	update[byte](tb, tree, 0xB1, 0xA, B1)
	update[byte](tb, tree, 0xC, 0xB0, CFromB0)
	update[byte](tb, tree, 0xC, 0xB1, CFromB1)

	update[byte](tb, tree, 0xD0, 0xC, DViaB0, WithNextNearestAncestor(asHash(0xB0)))
	update[byte](tb, tree, 0xD1, 0xC, DViaB1, WithNextNearestAncestor(asHash(0xB1)))

	update[byte](tb, tree, 0xE, 0xD0, EFromD0)
	update[byte](tb, tree, 0xE, 0xD1, EFromD1)
	return tree
}

type pathTest struct {
	assertMsg string
	path      func(*Tree) []Snapshot
	want      [][]string
}

func (tt *pathTest) assert(tb testing.TB, tr *Tree) {
	assert.Equal(tb, asLabelSets(tt.want...), labelsOfPath(tb, tt.path(tr)), tt.assertMsg)
}

func TestAmbiguousTree(t *testing.T) {
	tree := ambiguousTree(t)
	tests := []pathTest{
		{
			path: func(tr *Tree) []Snapshot {
				return tr.Snapshots(asHash(0xE), 999, false, WithNextNearestAncestor(asHash(0xD0)))
			},
			want: [][]string{{EFromD0}, {DViaB0}, {CFromB0}, {B0}, {A}, {}},
		},
		{
			path: func(tr *Tree) []Snapshot {
				return tr.Snapshots(asHash(0xE), 999, false, WithNextNearestAncestor(asHash(0xD1)))
			},
			want: [][]string{{EFromD1}, {DViaB1}, {CFromB1}, {B1}, {A}, {}},
		},
	}
	for _, tt := range tests {
		tt.assert(t, tree)
	}
}

func TestCapZero(t *testing.T) {
	// All tests use [ambiguousTree]; see its comment for the layer structure.
	tests := []struct {
		name      string
		root      common.Hash
		opts      []LibEVMOption
		pathTests []pathTest
	}{
		{
			root: asHash(0xA),
			opts: []LibEVMOption{PreserveDescendantsOnCapZero()},
			pathTests: []pathTest{
				{
					path: func(tr *Tree) []Snapshot {
						// Including the disk demonstrates that the capped layer has
						// been flattened into it.
						return tr.Snapshots(asHash(0xE), 999, false /*nodisk*/, WithNextNearestAncestor(asHash(0xD0)))
					},
					want: [][]string{{EFromD0}, {DViaB0}, {CFromB0}, {B0}, {A}},
				},
				{
					path: func(tr *Tree) []Snapshot {
						return tr.Snapshots(asHash(0xE), 999, false, WithNextNearestAncestor(asHash(0xD1)))
					},
					want: [][]string{{EFromD1}, {DViaB1}, {CFromB1}, {B1}, {A}},
				},
			},
		},
		{
			root: asHash(0xC),
			opts: []LibEVMOption{
				WithNextNearestAncestor(asHash(0xB0)),
				PreserveDescendantsOnCapZero(),
			},
			pathTests: []pathTest{
				{
					path: func(tr *Tree) []Snapshot {
						return tr.Snapshots(asHash(0xE), 999, false, WithNextNearestAncestor(asHash(0xD0)))
					},
					want: [][]string{
						{EFromD0},
						{DViaB0},
						{CFromB0, B0, A}, // flattened to disk
					},
				},
				{
					path: func(tr *Tree) []Snapshot {
						return tr.Snapshots(asHash(0xE), 999, false, WithNextNearestAncestor(asHash(0xD1)))
					},
					want: [][]string{ /* E from D1 no longer exists */ },
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := ambiguousTree(t)
			require.NoError(t, tree.Cap(tt.root, 0, tt.opts...), "%T.Cap(%v, 0, ...)", tree, tt.root)
			for _, pt := range tt.pathTests {
				pt.assert(t, tree)
			}
		})
	}
}

func TestIsDescendantOf(t *testing.T) {
	tree := ambiguousTree(t)

	type test struct {
		name           string
		snap, ancestor Snapshot
		want           bool
	}

	tests := []test{
		{
			name:     "nil snap",
			ancestor: tree.Snapshot(asHash(0xA)),
			want:     false,
		},
		{
			name: "nil ancestor",
			snap: tree.Snapshot(asHash(0xA)),
			want: false,
		},
		{
			name:     "D1 is descendant of B1",
			snap:     tree.Snapshot(asHash(0xD1)),
			ancestor: tree.Snapshot(asHash(0xB1)),
			want:     true,
		},
		{
			name:     "D1 is not descendant of B0",
			snap:     tree.Snapshot(asHash(0xD1)),
			ancestor: tree.Snapshot(asHash(0xB0)),
			want:     false,
		},
		{
			name:     "B1 is ancestor of D1",
			snap:     tree.Snapshot(asHash(0xB1)),
			ancestor: tree.Snapshot(asHash(0xD1)),
			want:     false,
		},
		{
			name:     "E via D0 is descendant of B0",
			snap:     tree.Snapshot(asHash(0xE), WithNextNearestAncestor(asHash(0xD0))),
			ancestor: tree.Snapshot(asHash(0xB0)),
			want:     true,
		},
		{
			name:     "E via D1 is not descendant of B0",
			snap:     tree.Snapshot(asHash(0xE), WithNextNearestAncestor(asHash(0xD1))),
			ancestor: tree.Snapshot(asHash(0xB0)),
			want:     false,
		},
	}

	disk := tree.disklayer()
	for _, snaps := range tree.ancestry {
		for _, s := range snaps {
			diff, ok := s.(*diffLayer)
			if !ok {
				continue
			}
			tests = append(tests, test{
				name:     fmt.Sprintf("%s is descendant of disk", labelsOf(t, diff)),
				snap:     diff,
				ancestor: disk,
				want:     true,
			})
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isDescendendantOf(tt.snap, tt.ancestor))
		})
	}
}
