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

// IPrecompileWrapper is an auto generated low-level Go binding around an user-defined struct.
type IPrecompileWrapper struct {
	Val *big.Int
}

// PrecompileTestMetaData contains all meta data concerning the PrecompileTest contract.
var PrecompileTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPrecompile\",\"name\":\"_precompile\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"func\",\"type\":\"string\"}],\"name\":\"Called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"Echo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"x\",\"type\":\"string\"}],\"name\":\"Echo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"EchoingFallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"internalType\":\"structIPrecompile.Wrapper\",\"name\":\"x\",\"type\":\"tuple\"}],\"name\":\"Extract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"bytes2\",\"name\":\"y\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"}],\"name\":\"HashPacked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NeitherViewNorPure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Pure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"RevertWith\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Self\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"View\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052604051611d49380380611d49833981810160405281019061002591906100ce565b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050506100f9565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61008c82610063565b9050919050565b5f61009d82610082565b9050919050565b6100ad81610093565b81146100b7575f5ffd5b50565b5f815190506100c8816100a4565b92915050565b5f602082840312156100e3576100e261005f565b5b5f6100f0848285016100ba565b91505092915050565b608051611bf26101575f395f81816101d2015281816102b101528181610356015281816104de015281816106420152818161077c0152818161086a015281816108a1015281816109b401528181610abf0152610c110152611bf25ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c8063ad8108a41161006f578063ad8108a414610101578063b4a818081461011d578063c62c692f14610139578063d7cc1f3714610143578063db84d7c01461015f578063f5ad2fb81461017b576100a7565b80631686f265146100ab57806334d6d9be146100b5578063406dade3146100d1578063a7263000146100db578063a93cbd97146100e5575b5f5ffd5b6100b3610185565b005b6100cf60048036038101906100ca9190610d97565b6101cf565b005b6100d96102ae565b005b6100e3610490565b005b6100ff60048036038101906100fa9190610efe565b6104da565b005b61011b60048036038101906101169190610fb5565b610640565b005b61013760048036038101906101329190610efe565b610722565b005b610141610868565b005b61015d6004803603810190610158919061108f565b610987565b005b6101796004803603810190610174919061117d565b610a96565b005b610183610bc4565b005b610198631686f26560e01b60015f610c0d565b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516101c59061121e565b60405180910390a1565b807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166334d6d9be836040518263ffffffff1660e01b8152600401610229919061124b565b602060405180830381865afa158015610244573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102689190611278565b14610276576102756112a3565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516102a39061131a565b60405180910390a150565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631b7fb4b0602a6040518263ffffffff1660e01b815260040160206040518083038185885af115801561031b573d5f5f3e3d5ffd5b50505050506040513d601f19601f820116820180604052508101906103409190611278565b9050602a8114610353576103526112a3565b5b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16602a636fb1b0e960e01b604051602401604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610405919061138a565b5f6040518083038185875af1925050503d805f811461043f576040519150601f19603f3d011682016040523d82523d5f602084013e610444565b606091505b505090508015610457576104566112a3565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610484906113ea565b60405180910390a15050565b6104a363a726300060e01b600180610c0d565b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516104d090611452565b60405180910390a1565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a93cbd9760e01b8460405160240161052c91906114b8565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610596919061138a565b5f604051808303815f865af19150503d805f81146105cf576040519150601f19603f3d011682016040523d82523d5f602084013e6105d4565b606091505b509150915081156105e8576105e76112a3565b5b8280519060200120818051906020012014610606576106056112a3565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a60405161063390611522565b60405180910390a1505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663ad8108a4826040518263ffffffff1660e01b81526004016106999190611569565b602060405180830381865afa1580156106b4573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106d89190611596565b815f0151146106ea576106e96112a3565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516107179061160b565b60405180910390a150565b5f5f8260405160240161073591906114b8565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16836040516107bf919061138a565b5f60405180830381855afa9150503d805f81146107f7576040519150601f19603f3d011682016040523d82523d5f602084013e6107fc565b606091505b50915091508161080f5761080e6112a3565b5b828051906020012081805190602001201461082d5761082c6112a3565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a60405161085a90611673565b60405180910390a150505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c62c692f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610908573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061092c91906116a5565b73ffffffffffffffffffffffffffffffffffffffff16146109505761094f6112a3565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a60405161097d9061171a565b60405180910390a1565b82828260405160200161099c939291906117bd565b604051602081830303815290604052805190602001207f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d7cc1f378585856040518463ffffffff1660e01b8152600401610a0f93929190611817565b602060405180830381865afa158015610a2a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a4e919061187f565b14610a5c57610a5b6112a3565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610a89906118f4565b60405180910390a1505050565b80604051602001610aa79190611956565b604051602081830303815290604052805190602001207f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663db84d7c0836040518263ffffffff1660e01b8152600401610b1691906119a4565b5f60405180830381865afa158015610b30573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610b589190611a32565b604051602001610b689190611956565b6040516020818303038152906040528051906020012014610b8c57610b8b6112a3565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610bb990611ac3565b60405180910390a150565b610bd663f5ad2fb860e01b5f5f610c0d565b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610c0390611b2b565b60405180910390a1565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1685604051602401604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610cb7919061138a565b5f604051808303815f865af19150503d805f8114610cf0576040519150601f19603f3d011682016040523d82523d5f602084013e610cf5565b606091505b509150915081610d0857610d076112a3565b5b5f5f82806020019051810190610d1e9190611b7e565b9150915085151582151514610d3657610d356112a3565b5b84151581151514610d4a57610d496112a3565b5b50505050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610d7681610d64565b8114610d80575f5ffd5b50565b5f81359050610d9181610d6d565b92915050565b5f60208284031215610dac57610dab610d5c565b5b5f610db984828501610d83565b91505092915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e1082610dca565b810181811067ffffffffffffffff82111715610e2f57610e2e610dda565b5b80604052505050565b5f610e41610d53565b9050610e4d8282610e07565b919050565b5f67ffffffffffffffff821115610e6c57610e6b610dda565b5b610e7582610dca565b9050602081019050919050565b828183375f83830152505050565b5f610ea2610e9d84610e52565b610e38565b905082815260208101848484011115610ebe57610ebd610dc6565b5b610ec9848285610e82565b509392505050565b5f82601f830112610ee557610ee4610dc2565b5b8135610ef5848260208601610e90565b91505092915050565b5f60208284031215610f1357610f12610d5c565b5b5f82013567ffffffffffffffff811115610f3057610f2f610d60565b5b610f3c84828501610ed1565b91505092915050565b5f5ffd5b5f819050919050565b610f5b81610f49565b8114610f65575f5ffd5b50565b5f81359050610f7681610f52565b92915050565b5f60208284031215610f9157610f90610f45565b5b610f9b6020610e38565b90505f610faa84828501610f68565b5f8301525092915050565b5f60208284031215610fca57610fc9610d5c565b5b5f610fd784828501610f7c565b91505092915050565b5f7fffff00000000000000000000000000000000000000000000000000000000000082169050919050565b61101481610fe0565b811461101e575f5ffd5b50565b5f8135905061102f8161100b565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61105e82611035565b9050919050565b61106e81611054565b8114611078575f5ffd5b50565b5f8135905061108981611065565b92915050565b5f5f5f606084860312156110a6576110a5610d5c565b5b5f6110b386828701610d83565b93505060206110c486828701611021565b92505060406110d58682870161107b565b9150509250925092565b5f67ffffffffffffffff8211156110f9576110f8610dda565b5b61110282610dca565b9050602081019050919050565b5f61112161111c846110df565b610e38565b90508281526020810184848401111561113d5761113c610dc6565b5b611148848285610e82565b509392505050565b5f82601f83011261116457611163610dc2565b5b813561117484826020860161110f565b91505092915050565b5f6020828403121561119257611191610d5c565b5b5f82013567ffffffffffffffff8111156111af576111ae610d60565b5b6111bb84828501611150565b91505092915050565b5f82825260208201905092915050565b7f56696577282900000000000000000000000000000000000000000000000000005f82015250565b5f6112086006836111c4565b9150611213826111d4565b602082019050919050565b5f6020820190508181035f830152611235816111fc565b9050919050565b61124581610d64565b82525050565b5f60208201905061125e5f83018461123c565b92915050565b5f8151905061127281610d6d565b92915050565b5f6020828403121561128d5761128c610d5c565b5b5f61129a84828501611264565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffd5b7f4563686f2875696e7432353629000000000000000000000000000000000000005f82015250565b5f611304600d836111c4565b915061130f826112d0565b602082019050919050565b5f6020820190508181035f830152611331816112f8565b9050919050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61136482611338565b61136e8185611342565b935061137e81856020860161134c565b80840191505092915050565b5f611395828461135a565b915081905092915050565b7f5472616e736665722829000000000000000000000000000000000000000000005f82015250565b5f6113d4600a836111c4565b91506113df826113a0565b602082019050919050565b5f6020820190508181035f830152611401816113c8565b9050919050565b7f4e656974686572566965774e6f725075726528290000000000000000000000005f82015250565b5f61143c6014836111c4565b915061144782611408565b602082019050919050565b5f6020820190508181035f83015261146981611430565b9050919050565b5f82825260208201905092915050565b5f61148a82611338565b6114948185611470565b93506114a481856020860161134c565b6114ad81610dca565b840191505092915050565b5f6020820190508181035f8301526114d08184611480565b905092915050565b7f52657665727457697468282e2e2e2900000000000000000000000000000000005f82015250565b5f61150c600f836111c4565b9150611517826114d8565b602082019050919050565b5f6020820190508181035f83015261153981611500565b9050919050565b61154981610f49565b82525050565b602082015f8201516115635f850182611540565b50505050565b5f60208201905061157c5f83018461154f565b92915050565b5f8151905061159081610f52565b92915050565b5f602082840312156115ab576115aa610d5c565b5b5f6115b884828501611582565b91505092915050565b7f45787472616374282e2e2e2900000000000000000000000000000000000000005f82015250565b5f6115f5600c836111c4565b9150611600826115c1565b602082019050919050565b5f6020820190508181035f830152611622816115e9565b9050919050565b7f4563686f696e6746616c6c6261636b282e2e2e290000000000000000000000005f82015250565b5f61165d6014836111c4565b915061166882611629565b602082019050919050565b5f6020820190508181035f83015261168a81611651565b9050919050565b5f8151905061169f81611065565b92915050565b5f602082840312156116ba576116b9610d5c565b5b5f6116c784828501611691565b91505092915050565b7f53656c66282900000000000000000000000000000000000000000000000000005f82015250565b5f6117046006836111c4565b915061170f826116d0565b602082019050919050565b5f6020820190508181035f830152611731816116f8565b9050919050565b5f819050919050565b61175261174d82610d64565b611738565b82525050565b5f819050919050565b61177261176d82610fe0565b611758565b82525050565b5f8160601b9050919050565b5f61178e82611778565b9050919050565b5f61179f82611784565b9050919050565b6117b76117b282611054565b611795565b82525050565b5f6117c88286611741565b6020820191506117d88285611761565b6002820191506117e882846117a6565b601482019150819050949350505050565b61180281610fe0565b82525050565b61181181611054565b82525050565b5f60608201905061182a5f83018661123c565b61183760208301856117f9565b6118446040830184611808565b949350505050565b5f819050919050565b61185e8161184c565b8114611868575f5ffd5b50565b5f8151905061187981611855565b92915050565b5f6020828403121561189457611893610d5c565b5b5f6118a18482850161186b565b91505092915050565b7f486173685061636b6564282e2e2e2900000000000000000000000000000000005f82015250565b5f6118de600f836111c4565b91506118e9826118aa565b602082019050919050565b5f6020820190508181035f83015261190b816118d2565b9050919050565b5f81519050919050565b5f81905092915050565b5f61193082611912565b61193a818561191c565b935061194a81856020860161134c565b80840191505092915050565b5f6119618284611926565b915081905092915050565b5f61197682611912565b61198081856111c4565b935061199081856020860161134c565b61199981610dca565b840191505092915050565b5f6020820190508181035f8301526119bc818461196c565b905092915050565b5f6119d66119d1846110df565b610e38565b9050828152602081018484840111156119f2576119f1610dc6565b5b6119fd84828561134c565b509392505050565b5f82601f830112611a1957611a18610dc2565b5b8151611a298482602086016119c4565b91505092915050565b5f60208284031215611a4757611a46610d5c565b5b5f82015167ffffffffffffffff811115611a6457611a63610d60565b5b611a7084828501611a05565b91505092915050565b7f4563686f28737472696e672900000000000000000000000000000000000000005f82015250565b5f611aad600c836111c4565b9150611ab882611a79565b602082019050919050565b5f6020820190508181035f830152611ada81611aa1565b9050919050565b7f50757265282900000000000000000000000000000000000000000000000000005f82015250565b5f611b156006836111c4565b9150611b2082611ae1565b602082019050919050565b5f6020820190508181035f830152611b4281611b09565b9050919050565b5f8115159050919050565b611b5d81611b49565b8114611b67575f5ffd5b50565b5f81519050611b7881611b54565b92915050565b5f5f60408385031215611b9457611b93610d5c565b5b5f611ba185828601611b6a565b9250506020611bb285828601611b6a565b915050925092905056fea264697066735822122015cec228d3431d1c7e411919e8c1d0eedf25bb9b34ed83e88e4dbcb442ad811764736f6c634300081c0033",
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

// EchoingFallback is a paid mutator transaction binding the contract method 0xb4a81808.
//
// Solidity: function EchoingFallback(bytes input) returns()
func (_PrecompileTest *PrecompileTestTransactor) EchoingFallback(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "EchoingFallback", input)
}

// EchoingFallback is a paid mutator transaction binding the contract method 0xb4a81808.
//
// Solidity: function EchoingFallback(bytes input) returns()
func (_PrecompileTest *PrecompileTestSession) EchoingFallback(input []byte) (*types.Transaction, error) {
	return _PrecompileTest.Contract.EchoingFallback(&_PrecompileTest.TransactOpts, input)
}

// EchoingFallback is a paid mutator transaction binding the contract method 0xb4a81808.
//
// Solidity: function EchoingFallback(bytes input) returns()
func (_PrecompileTest *PrecompileTestTransactorSession) EchoingFallback(input []byte) (*types.Transaction, error) {
	return _PrecompileTest.Contract.EchoingFallback(&_PrecompileTest.TransactOpts, input)
}

// Extract is a paid mutator transaction binding the contract method 0xad8108a4.
//
// Solidity: function Extract((int256) x) returns()
func (_PrecompileTest *PrecompileTestTransactor) Extract(opts *bind.TransactOpts, x IPrecompileWrapper) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "Extract", x)
}

// Extract is a paid mutator transaction binding the contract method 0xad8108a4.
//
// Solidity: function Extract((int256) x) returns()
func (_PrecompileTest *PrecompileTestSession) Extract(x IPrecompileWrapper) (*types.Transaction, error) {
	return _PrecompileTest.Contract.Extract(&_PrecompileTest.TransactOpts, x)
}

// Extract is a paid mutator transaction binding the contract method 0xad8108a4.
//
// Solidity: function Extract((int256) x) returns()
func (_PrecompileTest *PrecompileTestTransactorSession) Extract(x IPrecompileWrapper) (*types.Transaction, error) {
	return _PrecompileTest.Contract.Extract(&_PrecompileTest.TransactOpts, x)
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

// NeitherViewNorPure is a paid mutator transaction binding the contract method 0xa7263000.
//
// Solidity: function NeitherViewNorPure() returns()
func (_PrecompileTest *PrecompileTestTransactor) NeitherViewNorPure(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "NeitherViewNorPure")
}

// NeitherViewNorPure is a paid mutator transaction binding the contract method 0xa7263000.
//
// Solidity: function NeitherViewNorPure() returns()
func (_PrecompileTest *PrecompileTestSession) NeitherViewNorPure() (*types.Transaction, error) {
	return _PrecompileTest.Contract.NeitherViewNorPure(&_PrecompileTest.TransactOpts)
}

// NeitherViewNorPure is a paid mutator transaction binding the contract method 0xa7263000.
//
// Solidity: function NeitherViewNorPure() returns()
func (_PrecompileTest *PrecompileTestTransactorSession) NeitherViewNorPure() (*types.Transaction, error) {
	return _PrecompileTest.Contract.NeitherViewNorPure(&_PrecompileTest.TransactOpts)
}

// Pure is a paid mutator transaction binding the contract method 0xf5ad2fb8.
//
// Solidity: function Pure() returns()
func (_PrecompileTest *PrecompileTestTransactor) Pure(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "Pure")
}

// Pure is a paid mutator transaction binding the contract method 0xf5ad2fb8.
//
// Solidity: function Pure() returns()
func (_PrecompileTest *PrecompileTestSession) Pure() (*types.Transaction, error) {
	return _PrecompileTest.Contract.Pure(&_PrecompileTest.TransactOpts)
}

// Pure is a paid mutator transaction binding the contract method 0xf5ad2fb8.
//
// Solidity: function Pure() returns()
func (_PrecompileTest *PrecompileTestTransactorSession) Pure() (*types.Transaction, error) {
	return _PrecompileTest.Contract.Pure(&_PrecompileTest.TransactOpts)
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

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_PrecompileTest *PrecompileTestTransactor) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "Transfer")
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_PrecompileTest *PrecompileTestSession) Transfer() (*types.Transaction, error) {
	return _PrecompileTest.Contract.Transfer(&_PrecompileTest.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_PrecompileTest *PrecompileTestTransactorSession) Transfer() (*types.Transaction, error) {
	return _PrecompileTest.Contract.Transfer(&_PrecompileTest.TransactOpts)
}

// View is a paid mutator transaction binding the contract method 0x1686f265.
//
// Solidity: function View() returns()
func (_PrecompileTest *PrecompileTestTransactor) View(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompileTest.contract.Transact(opts, "View")
}

// View is a paid mutator transaction binding the contract method 0x1686f265.
//
// Solidity: function View() returns()
func (_PrecompileTest *PrecompileTestSession) View() (*types.Transaction, error) {
	return _PrecompileTest.Contract.View(&_PrecompileTest.TransactOpts)
}

// View is a paid mutator transaction binding the contract method 0x1686f265.
//
// Solidity: function View() returns()
func (_PrecompileTest *PrecompileTestTransactorSession) View() (*types.Transaction, error) {
	return _PrecompileTest.Contract.View(&_PrecompileTest.TransactOpts)
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
