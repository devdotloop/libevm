// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testprecompile

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

// TestSuiteMetaData contains all meta data concerning the TestSuite contract.
var TestSuiteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPrecompile\",\"name\":\"_precompile\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nonPayableErrorMsg\",\"type\":\"string\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"func\",\"type\":\"string\"}],\"name\":\"Called\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"Echo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"x\",\"type\":\"string\"}],\"name\":\"Echo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"EchoingFallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"val\",\"type\":\"int256\"}],\"internalType\":\"structIPrecompile.Wrapper\",\"name\":\"x\",\"type\":\"tuple\"}],\"name\":\"Extract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"bytes2\",\"name\":\"y\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"z\",\"type\":\"address\"}],\"name\":\"HashPacked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NeitherViewNorPure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Pure\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"RevertWith\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Self\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"View\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405260405161233c38038061233c83398181016040528101906100259190610227565b8173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1681525050805f90816100679190610491565b505050610560565b5f604051905090565b5f5ffd5b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100a982610080565b9050919050565b5f6100ba8261009f565b9050919050565b6100ca816100b0565b81146100d4575f5ffd5b50565b5f815190506100e5816100c1565b92915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610139826100f3565b810181811067ffffffffffffffff8211171561015857610157610103565b5b80604052505050565b5f61016a61006f565b90506101768282610130565b919050565b5f67ffffffffffffffff82111561019557610194610103565b5b61019e826100f3565b9050602081019050919050565b8281835e5f83830152505050565b5f6101cb6101c68461017b565b610161565b9050828152602081018484840111156101e7576101e66100ef565b5b6101f28482856101ab565b509392505050565b5f82601f83011261020e5761020d6100eb565b5b815161021e8482602086016101b9565b91505092915050565b5f5f6040838503121561023d5761023c610078565b5b5f61024a858286016100d7565b925050602083015167ffffffffffffffff81111561026b5761026a61007c565b5b610277858286016101fa565b9150509250929050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806102cf57607f821691505b6020821081036102e2576102e161028b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026103447fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610309565b61034e8683610309565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61039261038d61038884610366565b61036f565b610366565b9050919050565b5f819050919050565b6103ab83610378565b6103bf6103b782610399565b848454610315565b825550505050565b5f5f905090565b6103d66103c7565b6103e18184846103a2565b505050565b5b81811015610404576103f95f826103ce565b6001810190506103e7565b5050565b601f8211156104495761041a816102e8565b610423846102fa565b81016020851015610432578190505b61044661043e856102fa565b8301826103e6565b50505b505050565b5f82821c905092915050565b5f6104695f198460080261044e565b1980831691505092915050565b5f610481838361045a565b9150826002028217905092915050565b61049a82610281565b67ffffffffffffffff8111156104b3576104b2610103565b5b6104bd82546102b8565b6104c8828285610408565b5f60209050601f8311600181146104f9575f84156104e7578287015190505b6104f18582610476565b865550610558565b601f198416610507866102e8565b5f5b8281101561052e57848901518255600182019150602085019450602081019050610509565b8683101561054b5784890151610547601f89168261045a565b8355505b6001600288020188555050505b505050505050565b608051611d7e6105be5f395f81816101d2015281816102b10152818161041b0152818161057f015281816106b9015281816107a7015281816107de015281816108f1015281816109fc01528181610b4e0152610d020152611d7e5ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c8063ad8108a41161006f578063ad8108a414610101578063b4a818081461011d578063c62c692f14610139578063d7cc1f3714610143578063db84d7c01461015f578063f5ad2fb81461017b576100a7565b80631686f265146100ab57806334d6d9be146100b5578063406dade3146100d1578063a7263000146100db578063a93cbd97146100e5575b5f5ffd5b6100b3610185565b005b6100cf60048036038101906100ca9190610e1e565b6101cf565b005b6100d96102ae565b005b6100e36103cd565b005b6100ff60048036038101906100fa9190610f85565b610417565b005b61011b6004803603810190610116919061103c565b61057d565b005b61013760048036038101906101329190610f85565b61065f565b005b6101416107a5565b005b61015d60048036038101906101589190611116565b6108c4565b005b61017960048036038101906101749190611204565b6109d3565b005b610183610b01565b005b610198631686f26560e01b60015f610b4a565b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516101c5906112a5565b60405180910390a1565b807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166334d6d9be836040518263ffffffff1660e01b815260040161022991906112d2565b602060405180830381865afa158015610244573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061026891906112ff565b146102765761027561132a565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516102a3906113a1565b60405180910390a150565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16631b7fb4b0602a6040518263ffffffff1660e01b815260040160206040518083038185885af115801561031b573d5f5f3e3d5ffd5b50505050506040513d601f19601f8201168201806040525081019061034091906112ff565b9050602a81146103535761035261132a565b5b61036b5f60405180602001604052805f815250610c90565b61039560015f60405160200161038191906114b8565b604051602081830303815290604052610c90565b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516103c290611518565b60405180910390a150565b6103e063a726300060e01b600180610b4a565b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a60405161040d90611580565b60405180910390a1565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a93cbd9760e01b8460405160240161046991906115fe565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516104d39190611658565b5f604051808303815f865af19150503d805f811461050c576040519150601f19603f3d011682016040523d82523d5f602084013e610511565b606091505b509150915081156105255761052461132a565b5b82805190602001208180519060200120146105435761054261132a565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610570906116b8565b60405180910390a1505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663ad8108a4826040518263ffffffff1660e01b81526004016105d691906116ff565b602060405180830381865afa1580156105f1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610615919061172c565b815f0151146106275761062661132a565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610654906117a1565b60405180910390a150565b5f5f8260405160240161067291906115fe565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16836040516106fc9190611658565b5f60405180830381855afa9150503d805f8114610734576040519150601f19603f3d011682016040523d82523d5f602084013e610739565b606091505b50915091508161074c5761074b61132a565b5b828051906020012081805190602001201461076a5761076961132a565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a60405161079790611809565b60405180910390a150505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c62c692f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610845573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610869919061183b565b73ffffffffffffffffffffffffffffffffffffffff161461088d5761088c61132a565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516108ba906118b0565b60405180910390a1565b8282826040516020016108d993929190611953565b604051602081830303815290604052805190602001207f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d7cc1f378585856040518463ffffffff1660e01b815260040161094c939291906119ad565b602060405180830381865afa158015610967573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061098b9190611a15565b146109995761099861132a565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a6040516109c690611a8a565b60405180910390a1505050565b806040516020016109e49190611ae2565b604051602081830303815290604052805190602001207f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663db84d7c0836040518263ffffffff1660e01b8152600401610a539190611b30565b5f60405180830381865afa158015610a6d573d5f5f3e3d5ffd5b505050506040513d5f823e3d601f19601f82011682018060405250810190610a959190611bbe565b604051602001610aa59190611ae2565b6040516020818303038152906040528051906020012014610ac957610ac861132a565b5b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610af690611c4f565b60405180910390a150565b610b1363f5ad2fb860e01b5f5f610b4a565b7f3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a604051610b4090611cb7565b60405180910390a1565b5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1685604051602401604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610bf49190611658565b5f604051808303815f865af19150503d805f8114610c2d576040519150601f19603f3d011682016040523d82523d5f602084013e610c32565b606091505b509150915081610c4557610c4461132a565b5b5f5f82806020019051810190610c5b9190611d0a565b9150915085151582151514610c7357610c7261132a565b5b84151581151514610c8757610c8661132a565b5b50505050505050565b5f636fb1b0e960e01b604051602401604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505090505f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168584604051610d469190611658565b5f6040518083038185875af1925050503d805f8114610d80576040519150601f19603f3d011682016040523d82523d5f602084013e610d85565b606091505b50915091505f845103610da55781610da057610d9f61132a565b5b610dd3565b8115610db457610db361132a565b5b8380519060200120818051906020012014610dd257610dd161132a565b5b5b5050505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610dfd81610deb565b8114610e07575f5ffd5b50565b5f81359050610e1881610df4565b92915050565b5f60208284031215610e3357610e32610de3565b5b5f610e4084828501610e0a565b91505092915050565b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e9782610e51565b810181811067ffffffffffffffff82111715610eb657610eb5610e61565b5b80604052505050565b5f610ec8610dda565b9050610ed48282610e8e565b919050565b5f67ffffffffffffffff821115610ef357610ef2610e61565b5b610efc82610e51565b9050602081019050919050565b828183375f83830152505050565b5f610f29610f2484610ed9565b610ebf565b905082815260208101848484011115610f4557610f44610e4d565b5b610f50848285610f09565b509392505050565b5f82601f830112610f6c57610f6b610e49565b5b8135610f7c848260208601610f17565b91505092915050565b5f60208284031215610f9a57610f99610de3565b5b5f82013567ffffffffffffffff811115610fb757610fb6610de7565b5b610fc384828501610f58565b91505092915050565b5f5ffd5b5f819050919050565b610fe281610fd0565b8114610fec575f5ffd5b50565b5f81359050610ffd81610fd9565b92915050565b5f6020828403121561101857611017610fcc565b5b6110226020610ebf565b90505f61103184828501610fef565b5f8301525092915050565b5f6020828403121561105157611050610de3565b5b5f61105e84828501611003565b91505092915050565b5f7fffff00000000000000000000000000000000000000000000000000000000000082169050919050565b61109b81611067565b81146110a5575f5ffd5b50565b5f813590506110b681611092565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6110e5826110bc565b9050919050565b6110f5816110db565b81146110ff575f5ffd5b50565b5f81359050611110816110ec565b92915050565b5f5f5f6060848603121561112d5761112c610de3565b5b5f61113a86828701610e0a565b935050602061114b868287016110a8565b925050604061115c86828701611102565b9150509250925092565b5f67ffffffffffffffff8211156111805761117f610e61565b5b61118982610e51565b9050602081019050919050565b5f6111a86111a384611166565b610ebf565b9050828152602081018484840111156111c4576111c3610e4d565b5b6111cf848285610f09565b509392505050565b5f82601f8301126111eb576111ea610e49565b5b81356111fb848260208601611196565b91505092915050565b5f6020828403121561121957611218610de3565b5b5f82013567ffffffffffffffff81111561123657611235610de7565b5b611242848285016111d7565b91505092915050565b5f82825260208201905092915050565b7f56696577282900000000000000000000000000000000000000000000000000005f82015250565b5f61128f60068361124b565b915061129a8261125b565b602082019050919050565b5f6020820190508181035f8301526112bc81611283565b9050919050565b6112cc81610deb565b82525050565b5f6020820190506112e55f8301846112c3565b92915050565b5f815190506112f981610df4565b92915050565b5f6020828403121561131457611313610de3565b5b5f611321848285016112eb565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffd5b7f4563686f2875696e7432353629000000000000000000000000000000000000005f82015250565b5f61138b600d8361124b565b915061139682611357565b602082019050919050565b5f6020820190508181035f8301526113b88161137f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061140357607f821691505b602082108103611416576114156113bf565b5b50919050565b5f81905092915050565b5f819050815f5260205f209050919050565b5f8154611444816113ec565b61144e818661141c565b9450600182165f8114611468576001811461147d576114af565b60ff19831686528115158202860193506114af565b61148685611426565b5f5b838110156114a757815481890152600182019150602081019050611488565b838801955050505b50505092915050565b5f6114c38284611438565b915081905092915050565b7f5472616e736665722829000000000000000000000000000000000000000000005f82015250565b5f611502600a8361124b565b915061150d826114ce565b602082019050919050565b5f6020820190508181035f83015261152f816114f6565b9050919050565b7f4e656974686572566965774e6f725075726528290000000000000000000000005f82015250565b5f61156a60148361124b565b915061157582611536565b602082019050919050565b5f6020820190508181035f8301526115978161155e565b9050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6115d08261159e565b6115da81856115a8565b93506115ea8185602086016115b8565b6115f381610e51565b840191505092915050565b5f6020820190508181035f83015261161681846115c6565b905092915050565b5f81905092915050565b5f6116328261159e565b61163c818561161e565b935061164c8185602086016115b8565b80840191505092915050565b5f6116638284611628565b915081905092915050565b7f52657665727457697468282e2e2e2900000000000000000000000000000000005f82015250565b5f6116a2600f8361124b565b91506116ad8261166e565b602082019050919050565b5f6020820190508181035f8301526116cf81611696565b9050919050565b6116df81610fd0565b82525050565b602082015f8201516116f95f8501826116d6565b50505050565b5f6020820190506117125f8301846116e5565b92915050565b5f8151905061172681610fd9565b92915050565b5f6020828403121561174157611740610de3565b5b5f61174e84828501611718565b91505092915050565b7f45787472616374282e2e2e2900000000000000000000000000000000000000005f82015250565b5f61178b600c8361124b565b915061179682611757565b602082019050919050565b5f6020820190508181035f8301526117b88161177f565b9050919050565b7f4563686f696e6746616c6c6261636b282e2e2e290000000000000000000000005f82015250565b5f6117f360148361124b565b91506117fe826117bf565b602082019050919050565b5f6020820190508181035f830152611820816117e7565b9050919050565b5f81519050611835816110ec565b92915050565b5f602082840312156118505761184f610de3565b5b5f61185d84828501611827565b91505092915050565b7f53656c66282900000000000000000000000000000000000000000000000000005f82015250565b5f61189a60068361124b565b91506118a582611866565b602082019050919050565b5f6020820190508181035f8301526118c78161188e565b9050919050565b5f819050919050565b6118e86118e382610deb565b6118ce565b82525050565b5f819050919050565b61190861190382611067565b6118ee565b82525050565b5f8160601b9050919050565b5f6119248261190e565b9050919050565b5f6119358261191a565b9050919050565b61194d611948826110db565b61192b565b82525050565b5f61195e82866118d7565b60208201915061196e82856118f7565b60028201915061197e828461193c565b601482019150819050949350505050565b61199881611067565b82525050565b6119a7816110db565b82525050565b5f6060820190506119c05f8301866112c3565b6119cd602083018561198f565b6119da604083018461199e565b949350505050565b5f819050919050565b6119f4816119e2565b81146119fe575f5ffd5b50565b5f81519050611a0f816119eb565b92915050565b5f60208284031215611a2a57611a29610de3565b5b5f611a3784828501611a01565b91505092915050565b7f486173685061636b6564282e2e2e2900000000000000000000000000000000005f82015250565b5f611a74600f8361124b565b9150611a7f82611a40565b602082019050919050565b5f6020820190508181035f830152611aa181611a68565b9050919050565b5f81519050919050565b5f611abc82611aa8565b611ac6818561141c565b9350611ad68185602086016115b8565b80840191505092915050565b5f611aed8284611ab2565b915081905092915050565b5f611b0282611aa8565b611b0c818561124b565b9350611b1c8185602086016115b8565b611b2581610e51565b840191505092915050565b5f6020820190508181035f830152611b488184611af8565b905092915050565b5f611b62611b5d84611166565b610ebf565b905082815260208101848484011115611b7e57611b7d610e4d565b5b611b898482856115b8565b509392505050565b5f82601f830112611ba557611ba4610e49565b5b8151611bb5848260208601611b50565b91505092915050565b5f60208284031215611bd357611bd2610de3565b5b5f82015167ffffffffffffffff811115611bf057611bef610de7565b5b611bfc84828501611b91565b91505092915050565b7f4563686f28737472696e672900000000000000000000000000000000000000005f82015250565b5f611c39600c8361124b565b9150611c4482611c05565b602082019050919050565b5f6020820190508181035f830152611c6681611c2d565b9050919050565b7f50757265282900000000000000000000000000000000000000000000000000005f82015250565b5f611ca160068361124b565b9150611cac82611c6d565b602082019050919050565b5f6020820190508181035f830152611cce81611c95565b9050919050565b5f8115159050919050565b611ce981611cd5565b8114611cf3575f5ffd5b50565b5f81519050611d0481611ce0565b92915050565b5f5f60408385031215611d2057611d1f610de3565b5b5f611d2d85828601611cf6565b9250506020611d3e85828601611cf6565b915050925092905056fea26469706673582212203f8f80491d9ec37e8b661dc7940f888a634de87866cea3d27267b7366c437dc964736f6c634300081c0033",
}

// TestSuiteABI is the input ABI used to generate the binding from.
// Deprecated: Use TestSuiteMetaData.ABI instead.
var TestSuiteABI = TestSuiteMetaData.ABI

// TestSuiteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestSuiteMetaData.Bin instead.
var TestSuiteBin = TestSuiteMetaData.Bin

// DeployTestSuite deploys a new Ethereum contract, binding an instance of TestSuite to it.
func DeployTestSuite(auth *bind.TransactOpts, backend bind.ContractBackend, _precompile common.Address, nonPayableErrorMsg string) (common.Address, *types.Transaction, *TestSuite, error) {
	parsed, err := TestSuiteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestSuiteBin), backend, _precompile, nonPayableErrorMsg)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestSuite{TestSuiteCaller: TestSuiteCaller{contract: contract}, TestSuiteTransactor: TestSuiteTransactor{contract: contract}, TestSuiteFilterer: TestSuiteFilterer{contract: contract}}, nil
}

// TestSuite is an auto generated Go binding around an Ethereum contract.
type TestSuite struct {
	TestSuiteCaller     // Read-only binding to the contract
	TestSuiteTransactor // Write-only binding to the contract
	TestSuiteFilterer   // Log filterer for contract events
}

// TestSuiteCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestSuiteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSuiteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestSuiteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSuiteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestSuiteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSuiteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSuiteSession struct {
	Contract     *TestSuite        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestSuiteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestSuiteCallerSession struct {
	Contract *TestSuiteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestSuiteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestSuiteTransactorSession struct {
	Contract     *TestSuiteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestSuiteRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestSuiteRaw struct {
	Contract *TestSuite // Generic contract binding to access the raw methods on
}

// TestSuiteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestSuiteCallerRaw struct {
	Contract *TestSuiteCaller // Generic read-only contract binding to access the raw methods on
}

// TestSuiteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestSuiteTransactorRaw struct {
	Contract *TestSuiteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestSuite creates a new instance of TestSuite, bound to a specific deployed contract.
func NewTestSuite(address common.Address, backend bind.ContractBackend) (*TestSuite, error) {
	contract, err := bindTestSuite(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestSuite{TestSuiteCaller: TestSuiteCaller{contract: contract}, TestSuiteTransactor: TestSuiteTransactor{contract: contract}, TestSuiteFilterer: TestSuiteFilterer{contract: contract}}, nil
}

// NewTestSuiteCaller creates a new read-only instance of TestSuite, bound to a specific deployed contract.
func NewTestSuiteCaller(address common.Address, caller bind.ContractCaller) (*TestSuiteCaller, error) {
	contract, err := bindTestSuite(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestSuiteCaller{contract: contract}, nil
}

// NewTestSuiteTransactor creates a new write-only instance of TestSuite, bound to a specific deployed contract.
func NewTestSuiteTransactor(address common.Address, transactor bind.ContractTransactor) (*TestSuiteTransactor, error) {
	contract, err := bindTestSuite(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestSuiteTransactor{contract: contract}, nil
}

// NewTestSuiteFilterer creates a new log filterer instance of TestSuite, bound to a specific deployed contract.
func NewTestSuiteFilterer(address common.Address, filterer bind.ContractFilterer) (*TestSuiteFilterer, error) {
	contract, err := bindTestSuite(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestSuiteFilterer{contract: contract}, nil
}

// bindTestSuite binds a generic wrapper to an already deployed contract.
func bindTestSuite(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestSuiteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSuite *TestSuiteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSuite.Contract.TestSuiteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSuite *TestSuiteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSuite.Contract.TestSuiteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSuite *TestSuiteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSuite.Contract.TestSuiteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSuite *TestSuiteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSuite.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSuite *TestSuiteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSuite.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSuite *TestSuiteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSuite.Contract.contract.Transact(opts, method, params...)
}

// Echo is a paid mutator transaction binding the contract method 0x34d6d9be.
//
// Solidity: function Echo(uint256 x) returns()
func (_TestSuite *TestSuiteTransactor) Echo(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "Echo", x)
}

// Echo is a paid mutator transaction binding the contract method 0x34d6d9be.
//
// Solidity: function Echo(uint256 x) returns()
func (_TestSuite *TestSuiteSession) Echo(x *big.Int) (*types.Transaction, error) {
	return _TestSuite.Contract.Echo(&_TestSuite.TransactOpts, x)
}

// Echo is a paid mutator transaction binding the contract method 0x34d6d9be.
//
// Solidity: function Echo(uint256 x) returns()
func (_TestSuite *TestSuiteTransactorSession) Echo(x *big.Int) (*types.Transaction, error) {
	return _TestSuite.Contract.Echo(&_TestSuite.TransactOpts, x)
}

// Echo0 is a paid mutator transaction binding the contract method 0xdb84d7c0.
//
// Solidity: function Echo(string x) returns()
func (_TestSuite *TestSuiteTransactor) Echo0(opts *bind.TransactOpts, x string) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "Echo0", x)
}

// Echo0 is a paid mutator transaction binding the contract method 0xdb84d7c0.
//
// Solidity: function Echo(string x) returns()
func (_TestSuite *TestSuiteSession) Echo0(x string) (*types.Transaction, error) {
	return _TestSuite.Contract.Echo0(&_TestSuite.TransactOpts, x)
}

// Echo0 is a paid mutator transaction binding the contract method 0xdb84d7c0.
//
// Solidity: function Echo(string x) returns()
func (_TestSuite *TestSuiteTransactorSession) Echo0(x string) (*types.Transaction, error) {
	return _TestSuite.Contract.Echo0(&_TestSuite.TransactOpts, x)
}

// EchoingFallback is a paid mutator transaction binding the contract method 0xb4a81808.
//
// Solidity: function EchoingFallback(bytes input) returns()
func (_TestSuite *TestSuiteTransactor) EchoingFallback(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "EchoingFallback", input)
}

// EchoingFallback is a paid mutator transaction binding the contract method 0xb4a81808.
//
// Solidity: function EchoingFallback(bytes input) returns()
func (_TestSuite *TestSuiteSession) EchoingFallback(input []byte) (*types.Transaction, error) {
	return _TestSuite.Contract.EchoingFallback(&_TestSuite.TransactOpts, input)
}

// EchoingFallback is a paid mutator transaction binding the contract method 0xb4a81808.
//
// Solidity: function EchoingFallback(bytes input) returns()
func (_TestSuite *TestSuiteTransactorSession) EchoingFallback(input []byte) (*types.Transaction, error) {
	return _TestSuite.Contract.EchoingFallback(&_TestSuite.TransactOpts, input)
}

// Extract is a paid mutator transaction binding the contract method 0xad8108a4.
//
// Solidity: function Extract((int256) x) returns()
func (_TestSuite *TestSuiteTransactor) Extract(opts *bind.TransactOpts, x IPrecompileWrapper) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "Extract", x)
}

// Extract is a paid mutator transaction binding the contract method 0xad8108a4.
//
// Solidity: function Extract((int256) x) returns()
func (_TestSuite *TestSuiteSession) Extract(x IPrecompileWrapper) (*types.Transaction, error) {
	return _TestSuite.Contract.Extract(&_TestSuite.TransactOpts, x)
}

// Extract is a paid mutator transaction binding the contract method 0xad8108a4.
//
// Solidity: function Extract((int256) x) returns()
func (_TestSuite *TestSuiteTransactorSession) Extract(x IPrecompileWrapper) (*types.Transaction, error) {
	return _TestSuite.Contract.Extract(&_TestSuite.TransactOpts, x)
}

// HashPacked is a paid mutator transaction binding the contract method 0xd7cc1f37.
//
// Solidity: function HashPacked(uint256 x, bytes2 y, address z) returns()
func (_TestSuite *TestSuiteTransactor) HashPacked(opts *bind.TransactOpts, x *big.Int, y [2]byte, z common.Address) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "HashPacked", x, y, z)
}

// HashPacked is a paid mutator transaction binding the contract method 0xd7cc1f37.
//
// Solidity: function HashPacked(uint256 x, bytes2 y, address z) returns()
func (_TestSuite *TestSuiteSession) HashPacked(x *big.Int, y [2]byte, z common.Address) (*types.Transaction, error) {
	return _TestSuite.Contract.HashPacked(&_TestSuite.TransactOpts, x, y, z)
}

// HashPacked is a paid mutator transaction binding the contract method 0xd7cc1f37.
//
// Solidity: function HashPacked(uint256 x, bytes2 y, address z) returns()
func (_TestSuite *TestSuiteTransactorSession) HashPacked(x *big.Int, y [2]byte, z common.Address) (*types.Transaction, error) {
	return _TestSuite.Contract.HashPacked(&_TestSuite.TransactOpts, x, y, z)
}

// NeitherViewNorPure is a paid mutator transaction binding the contract method 0xa7263000.
//
// Solidity: function NeitherViewNorPure() returns()
func (_TestSuite *TestSuiteTransactor) NeitherViewNorPure(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "NeitherViewNorPure")
}

// NeitherViewNorPure is a paid mutator transaction binding the contract method 0xa7263000.
//
// Solidity: function NeitherViewNorPure() returns()
func (_TestSuite *TestSuiteSession) NeitherViewNorPure() (*types.Transaction, error) {
	return _TestSuite.Contract.NeitherViewNorPure(&_TestSuite.TransactOpts)
}

// NeitherViewNorPure is a paid mutator transaction binding the contract method 0xa7263000.
//
// Solidity: function NeitherViewNorPure() returns()
func (_TestSuite *TestSuiteTransactorSession) NeitherViewNorPure() (*types.Transaction, error) {
	return _TestSuite.Contract.NeitherViewNorPure(&_TestSuite.TransactOpts)
}

// Pure is a paid mutator transaction binding the contract method 0xf5ad2fb8.
//
// Solidity: function Pure() returns()
func (_TestSuite *TestSuiteTransactor) Pure(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "Pure")
}

// Pure is a paid mutator transaction binding the contract method 0xf5ad2fb8.
//
// Solidity: function Pure() returns()
func (_TestSuite *TestSuiteSession) Pure() (*types.Transaction, error) {
	return _TestSuite.Contract.Pure(&_TestSuite.TransactOpts)
}

// Pure is a paid mutator transaction binding the contract method 0xf5ad2fb8.
//
// Solidity: function Pure() returns()
func (_TestSuite *TestSuiteTransactorSession) Pure() (*types.Transaction, error) {
	return _TestSuite.Contract.Pure(&_TestSuite.TransactOpts)
}

// RevertWith is a paid mutator transaction binding the contract method 0xa93cbd97.
//
// Solidity: function RevertWith(bytes err) returns()
func (_TestSuite *TestSuiteTransactor) RevertWith(opts *bind.TransactOpts, err []byte) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "RevertWith", err)
}

// RevertWith is a paid mutator transaction binding the contract method 0xa93cbd97.
//
// Solidity: function RevertWith(bytes err) returns()
func (_TestSuite *TestSuiteSession) RevertWith(err []byte) (*types.Transaction, error) {
	return _TestSuite.Contract.RevertWith(&_TestSuite.TransactOpts, err)
}

// RevertWith is a paid mutator transaction binding the contract method 0xa93cbd97.
//
// Solidity: function RevertWith(bytes err) returns()
func (_TestSuite *TestSuiteTransactorSession) RevertWith(err []byte) (*types.Transaction, error) {
	return _TestSuite.Contract.RevertWith(&_TestSuite.TransactOpts, err)
}

// Self is a paid mutator transaction binding the contract method 0xc62c692f.
//
// Solidity: function Self() returns()
func (_TestSuite *TestSuiteTransactor) Self(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "Self")
}

// Self is a paid mutator transaction binding the contract method 0xc62c692f.
//
// Solidity: function Self() returns()
func (_TestSuite *TestSuiteSession) Self() (*types.Transaction, error) {
	return _TestSuite.Contract.Self(&_TestSuite.TransactOpts)
}

// Self is a paid mutator transaction binding the contract method 0xc62c692f.
//
// Solidity: function Self() returns()
func (_TestSuite *TestSuiteTransactorSession) Self() (*types.Transaction, error) {
	return _TestSuite.Contract.Self(&_TestSuite.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_TestSuite *TestSuiteTransactor) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "Transfer")
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_TestSuite *TestSuiteSession) Transfer() (*types.Transaction, error) {
	return _TestSuite.Contract.Transfer(&_TestSuite.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_TestSuite *TestSuiteTransactorSession) Transfer() (*types.Transaction, error) {
	return _TestSuite.Contract.Transfer(&_TestSuite.TransactOpts)
}

// View is a paid mutator transaction binding the contract method 0x1686f265.
//
// Solidity: function View() returns()
func (_TestSuite *TestSuiteTransactor) View(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSuite.contract.Transact(opts, "View")
}

// View is a paid mutator transaction binding the contract method 0x1686f265.
//
// Solidity: function View() returns()
func (_TestSuite *TestSuiteSession) View() (*types.Transaction, error) {
	return _TestSuite.Contract.View(&_TestSuite.TransactOpts)
}

// View is a paid mutator transaction binding the contract method 0x1686f265.
//
// Solidity: function View() returns()
func (_TestSuite *TestSuiteTransactorSession) View() (*types.Transaction, error) {
	return _TestSuite.Contract.View(&_TestSuite.TransactOpts)
}

// TestSuiteCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the TestSuite contract.
type TestSuiteCalledIterator struct {
	Event *TestSuiteCalled // Event containing the contract specifics and raw log

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
func (it *TestSuiteCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSuiteCalled)
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
		it.Event = new(TestSuiteCalled)
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
func (it *TestSuiteCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSuiteCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSuiteCalled represents a Called event raised by the TestSuite contract.
type TestSuiteCalled struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0x3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a.
//
// Solidity: event Called(string func)
func (_TestSuite *TestSuiteFilterer) FilterCalled(opts *bind.FilterOpts) (*TestSuiteCalledIterator, error) {

	logs, sub, err := _TestSuite.contract.FilterLogs(opts, "Called")
	if err != nil {
		return nil, err
	}
	return &TestSuiteCalledIterator{contract: _TestSuite.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x3fd5724095fcb65d3b7e6d79ac130db7e9dbf891d59bb0da292d471bc40d564a.
//
// Solidity: event Called(string func)
func (_TestSuite *TestSuiteFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *TestSuiteCalled) (event.Subscription, error) {

	logs, sub, err := _TestSuite.contract.WatchLogs(opts, "Called")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSuiteCalled)
				if err := _TestSuite.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_TestSuite *TestSuiteFilterer) ParseCalled(log types.Log) (*TestSuiteCalled, error) {
	event := new(TestSuiteCalled)
	if err := _TestSuite.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
