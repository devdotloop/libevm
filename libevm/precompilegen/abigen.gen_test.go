// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PrecompileTestMetaData contains all meta data concerning the PrecompileTest contract.
var PrecompileTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPrecompile\",\"name\":\"_precompile\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"func\",\"type\":\"string\"}],\"name\":\"Called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"Echo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"x\",\"type\":\"string\"}],\"name\":\"Echo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"bytes2\",\"name\":\"y\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"}],\"name\":\"HashPacked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"RevertWith\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Self\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051611253380380611253833981810160405281019061003191906100da565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505050610105565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100988261006f565b9050919050565b5f6100a98261008e565b9050919050565b6100b98161009f565b81146100c3575f5ffd5b50565b5f815190506100d4816100b0565b92915050565b5f602082840312156100ef576100ee61006b565b5b5f6100fc848285016100c6565b91505092915050565b60805161111461013f5f395f818160d6015281816101b60152818161031a0152818161035101528181610464015261056f01526111145ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806334d6d9be14610059578063a93cbd9714610075578063c62c692f14610091578063d7cc1f371461009b578063db84d7c0146100b7575b5f5ffd5b610073600480360381019061006e91906106b8565b6100d3565b005b61008f600480360381019061008a919061081f565b6101b2565b005b610099610318565b005b6100b560048036038101906100b09190610915565b610437565b005b6100d160048036038101906100cc9190610a03565b610546565b005b807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166334d6d9be836040518263ffffffff1660e01b815260040161012d9190610a59565b602060405180830381865afa158015610148573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061016c9190610a86565b1461017a57610179610ab1565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516101a790610b38565b60405180910390a150565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a93cbd9760e01b846040516024016102049190610bb6565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161026e9190610c10565b5f604051808303815f865af19150503d805f81146102a7576040519150601f19603f3d011682016040523d82523d5f602084013e6102ac565b606091505b509150915081156102c0576102bf610ab1565b5b82805190602001208180519060200120146102de576102dd610ab1565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a60405161030b90610c70565b60405180910390a1505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c62c692f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103b8573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103dc9190610ca2565b73ffffffffffffffffffffffffffffffffffffffff1614610400576103ff610ab1565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a60405161042d90610d17565b60405180910390a1565b82828260405160200161044c93929190610dba565b604051602081830303815290604052805190602001207f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d7cc1f378585856040518463ffffffff1660e01b81526004016104bf93929190610e14565b602060405180830381865afa1580156104da573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104fe9190610e7c565b1461050c5761050b610ab1565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a60405161053990610ef1565b60405180910390a1505050565b806040516020016105579190610f53565b604051602081830303815290604052805190602001207f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663db84d7c0836040518263ffffffff1660e01b81526004016105c69190610fa1565b5f60405180830381865afa1580156105e0573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610608919061102f565b6040516020016106189190610f53565b604051602081830303815290604052805190602001201461063c5761063b610ab1565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610669906110c0565b60405180910390a150565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b61069781610685565b81146106a1575f5ffd5b50565b5f813590506106b28161068e565b92915050565b5f602082840312156106cd576106cc61067d565b5b5f6106da848285016106a4565b91505092915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610731826106eb565b810181811067ffffffffffffffff821117156107505761074f6106fb565b5b80604052505050565b5f610762610674565b905061076e8282610728565b919050565b5f67ffffffffffffffff82111561078d5761078c6106fb565b5b610796826106eb565b9050602081019050919050565b828183375f83830152505050565b5f6107c36107be84610773565b610759565b9050828152602081018484840111156107df576107de6106e7565b5b6107ea8482856107a3565b509392505050565b5f82601f830112610806576108056106e3565b5b81356108168482602086016107b1565b91505092915050565b5f602082840312156108345761083361067d565b5b5f82013567ffffffffffffffff81111561085157610850610681565b5b61085d848285016107f2565b91505092915050565b5f7fffff00000000000000000000000000000000000000000000000000000000000082169050919050565b61089a81610866565b81146108a4575f5ffd5b50565b5f813590506108b581610891565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108e4826108bb565b9050919050565b6108f4816108da565b81146108fe575f5ffd5b50565b5f8135905061090f816108eb565b92915050565b5f5f5f6060848603121561092c5761092b61067d565b5b5f610939868287016106a4565b935050602061094a868287016108a7565b925050604061095b86828701610901565b9150509250925092565b5f67ffffffffffffffff82111561097f5761097e6106fb565b5b610988826106eb565b9050602081019050919050565b5f6109a76109a284610965565b610759565b9050828152602081018484840111156109c3576109c26106e7565b5b6109ce8482856107a3565b509392505050565b5f82601f8301126109ea576109e96106e3565b5b81356109fa848260208601610995565b91505092915050565b5f60208284031215610a1857610a1761067d565b5b5f82013567ffffffffffffffff811115610a3557610a34610681565b5b610a41848285016109d6565b91505092915050565b610a5381610685565b82525050565b5f602082019050610a6c5f830184610a4a565b92915050565b5f81519050610a808161068e565b92915050565b5f60208284031215610a9b57610a9a61067d565b5b5f610aa884828501610a72565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffd5b5f82825260208201905092915050565b7f4563686f2875696e7432353629000000000000000000000000000000000000005f82015250565b5f610b22600d83610ade565b9150610b2d82610aee565b602082019050919050565b5f6020820190508181035f830152610b4f81610b16565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610b8882610b56565b610b928185610b60565b9350610ba2818560208601610b70565b610bab816106eb565b840191505092915050565b5f6020820190508181035f830152610bce8184610b7e565b905092915050565b5f81905092915050565b5f610bea82610b56565b610bf48185610bd6565b9350610c04818560208601610b70565b80840191505092915050565b5f610c1b8284610be0565b915081905092915050565b7f52657665727457697468282e2e2e2900000000000000000000000000000000005f82015250565b5f610c5a600f83610ade565b9150610c6582610c26565b602082019050919050565b5f6020820190508181035f830152610c8781610c4e565b9050919050565b5f81519050610c9c816108eb565b92915050565b5f60208284031215610cb757610cb661067d565b5b5f610cc484828501610c8e565b91505092915050565b7f53656c66282900000000000000000000000000000000000000000000000000005f82015250565b5f610d01600683610ade565b9150610d0c82610ccd565b602082019050919050565b5f6020820190508181035f830152610d2e81610cf5565b9050919050565b5f819050919050565b610d4f610d4a82610685565b610d35565b82525050565b5f819050919050565b610d6f610d6a82610866565b610d55565b82525050565b5f8160601b9050919050565b5f610d8b82610d75565b9050919050565b5f610d9c82610d81565b9050919050565b610db4610daf826108da565b610d92565b82525050565b5f610dc58286610d3e565b602082019150610dd58285610d5e565b600282019150610de58284610da3565b601482019150819050949350505050565b610dff81610866565b82525050565b610e0e816108da565b82525050565b5f606082019050610e275f830186610a4a565b610e346020830185610df6565b610e416040830184610e05565b949350505050565b5f819050919050565b610e5b81610e49565b8114610e65575f5ffd5b50565b5f81519050610e7681610e52565b92915050565b5f60208284031215610e9157610e9061067d565b5b5f610e9e84828501610e68565b91505092915050565b7f486173685061636b6564282e2e2e2900000000000000000000000000000000005f82015250565b5f610edb600f83610ade565b9150610ee682610ea7565b602082019050919050565b5f6020820190508181035f830152610f0881610ecf565b9050919050565b5f81519050919050565b5f81905092915050565b5f610f2d82610f0f565b610f378185610f19565b9350610f47818560208601610b70565b80840191505092915050565b5f610f5e8284610f23565b915081905092915050565b5f610f7382610f0f565b610f7d8185610ade565b9350610f8d818560208601610b70565b610f96816106eb565b840191505092915050565b5f6020820190508181035f830152610fb98184610f69565b905092915050565b5f610fd3610fce84610965565b610759565b905082815260208101848484011115610fef57610fee6106e7565b5b610ffa848285610b70565b509392505050565b5f82601f830112611016576110156106e3565b5b8151611026848260208601610fc1565b91505092915050565b5f602082840312156110445761104361067d565b5b5f82015167ffffffffffffffff81111561106157611060610681565b5b61106d84828501611002565b91505092915050565b7f4563686f28737472696e672900000000000000000000000000000000000000005f82015250565b5f6110aa600c83610ade565b91506110b582611076565b602082019050919050565b5f6020820190508181035f8301526110d78161109e565b905091905056fea2646970667358221220dbd8e8b7b9412c1c2ca9fc32a128eab4ffca4ec73c31dc078ba49163bd3f407d64736f6c634300081c0033",
}

// PrecompileTestABI is the input ABI used to generate the binding from.
// Deprecated: Use PrecompileTestMetaData.ABI instead.
var PrecompileTestABI = PrecompileTestMetaData.ABI

// PrecompileTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PrecompileTestMetaData.Bin instead.
var PrecompileTestBin = PrecompileTestMetaData.Bin

// DeployPrecompileTest deploys a new Ethereum contract, binding an instance of PrecompileTest to it.
func DeployPrecompileTest(auth *bind.TransactOpts, backend bind.ContractBackend, _precompile common.Address) (common.Address, *types.Transaction, *PrecompileTest, error) {
	parsed, err := PrecompileTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PrecompileTestBin), backend, _precompile)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrecompileTest{PrecompileTestCaller: PrecompileTestCaller{contract: contract}, PrecompileTestTransactor: PrecompileTestTransactor{contract: contract}, PrecompileTestFilterer: PrecompileTestFilterer{contract: contract}}, nil
}

// PrecompileTest is an auto generated Go binding around an Ethereum contract.
type PrecompileTest struct {
	PrecompileTestCaller     // Read-only binding to the contract
	PrecompileTestTransactor // Write-only binding to the contract
	PrecompileTestFilterer   // Log filterer for contract events
}

// PrecompileTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompileTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompileTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompileTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompileTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompileTestSession struct {
	Contract     *PrecompileTest   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrecompileTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompileTestCallerSession struct {
	Contract *PrecompileTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PrecompileTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompileTestTransactorSession struct {
	Contract     *PrecompileTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PrecompileTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompileTestRaw struct {
	Contract *PrecompileTest // Generic contract binding to access the raw methods on
}

// PrecompileTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompileTestCallerRaw struct {
	Contract *PrecompileTestCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompileTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompileTestTransactorRaw struct {
	Contract *PrecompileTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompileTest creates a new instance of PrecompileTest, bound to a specific deployed contract.
func NewPrecompileTest(address common.Address, backend bind.ContractBackend) (*PrecompileTest, error) {
	contract, err := bindPrecompileTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrecompileTest{PrecompileTestCaller: PrecompileTestCaller{contract: contract}, PrecompileTestTransactor: PrecompileTestTransactor{contract: contract}, PrecompileTestFilterer: PrecompileTestFilterer{contract: contract}}, nil
}

// NewPrecompileTestCaller creates a new read-only instance of PrecompileTest, bound to a specific deployed contract.
func NewPrecompileTestCaller(address common.Address, caller bind.ContractCaller) (*PrecompileTestCaller, error) {
	contract, err := bindPrecompileTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileTestCaller{contract: contract}, nil
}

// NewPrecompileTestTransactor creates a new write-only instance of PrecompileTest, bound to a specific deployed contract.
func NewPrecompileTestTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompileTestTransactor, error) {
	contract, err := bindPrecompileTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompileTestTransactor{contract: contract}, nil
}

// NewPrecompileTestFilterer creates a new log filterer instance of PrecompileTest, bound to a specific deployed contract.
func NewPrecompileTestFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompileTestFilterer, error) {
	contract, err := bindPrecompileTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompileTestFilterer{contract: contract}, nil
}

// bindPrecompileTest binds a generic wrapper to an already deployed contract.
func bindPrecompileTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PrecompileTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompileTest *PrecompileTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrecompileTest.Contract.PrecompileTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompileTest *PrecompileTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileTest.Contract.PrecompileTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompileTest *PrecompileTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompileTest.Contract.PrecompileTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompileTest *PrecompileTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrecompileTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompileTest *PrecompileTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompileTest *PrecompileTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompileTest.Contract.contract.Transact(opts, method, params...)
}

// Echo is a paid mutator transaction binding the contract method 0x34d6d9be.
//
// Solidity: function Echo(uint256 x) returns()
func (_PrecompileTest *PrecompileTestTransactor) Echo(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "Echo", x)
}

// Echo is a paid mutator transaction binding the contract method 0x34d6d9be.
//
// Solidity: function Echo(uint256 x) returns()
func (_PrecompileTest *PrecompileTestSession) Echo(x *big.Int) (*types.Transaction, error) {
	return _PrecompileTest.Contract.Echo(&_PrecompileTest.TransactOpts, x)
}

// Echo is a paid mutator transaction binding the contract method 0x34d6d9be.
//
// Solidity: function Echo(uint256 x) returns()
func (_PrecompileTest *PrecompileTestTransactorSession) Echo(x *big.Int) (*types.Transaction, error) {
	return _PrecompileTest.Contract.Echo(&_PrecompileTest.TransactOpts, x)
}

// Echo0 is a paid mutator transaction binding the contract method 0xdb84d7c0.
//
// Solidity: function Echo(string x) returns()
func (_PrecompileTest *PrecompileTestTransactor) Echo0(opts *bind.TransactOpts, x string) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "Echo0", x)
}

// Echo0 is a paid mutator transaction binding the contract method 0xdb84d7c0.
//
// Solidity: function Echo(string x) returns()
func (_PrecompileTest *PrecompileTestSession) Echo0(x string) (*types.Transaction, error) {
	return _PrecompileTest.Contract.Echo0(&_PrecompileTest.TransactOpts, x)
}

// Echo0 is a paid mutator transaction binding the contract method 0xdb84d7c0.
//
// Solidity: function Echo(string x) returns()
func (_PrecompileTest *PrecompileTestTransactorSession) Echo0(x string) (*types.Transaction, error) {
	return _PrecompileTest.Contract.Echo0(&_PrecompileTest.TransactOpts, x)
}

// HashPacked is a paid mutator transaction binding the contract method 0xd7cc1f37.
//
// Solidity: function HashPacked(uint256 x, bytes2 y, address z) returns()
func (_PrecompileTest *PrecompileTestTransactor) HashPacked(opts *bind.TransactOpts, x *big.Int, y [2]byte, z common.Address) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "HashPacked", x, y, z)
}

// HashPacked is a paid mutator transaction binding the contract method 0xd7cc1f37.
//
// Solidity: function HashPacked(uint256 x, bytes2 y, address z) returns()
func (_PrecompileTest *PrecompileTestSession) HashPacked(x *big.Int, y [2]byte, z common.Address) (*types.Transaction, error) {
	return _PrecompileTest.Contract.HashPacked(&_PrecompileTest.TransactOpts, x, y, z)
}

// HashPacked is a paid mutator transaction binding the contract method 0xd7cc1f37.
//
// Solidity: function HashPacked(uint256 x, bytes2 y, address z) returns()
func (_PrecompileTest *PrecompileTestTransactorSession) HashPacked(x *big.Int, y [2]byte, z common.Address) (*types.Transaction, error) {
	return _PrecompileTest.Contract.HashPacked(&_PrecompileTest.TransactOpts, x, y, z)
}

// RevertWith is a paid mutator transaction binding the contract method 0xa93cbd97.
//
// Solidity: function RevertWith(bytes err) returns()
func (_PrecompileTest *PrecompileTestTransactor) RevertWith(opts *bind.TransactOpts, err []byte) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "RevertWith", err)
}

// RevertWith is a paid mutator transaction binding the contract method 0xa93cbd97.
//
// Solidity: function RevertWith(bytes err) returns()
func (_PrecompileTest *PrecompileTestSession) RevertWith(err []byte) (*types.Transaction, error) {
	return _PrecompileTest.Contract.RevertWith(&_PrecompileTest.TransactOpts, err)
}

// RevertWith is a paid mutator transaction binding the contract method 0xa93cbd97.
//
// Solidity: function RevertWith(bytes err) returns()
func (_PrecompileTest *PrecompileTestTransactorSession) RevertWith(err []byte) (*types.Transaction, error) {
	return _PrecompileTest.Contract.RevertWith(&_PrecompileTest.TransactOpts, err)
}

// Self is a paid mutator transaction binding the contract method 0xc62c692f.
//
// Solidity: function Self() returns()
func (_PrecompileTest *PrecompileTestTransactor) Self(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "Self")
}

// Self is a paid mutator transaction binding the contract method 0xc62c692f.
//
// Solidity: function Self() returns()
func (_PrecompileTest *PrecompileTestSession) Self() (*types.Transaction, error) {
	return _PrecompileTest.Contract.Self(&_PrecompileTest.TransactOpts)
}

// Self is a paid mutator transaction binding the contract method 0xc62c692f.
//
// Solidity: function Self() returns()
func (_PrecompileTest *PrecompileTestTransactorSession) Self() (*types.Transaction, error) {
	return _PrecompileTest.Contract.Self(&_PrecompileTest.TransactOpts)
}

// PrecompileTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the PrecompileTest contract.
type PrecompileTestCalledIterator struct {
	Event *PrecompileTestCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PrecompileTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrecompileTestCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PrecompileTestCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PrecompileTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrecompileTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrecompileTestCalled represents a Called event raised by the PrecompileTest contract.
type PrecompileTestCalled struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0x3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a.
//
// Solidity: event Called(string func)
func (_PrecompileTest *PrecompileTestFilterer) FilterCalled(opts *bind.FilterOpts) (*PrecompileTestCalledIterator, error) {

	logs, sub, err := _PrecompileTest.contract.FilterLogs(opts, "Called")
	if err != nil {
		return nil, err
	}
	return &PrecompileTestCalledIterator{contract: _PrecompileTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a.
//
// Solidity: event Called(string func)
func (_PrecompileTest *PrecompileTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *PrecompileTestCalled) (event.Subscription, error) {

	logs, sub, err := _PrecompileTest.contract.WatchLogs(opts, "Called")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrecompileTestCalled)
				if err := _PrecompileTest.contract.UnpackLog(event, "Called", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCalled is a log parse operation binding the contract event 0x3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a.
//
// Solidity: event Called(string func)
func (_PrecompileTest *PrecompileTestFilterer) ParseCalled(log types.Log) (*PrecompileTestCalled, error) {
	event := new(PrecompileTestCalled)
	if err := _PrecompileTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
