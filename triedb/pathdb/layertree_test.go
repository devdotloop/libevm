package pathdb

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/stretchr/testify/require"
)

func TestPathLoopNonCanonicalStatePersistsToDisk(t *testing.T) {
	disk := rawdb.NewMemoryDatabase()
	db := New(disk, Defaults)
	defer db.Close()
	nodes := trienode.NewMergedNodeSet()

	// Let's start at state r
	r := common.HexToHash("0x0001")

	// Possible state transitions A, B
	rA := common.HexToHash("0x0011")
	rB := common.HexToHash("0x0101")

	// Result of applying A and B (we assume they commute)
	rAB := common.HexToHash("0x0111")

	// Build path: r -> r+A -> r+A+B
	require.NoError(t, db.Update(r, db.tree.bottom().root, 1, nodes, nil))
	require.NoError(t, db.Update(rA, r, 2, nodes, nil))
	require.NoError(t, db.Update(rAB, rA, 3, nodes, nil))

	// Now we insert path: r -> r+B -> r+B+A
	require.NoError(t, db.Update(rB, r, 2, nodes, nil))
	require.NoError(t, db.Update(rAB, rB, 3, nodes, nil))

	// Blockchain sets the preferred head back to the original path
	// But this is not reflected in triedb

	// An additional state transition
	rABC := common.HexToHash("0x1111")
	// This will be added to the non-preferred path, as when rAB is added the
	// second time, its parent becomes rB, as it overwrites the previous entry
	// for rAB with parent rA.
	// r -> r+B -> r+B+A -> r+B+A+C
	require.NoError(t, db.Update(rABC, rAB, 4, nodes, nil))

	// calling cap with 2 here, but we can imagine this happens at the end of a
	// longer chain.
	db.tree.cap(rABC, 2)

	t.Log(rawdb.ReadStateID(disk, rA))  // this is <nil>, but rA was on the canonical chain
	t.Log(*rawdb.ReadStateID(disk, rB)) // this is on disk, but rB is not on the canonical chain

	// On a clean shutdown, the journal will perist the difflayers, and continuing
	// is possible.

	// On an unclean shutdown, the journal will not be persisted, so recovery is
	// not possible, as rB does not correspond to any state in the canonical chain.
}

func TestPathLoopCapNonInsert(t *testing.T) {
	disk := rawdb.NewMemoryDatabase()
	db := New(disk, Defaults)
	defer db.Close()
	nodes := trienode.NewMergedNodeSet()

	// base -> r1 -> r2 -> r3 -> r4
	root1 := common.HexToHash("0x01")
	root2 := common.HexToHash("0x02")
	root3 := common.HexToHash("0x03")
	root4 := common.HexToHash("0x04")

	require.NoError(t, db.Update(root1, db.tree.bottom().root, 1, nodes, nil))
	require.NoError(t, db.Update(root2, root1, 2, nodes, nil))
	require.NoError(t, db.Update(root3, root2, 3, nodes, nil))
	require.NoError(t, db.Update(root4, root3, 4, nodes, nil))

	// Maybe it's also possible to directly create r4 via r1
	// base -> r1 -> r2 -> r3 -> r4
	//           \______________/
	require.NoError(t, db.Update(root4, root1, 2, nodes, nil))

	// If we try to call cap at r3 with 1, layer r4 will be removed.
	// This is not an issue if we only ever call cap during insert,
	// but we may want to call cap at other times like accept, and
	// still preserve the verified tree.
	db.tree.cap(root3, 1)
	t.Log(rawdb.ReadStateID(disk, root4)) // this is <nil>, but r4 was on the verified chain
}

func TestPathLoopWrongPersistedStateID(t *testing.T) {
	disk := rawdb.NewMemoryDatabase()
	db := New(disk, Defaults)
	defer db.Close()
	nodes := trienode.NewMergedNodeSet()

	// base -> r1 -> r2 -> r3 -> r4
	root1 := common.HexToHash("0x01")
	root2 := common.HexToHash("0x02")
	root3 := common.HexToHash("0x03")
	root4 := common.HexToHash("0x04")

	t.Log(db.tree.bottom().stateID())
	require.NoError(t, db.Update(root1, db.tree.bottom().root, 1, nodes, nil))
	require.NoError(t, db.Update(root2, root1, 2, nodes, nil))
	require.NoError(t, db.Update(root3, root2, 3, nodes, nil))
	require.NoError(t, db.Update(root4, root3, 4, nodes, nil))

	// Maybe it's also possible to directly create r4 via r1
	// base -> r1 -> r2 -> r3 -> r4
	//           \______________/
	require.NoError(t, db.Update(root4, root1, 2, nodes, nil))

	// Let's continue with r5
	// base -> r1 -> r2 -> r3 -> r4 -> r5
	//           \______________/
	root5 := common.HexToHash("0x05")
	require.NoError(t, db.Update(root5, root4, 5, nodes, nil))

	db.tree.cap(root5, 1)
	t.Log(*rawdb.ReadStateID(disk, root4)) // This is on disk, but the number should be 4, not 2
	// Note this also corrupts the stateID number for r5, which is now 3 instead of 5.
	t.Log(db.tree.get(root5).stateID())
}

func TestPathLoopKeepsLayersInMemory(t *testing.T) {
	disk := rawdb.NewMemoryDatabase()
	db := New(disk, Defaults)
	defer db.Close()
	nodes := trienode.NewMergedNodeSet()

	// base -> r1 -> r2 -> r3 -> r4
	root1 := common.HexToHash("0x01")
	root2 := common.HexToHash("0x02")
	root3 := common.HexToHash("0x03")
	root4 := common.HexToHash("0x04")

	t.Log(db.tree.bottom().stateID())
	require.NoError(t, db.Update(root1, db.tree.bottom().root, 1, nodes, nil))
	require.NoError(t, db.Update(root2, root1, 2, nodes, nil))
	require.NoError(t, db.Update(root3, root2, 3, nodes, nil))
	require.NoError(t, db.Update(root4, root3, 4, nodes, nil))

	// Maybe it's also possible to directly create r4 via r1
	// base -> r1 -> r2 -> r3 -> r4
	//           \______________/
	require.NoError(t, db.Update(root4, root1, 2, nodes, nil))

	// Let's continue with r5
	// base -> r1 -> r2 -> r3 -> r4 -> r5
	//           \______________/
	root5 := common.HexToHash("0x05")
	require.NoError(t, db.Update(root5, root4, 5, nodes, nil))

	db.tree.cap(root5, 2)
	// We would expect this to remove r1, r2, and r3.
	// But it doesn't because of the path loop.
	for _, layer := range db.tree.layers {
		t.Log(layer.rootHash(), layer.stateID())
	}
}
