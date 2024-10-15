package state

import "github.com/ava-labs/libevm/triedb/triedbopts"

type Option interface {
	stateopt()
}

func WithTrieDBOptions(opts ...triedbopts.Option) Option {
	return triedbOptions(opts)
}

type triedbOptions []triedbopts.Option

func (triedbOptions) stateopt() {}

func FilterTrieDBOptions(opts []Option) (topts []triedbopts.Option) {
	for _, o := range opts {
		if to, ok := o.(triedbOptions); ok {
			topts = append(topts, ([]triedbopts.Option)(to)...)
		}
	}
	return topts
}
