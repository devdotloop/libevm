// Copyright 2025 the libevm authors.
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

package miner

import (
	"context"
	"testing"

	"github.com/ava-labs/libevm/consensus"
	"github.com/ava-labs/libevm/consensus/ethash"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

type stubEngine struct {
	consensus.Engine
}

func TestBlockBuilder(t *testing.T) {
	ctx := context.Background()

	worker, backend := newTestWorker(t, ethashChainConfig, ethash.NewFaker(), rawdb.NewMemoryDatabase(), 0)

	var txs []*types.Transaction
	for i := 0; i < 10; i++ {
		tx := backend.newRandomTx(i%3 == 0 /*contract creation*/)
		errs := backend.TxPool().Add([]*types.Transaction{tx}, true, false)
		require.NoError(t, errs[0])
		txs = append(txs, tx)
	}

	builder := newBlockBuilder(worker)
	defer builder.Close()

	block, err := builder.Build(ctx)
	require.NoErrorf(t, err, "%T.Build()", builder)

	// The upstream [newTestWorker] constructor prepopulates the tx pool with
	// [pendingTxs].
	want := append(types.Transactions(pendingTxs), txs...)
	txEq := cmp.Comparer(func(a, b *types.Transaction) bool {
		return a.Hash() == b.Hash()
	})
	if diff := cmp.Diff(want, block.Transactions(), txEq); diff != "" {
		t.Errorf("%T.Build().Transactions() diff (-want +got):\n%s", builder, diff)
	}
}
