package triedb

import (
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/trie"
	"github.com/ava-labs/libevm/trie/trienode"
	"github.com/ava-labs/libevm/triedb/pathdb"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
)

func TestLayerTree(t *testing.T) {
	// New database with 0 history limit
	var (
		historyLimit = uint64(0)
		disk         = rawdb.NewMemoryDatabase()
		db           = NewDatabase(disk, &Config{
			PathDB: &pathdb.Config{
				StateHistory:   historyLimit,
				CleanCacheSize: 256 * 1024,
				DirtyCacheSize: 256 * 1024,
			},
		})
	)

	// Create an initial state with two accounts
	a1 := common.Address{0x1}
	a2 := common.Address{0x2}
	a3 := common.Address{0x3}
	balance := uint256.NewInt(100)
	half := uint256.NewInt(50)

	root := types.EmptyRootHash
	root = nextState(t, db, root, 0, addBalance(a1, balance), addBalance(a2, balance))

	t.Logf("Root: %x", root)

	// We're going to make the same state C from R, but with paths A and B:
	// R -> A -> C
	// R -> B0 -> B1 -> C
	A := nextState(t, db, root, 1, pay(a1, a2, half), pay(a1, a2, half))
	B0 := nextState(t, db, root, 1, pay(a2, a1, half))
	B1 := nextState(t, db, B0, 2, pay(a2, a1, half))

	// a1 -> a3 payment is to make C different from R
	// reorder these two lines gives different diskLayer commits
	cFromB := nextState(t, db, B1, 3, pay(a1, a2, balance), pay(a1, a3, half))
	cFromA := nextState(t, db, A, 2, pay(a2, a1, balance), pay(a1, a3, half))

	assert.Equal(t, cFromA, cFromB)
	D := nextState(t, db, cFromA, 3, pay(a1, a3, half))

	t.Logf(
		"Roots: A: %x, B0: %x, B1: %x, C: %x, D: %x",
		A, B0, B1, cFromA, D,
	)

	// Also interesting behavior with CapTree(D, 2)
	if err := db.CapTree(D, 3); err != nil {
		t.Fatalf("Failed to cap tree, err: %v", err)
	}

	t.Logf("Successfully capped tree")
}

func nextState(
	t *testing.T, db *Database, parent common.Hash, block uint64, transitions ...stateTransition,
) common.Hash {
	tr, err := trie.NewStateTrie(trie.StateTrieID(parent), db)
	if err != nil {
		t.Fatalf("Failed to create trie, err: %v", err)
	}
	for _, transition := range transitions {
		transition(t, tr)
	}
	next, set, err := tr.Commit(false)
	if err != nil {
		t.Fatalf("Failed to commit trie, err: %v", err)
	}
	if err := db.Update(next, parent, block, trienode.NewWithNodeSet(set), nil); err != nil {
		t.Fatalf("Failed to update database, err: %v", err)
	}
	return next
}

type stateTransition func(*testing.T, *trie.StateTrie)

func addBalance(addr common.Address, amount *uint256.Int) stateTransition {
	return func(t *testing.T, tr *trie.StateTrie) {
		acc, err := tr.GetAccount(addr)
		if err != nil {
			t.Fatalf("Failed to get account, err: %v", err)
		}
		if acc == nil {
			acc = types.NewEmptyStateAccount()
		}

		acc.Balance = acc.Balance.Add(acc.Balance, amount)
		if err := tr.UpdateAccount(addr, acc); err != nil {
			t.Fatalf("Failed to update trie, err: %v", err)
		}
	}
}

func pay(from, to common.Address, amount *uint256.Int) stateTransition {
	return func(t *testing.T, tr *trie.StateTrie) {
		fromAcc, err := tr.GetAccount(from)
		if err != nil {
			t.Fatalf("Failed to get account, err: %v", err)
		}
		if fromAcc == nil {
			t.Fatalf("Account not found")
		}

		toAcc, err := tr.GetAccount(to)
		if err != nil {
			t.Fatalf("Failed to get account, err: %v", err)
		}
		if toAcc == nil {
			toAcc = types.NewEmptyStateAccount()
		}

		fromAcc.Balance = fromAcc.Balance.Sub(fromAcc.Balance, amount)
		toAcc.Balance = toAcc.Balance.Add(toAcc.Balance, amount)

		if err := tr.UpdateAccount(from, fromAcc); err != nil {
			t.Fatalf("Failed to update trie, err: %v", err)
		}
		if err := tr.UpdateAccount(to, toAcc); err != nil {
			t.Fatalf("Failed to update trie, err: %v", err)
		}
	}
}
