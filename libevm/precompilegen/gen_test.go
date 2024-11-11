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
package main

import (
	"math/big"
	"testing"

	"github.com/holiman/uint256"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/vm"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/libevm"
	"github.com/ava-labs/libevm/libevm/ethtest"
	"github.com/ava-labs/libevm/libevm/hookstest"
	"github.com/ava-labs/libevm/libevm/precompilegen/testprecompile"
)

//go:generate solc -o ./ --overwrite --abi IPrecompile.sol
//go:generate go run . -in IPrecompile.abi -out ./testprecompile/generated.go -package testprecompile

func TestGeneratedPrecompile(t *testing.T) {
	addr := ethtest.NewPseudoRand(424242).Address()

	hooks := hookstest.Stub{
		PrecompileOverrides: map[common.Address]libevm.PrecompiledContract{
			addr: testprecompile.New(precompile{}),
		},
	}
	hooks.Register(t)
}

type precompile struct{}

var _ testprecompile.Contract = precompile{}

func (precompile) Fallback(env vm.PrecompileEnvironment, x []byte) ([]byte, uint64, error) {
	return nil, 0, nil
}

func (precompile) Echo(env vm.PrecompileEnvironment, x *big.Int) (*big.Int, error) {
	return x, nil
}

func (precompile) Echo0(env vm.PrecompileEnvironment, x string) (string, error) {
	return x, nil
}

func (precompile) Extract(env vm.PrecompileEnvironment, x struct {
	Val *big.Int "json:\"val\""
}) (*big.Int, error) {
	return x.Val, nil
}

func (precompile) HashPacked(env vm.PrecompileEnvironment, x *big.Int, y [2]byte, z common.Address) (hash [32]byte, _ error) {
	copy(
		hash[:],
		crypto.Keccak256(
			uint256.MustFromBig(x).PaddedBytes(32),
			y[:],
			z.Bytes(),
		),
	)
	return hash, nil
}

func (precompile) RevertWith(env vm.PrecompileEnvironment, x []byte) error {
	return vm.RevertError(x)
}

func (precompile) Self(env vm.PrecompileEnvironment) (common.Address, error) {
	return env.Addresses().Self, nil
}
