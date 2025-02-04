package downloader

import "github.com/ava-labs/libevm/core/types"

func (f *fetchResult) BlockBody() *types.Body {
	return &types.Body{
		Transactions: f.Transactions,
		Uncles:       f.Uncles,
	}
}
