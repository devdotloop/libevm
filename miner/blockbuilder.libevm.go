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
	"errors"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/consensus"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/params"
)

type BlockBuilder struct {
	worker  *worker
	results chan *newPayloadResult
}

func NewBlockBuilder(eth Backend, config *Config, chainConfig *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine, isLocalBlock func(header *types.Header) bool) *BlockBuilder {
	return newBlockBuilder(newWorker(config, chainConfig, engine, eth, mux, isLocalBlock, true))
}

func newBlockBuilder(w *worker) *BlockBuilder {
	return &BlockBuilder{
		worker:  w,
		results: make(chan *newPayloadResult),
	}
}

func (b *BlockBuilder) Close() {
	b.worker.close()
	close(b.results)
}

var errBlockBuilderClosed = errors.New("BlockBuilder closed")

func (b *BlockBuilder) Build(ctx context.Context) (*types.Block, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-b.worker.exitCh:
		return nil, errBlockBuilderClosed

	case b.worker.getWorkCh <- &getWorkReq{
		params: &generateParams{
			timestamp:  uint64(time.Now().Unix()),
			coinbase:   b.worker.etherbase(),
			beaconRoot: &common.Hash{}, // TODO(arr4n): is this necessary? Cargo-culted from ava-labs/coreth
		},
		result: b.results,
	}:
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-b.worker.exitCh:
		return nil, errBlockBuilderClosed

	case r, ok := <-b.results:
		if !ok {
			return nil, errBlockBuilderClosed
		}
		if r.err != nil {
			return nil, r.err
		}
		return r.block, nil
	}
}
