// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package feeds_consumer

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"
)

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

type KeystoneFeedsPermissionHandlerPermission struct {
	Forwarder     common.Address
	WorkflowName  [10]byte
	ReportName    [2]byte
	WorkflowOwner common.Address
	IsAllowed     bool
}

var KeystoneFeedsConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"}],\"name\":\"ReportForwarderUnauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint224\",\"name\":\"price\",\"type\":\"uint224\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"name\":\"FeedReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rawReport\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"reportId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission\",\"name\":\"permission\",\"type\":\"tuple\"}],\"name\":\"ReportPermissionSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"feedId\",\"type\":\"bytes32\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"metadata\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rawReport\",\"type\":\"bytes\"}],\"name\":\"onReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"},{\"internalType\":\"bytes10\",\"name\":\"workflowName\",\"type\":\"bytes10\"},{\"internalType\":\"bytes2\",\"name\":\"reportName\",\"type\":\"bytes2\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"internalType\":\"structKeystoneFeedsPermissionHandler.Permission[]\",\"name\":\"permissions\",\"type\":\"tuple[]\"}],\"name\":\"setReportPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033806000816100675760405162461bcd60e51b815260206004820152601860248201527f43616e6e6f7420736574206f776e657220746f207a65726f000000000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0384811691909117909155811615610097576100978161009f565b505050610148565b336001600160a01b038216036100f75760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161005e565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b610ce9806101576000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063805f213211610050578063805f21321461014e5780638da5cb5b14610161578063f2fde38b1461018957600080fd5b806331d98b3f1461007757806341ed29e71461013157806379ba509714610146575b600080fd5b6100f3610085366004610856565b6000908152600360209081526040918290208251808401909352547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168084527c010000000000000000000000000000000000000000000000000000000090910463ffffffff169290910182905291565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff909316835263ffffffff9091166020830152015b60405180910390f35b61014461013f366004610986565b61019c565b005b6101446101de565b61014461015c366004610b21565b6102e0565b60005460405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610128565b610144610197366004610b8d565b6104f8565b6101a461050c565b60005b81518110156101da576101d28282815181106101c5576101c5610baf565b602002602001015161058f565b6001016101a7565b5050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610264576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b7f245e8e0654076d706bf5dea6eab892c9d7788b4aefa7f4923e837e43d200959b848484846040516103159493929190610c27565b60405180910390a16104f2565b81518110156104ec57604051806040016040528083838151811061034857610348610baf565b6020026020010151602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16815260200183838151811061038957610389610baf565b60200260200101516040015163ffffffff16815250600360008484815181106103b4576103b4610baf565b602090810291909101810151518252818101929092526040016000208251929091015163ffffffff167c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff909216919091179055815182908290811061043657610436610baf565b6020026020010151600001517f2c30f5cb3caf4239d0f994ce539d7ef24817fa550169c388e3a110f02e40197d83838151811061047557610475610baf565b60200260200101516020015184848151811061049357610493610baf565b6020026020010151604001516040516104dc9291907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff92909216825263ffffffff16602082015260400190565b60405180910390a2600101610322565b50505050505b50505050565b61050061050c565b61050981610761565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461058d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e657200000000000000000000604482015260640161025b565b565b600061064882600001518360600151846020015185604001516040805173ffffffffffffffffffffffffffffffffffffffff80871660208301528516918101919091527fffffffffffffffffffff00000000000000000000000000000000000000000000831660608201527fffff0000000000000000000000000000000000000000000000000000000000008216608082015260009060a001604051602081830303815290604052805190602001209050949350505050565b60808301516000828152600260205260409081902080549215157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00909316929092179091555190915081907f32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a390610755908590600060a08201905073ffffffffffffffffffffffffffffffffffffffff8084511683527fffffffffffffffffffff0000000000000000000000000000000000000000000060208501511660208401527fffff00000000000000000000000000000000000000000000000000000000000060408501511660408401528060608501511660608401525060808301511515608083015292915050565b60405180910390a25050565b3373ffffffffffffffffffffffffffffffffffffffff8216036107e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c66000000000000000000604482015260640161025b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b60006020828403121561086857600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff811182821017156108c1576108c161086f565b60405290565b6040516060810167ffffffffffffffff811182821017156108c1576108c161086f565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109315761093161086f565b604052919050565b600067ffffffffffffffff8211156109535761095361086f565b5060051b60200190565b803573ffffffffffffffffffffffffffffffffffffffff8116811461098157600080fd5b919050565b6000602080838503121561099957600080fd5b823567ffffffffffffffff8111156109b057600080fd5b8301601f810185136109c157600080fd5b80356109d46109cf82610939565b6108ea565b81815260a091820283018401918482019190888411156109f357600080fd5b938501935b83851015610acc5780858a031215610a105760008081fd5b610a1861089e565b610a218661095d565b8152868601357fffffffffffffffffffff0000000000000000000000000000000000000000000081168114610a565760008081fd5b818801526040868101357fffff00000000000000000000000000000000000000000000000000000000000081168114610a8f5760008081fd5b908201526060610aa087820161095d565b908201526080868101358015158114610ab95760008081fd5b90820152835293840193918501916109f8565b50979650505050505050565b60008083601f840112610aea57600080fd5b50813567ffffffffffffffff811115610b0257600080fd5b602083019150836020828501011115610b1a57600080fd5b9250929050565b60008060008060408587031215610b3757600080fd5b843567ffffffffffffffff80821115610b4f57600080fd5b610b5b88838901610ad8565b90965094506020870135915080821115610b7457600080fd5b50610b8187828801610ad8565b95989497509550505050565b600060208284031215610b9f57600080fd5b610ba88261095d565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000610c3b604083018688610bde565b8281036020840152610c4e818587610bde565b979650505050505050565b83851015610acc5780858a031215610c715760008081fd5b610c796108c7565b85358152868601357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff81168114610cac5760008081fd5b8188015260408681013563ffffffff81168114610cc95760008081fd5b9082015283529384019391850191610c5956fea164736f6c6343000818000a",
}

var KeystoneFeedsConsumerABI = KeystoneFeedsConsumerMetaData.ABI

var KeystoneFeedsConsumerBin = KeystoneFeedsConsumerMetaData.Bin

func DeployKeystoneFeedsConsumer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KeystoneFeedsConsumer, error) {
	parsed, err := KeystoneFeedsConsumerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KeystoneFeedsConsumerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KeystoneFeedsConsumer{address: address, abi: *parsed, KeystoneFeedsConsumerCaller: KeystoneFeedsConsumerCaller{contract: contract}, KeystoneFeedsConsumerTransactor: KeystoneFeedsConsumerTransactor{contract: contract}, KeystoneFeedsConsumerFilterer: KeystoneFeedsConsumerFilterer{contract: contract}}, nil
}

type KeystoneFeedsConsumer struct {
	address common.Address
	abi     abi.ABI
	KeystoneFeedsConsumerCaller
	KeystoneFeedsConsumerTransactor
	KeystoneFeedsConsumerFilterer
}

type KeystoneFeedsConsumerCaller struct {
	contract *bind.BoundContract
}

type KeystoneFeedsConsumerTransactor struct {
	contract *bind.BoundContract
}

type KeystoneFeedsConsumerFilterer struct {
	contract *bind.BoundContract
}

type KeystoneFeedsConsumerSession struct {
	Contract     *KeystoneFeedsConsumer
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type KeystoneFeedsConsumerCallerSession struct {
	Contract *KeystoneFeedsConsumerCaller
	CallOpts bind.CallOpts
}

type KeystoneFeedsConsumerTransactorSession struct {
	Contract     *KeystoneFeedsConsumerTransactor
	TransactOpts bind.TransactOpts
}

type KeystoneFeedsConsumerRaw struct {
	Contract *KeystoneFeedsConsumer
}

type KeystoneFeedsConsumerCallerRaw struct {
	Contract *KeystoneFeedsConsumerCaller
}

type KeystoneFeedsConsumerTransactorRaw struct {
	Contract *KeystoneFeedsConsumerTransactor
}

func NewKeystoneFeedsConsumer(address common.Address, backend bind.ContractBackend) (*KeystoneFeedsConsumer, error) {
	abi, err := abi.JSON(strings.NewReader(KeystoneFeedsConsumerABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindKeystoneFeedsConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumer{address: address, abi: abi, KeystoneFeedsConsumerCaller: KeystoneFeedsConsumerCaller{contract: contract}, KeystoneFeedsConsumerTransactor: KeystoneFeedsConsumerTransactor{contract: contract}, KeystoneFeedsConsumerFilterer: KeystoneFeedsConsumerFilterer{contract: contract}}, nil
}

func NewKeystoneFeedsConsumerCaller(address common.Address, caller bind.ContractCaller) (*KeystoneFeedsConsumerCaller, error) {
	contract, err := bindKeystoneFeedsConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumerCaller{contract: contract}, nil
}

func NewKeystoneFeedsConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*KeystoneFeedsConsumerTransactor, error) {
	contract, err := bindKeystoneFeedsConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumerTransactor{contract: contract}, nil
}

func NewKeystoneFeedsConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*KeystoneFeedsConsumerFilterer, error) {
	contract, err := bindKeystoneFeedsConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumerFilterer{contract: contract}, nil
}

func bindKeystoneFeedsConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KeystoneFeedsConsumerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeystoneFeedsConsumer.Contract.KeystoneFeedsConsumerCaller.contract.Call(opts, result, method, params...)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.KeystoneFeedsConsumerTransactor.contract.Transfer(opts)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.KeystoneFeedsConsumerTransactor.contract.Transact(opts, method, params...)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KeystoneFeedsConsumer.Contract.contract.Call(opts, result, method, params...)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.contract.Transfer(opts)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.contract.Transact(opts, method, params...)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerCaller) GetPrice(opts *bind.CallOpts, feedId [32]byte) (*big.Int, uint32, error) {
	var out []interface{}
	err := _KeystoneFeedsConsumer.contract.Call(opts, &out, "getPrice", feedId)

	if err != nil {
		return *new(*big.Int), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerSession) GetPrice(feedId [32]byte) (*big.Int, uint32, error) {
	return _KeystoneFeedsConsumer.Contract.GetPrice(&_KeystoneFeedsConsumer.CallOpts, feedId)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerCallerSession) GetPrice(feedId [32]byte) (*big.Int, uint32, error) {
	return _KeystoneFeedsConsumer.Contract.GetPrice(&_KeystoneFeedsConsumer.CallOpts, feedId)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KeystoneFeedsConsumer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerSession) Owner() (common.Address, error) {
	return _KeystoneFeedsConsumer.Contract.Owner(&_KeystoneFeedsConsumer.CallOpts)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerCallerSession) Owner() (common.Address, error) {
	return _KeystoneFeedsConsumer.Contract.Owner(&_KeystoneFeedsConsumer.CallOpts)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.contract.Transact(opts, "acceptOwnership")
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerSession) AcceptOwnership() (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.AcceptOwnership(&_KeystoneFeedsConsumer.TransactOpts)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.AcceptOwnership(&_KeystoneFeedsConsumer.TransactOpts)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactor) OnReport(opts *bind.TransactOpts, metadata []byte, rawReport []byte) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.contract.Transact(opts, "onReport", metadata, rawReport)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerSession) OnReport(metadata []byte, rawReport []byte) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.OnReport(&_KeystoneFeedsConsumer.TransactOpts, metadata, rawReport)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactorSession) OnReport(metadata []byte, rawReport []byte) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.OnReport(&_KeystoneFeedsConsumer.TransactOpts, metadata, rawReport)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactor) SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.contract.Transact(opts, "setReportPermissions", permissions)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerSession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.SetReportPermissions(&_KeystoneFeedsConsumer.TransactOpts, permissions)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactorSession) SetReportPermissions(permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.SetReportPermissions(&_KeystoneFeedsConsumer.TransactOpts, permissions)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.contract.Transact(opts, "transferOwnership", to)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.TransferOwnership(&_KeystoneFeedsConsumer.TransactOpts, to)
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _KeystoneFeedsConsumer.Contract.TransferOwnership(&_KeystoneFeedsConsumer.TransactOpts, to)
}

type KeystoneFeedsConsumerFeedReceivedIterator struct {
	Event *KeystoneFeedsConsumerFeedReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *KeystoneFeedsConsumerFeedReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeystoneFeedsConsumerFeedReceived)
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

	select {
	case log := <-it.logs:
		it.Event = new(KeystoneFeedsConsumerFeedReceived)
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

func (it *KeystoneFeedsConsumerFeedReceivedIterator) Error() error {
	return it.fail
}

func (it *KeystoneFeedsConsumerFeedReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type KeystoneFeedsConsumerFeedReceived struct {
	FeedId    [32]byte
	Price     *big.Int
	Timestamp uint32
	Raw       types.Log
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) FilterFeedReceived(opts *bind.FilterOpts, feedId [][32]byte) (*KeystoneFeedsConsumerFeedReceivedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _KeystoneFeedsConsumer.contract.FilterLogs(opts, "FeedReceived", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumerFeedReceivedIterator{contract: _KeystoneFeedsConsumer.contract, event: "FeedReceived", logs: logs, sub: sub}, nil
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) WatchFeedReceived(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerFeedReceived, feedId [][32]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _KeystoneFeedsConsumer.contract.WatchLogs(opts, "FeedReceived", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(KeystoneFeedsConsumerFeedReceived)
				if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "FeedReceived", log); err != nil {
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

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) ParseFeedReceived(log types.Log) (*KeystoneFeedsConsumerFeedReceived, error) {
	event := new(KeystoneFeedsConsumerFeedReceived)
	if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "FeedReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type KeystoneFeedsConsumerMessageReceivedIterator struct {
	Event *KeystoneFeedsConsumerMessageReceived

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *KeystoneFeedsConsumerMessageReceivedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeystoneFeedsConsumerMessageReceived)
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

	select {
	case log := <-it.logs:
		it.Event = new(KeystoneFeedsConsumerMessageReceived)
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

func (it *KeystoneFeedsConsumerMessageReceivedIterator) Error() error {
	return it.fail
}

func (it *KeystoneFeedsConsumerMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type KeystoneFeedsConsumerMessageReceived struct {
	Metadata  []byte
	RawReport []byte
	Raw       types.Log
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*KeystoneFeedsConsumerMessageReceivedIterator, error) {

	logs, sub, err := _KeystoneFeedsConsumer.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumerMessageReceivedIterator{contract: _KeystoneFeedsConsumer.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerMessageReceived) (event.Subscription, error) {

	logs, sub, err := _KeystoneFeedsConsumer.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(KeystoneFeedsConsumerMessageReceived)
				if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) ParseMessageReceived(log types.Log) (*KeystoneFeedsConsumerMessageReceived, error) {
	event := new(KeystoneFeedsConsumerMessageReceived)
	if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type KeystoneFeedsConsumerOwnershipTransferRequestedIterator struct {
	Event *KeystoneFeedsConsumerOwnershipTransferRequested

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *KeystoneFeedsConsumerOwnershipTransferRequestedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeystoneFeedsConsumerOwnershipTransferRequested)
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

	select {
	case log := <-it.logs:
		it.Event = new(KeystoneFeedsConsumerOwnershipTransferRequested)
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

func (it *KeystoneFeedsConsumerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

func (it *KeystoneFeedsConsumerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type KeystoneFeedsConsumerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeystoneFeedsConsumerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeystoneFeedsConsumer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumerOwnershipTransferRequestedIterator{contract: _KeystoneFeedsConsumer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeystoneFeedsConsumer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(KeystoneFeedsConsumerOwnershipTransferRequested)
				if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) ParseOwnershipTransferRequested(log types.Log) (*KeystoneFeedsConsumerOwnershipTransferRequested, error) {
	event := new(KeystoneFeedsConsumerOwnershipTransferRequested)
	if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type KeystoneFeedsConsumerOwnershipTransferredIterator struct {
	Event *KeystoneFeedsConsumerOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *KeystoneFeedsConsumerOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeystoneFeedsConsumerOwnershipTransferred)
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

	select {
	case log := <-it.logs:
		it.Event = new(KeystoneFeedsConsumerOwnershipTransferred)
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

func (it *KeystoneFeedsConsumerOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *KeystoneFeedsConsumerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type KeystoneFeedsConsumerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeystoneFeedsConsumerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeystoneFeedsConsumer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumerOwnershipTransferredIterator{contract: _KeystoneFeedsConsumer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KeystoneFeedsConsumer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(KeystoneFeedsConsumerOwnershipTransferred)
				if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) ParseOwnershipTransferred(log types.Log) (*KeystoneFeedsConsumerOwnershipTransferred, error) {
	event := new(KeystoneFeedsConsumerOwnershipTransferred)
	if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type KeystoneFeedsConsumerReportPermissionSetIterator struct {
	Event *KeystoneFeedsConsumerReportPermissionSet

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *KeystoneFeedsConsumerReportPermissionSetIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KeystoneFeedsConsumerReportPermissionSet)
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

	select {
	case log := <-it.logs:
		it.Event = new(KeystoneFeedsConsumerReportPermissionSet)
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

func (it *KeystoneFeedsConsumerReportPermissionSetIterator) Error() error {
	return it.fail
}

func (it *KeystoneFeedsConsumerReportPermissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type KeystoneFeedsConsumerReportPermissionSet struct {
	ReportId   [32]byte
	Permission KeystoneFeedsPermissionHandlerPermission
	Raw        types.Log
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*KeystoneFeedsConsumerReportPermissionSetIterator, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _KeystoneFeedsConsumer.contract.FilterLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return &KeystoneFeedsConsumerReportPermissionSetIterator{contract: _KeystoneFeedsConsumer.contract, event: "ReportPermissionSet", logs: logs, sub: sub}, nil
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerReportPermissionSet, reportId [][32]byte) (event.Subscription, error) {

	var reportIdRule []interface{}
	for _, reportIdItem := range reportId {
		reportIdRule = append(reportIdRule, reportIdItem)
	}

	logs, sub, err := _KeystoneFeedsConsumer.contract.WatchLogs(opts, "ReportPermissionSet", reportIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(KeystoneFeedsConsumerReportPermissionSet)
				if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
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

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumerFilterer) ParseReportPermissionSet(log types.Log) (*KeystoneFeedsConsumerReportPermissionSet, error) {
	event := new(KeystoneFeedsConsumerReportPermissionSet)
	if err := _KeystoneFeedsConsumer.contract.UnpackLog(event, "ReportPermissionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumer) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _KeystoneFeedsConsumer.abi.Events["FeedReceived"].ID:
		return _KeystoneFeedsConsumer.ParseFeedReceived(log)
	case _KeystoneFeedsConsumer.abi.Events["MessageReceived"].ID:
		return _KeystoneFeedsConsumer.ParseMessageReceived(log)
	case _KeystoneFeedsConsumer.abi.Events["OwnershipTransferRequested"].ID:
		return _KeystoneFeedsConsumer.ParseOwnershipTransferRequested(log)
	case _KeystoneFeedsConsumer.abi.Events["OwnershipTransferred"].ID:
		return _KeystoneFeedsConsumer.ParseOwnershipTransferred(log)
	case _KeystoneFeedsConsumer.abi.Events["ReportPermissionSet"].ID:
		return _KeystoneFeedsConsumer.ParseReportPermissionSet(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (KeystoneFeedsConsumerFeedReceived) Topic() common.Hash {
	return common.HexToHash("0x2c30f5cb3caf4239d0f994ce539d7ef24817fa550169c388e3a110f02e40197d")
}

func (KeystoneFeedsConsumerMessageReceived) Topic() common.Hash {
	return common.HexToHash("0x245e8e0654076d706bf5dea6eab892c9d7788b4aefa7f4923e837e43d200959b")
}

func (KeystoneFeedsConsumerOwnershipTransferRequested) Topic() common.Hash {
	return common.HexToHash("0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278")
}

func (KeystoneFeedsConsumerOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (KeystoneFeedsConsumerReportPermissionSet) Topic() common.Hash {
	return common.HexToHash("0x32a4ba3fa3351b11ad555d4c8ec70a744e8705607077a946807030d64b6ab1a3")
}

func (_KeystoneFeedsConsumer *KeystoneFeedsConsumer) Address() common.Address {
	return _KeystoneFeedsConsumer.address
}

type KeystoneFeedsConsumerInterface interface {
	GetPrice(opts *bind.CallOpts, feedId [32]byte) (*big.Int, uint32, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	OnReport(opts *bind.TransactOpts, metadata []byte, rawReport []byte) (*types.Transaction, error)

	SetReportPermissions(opts *bind.TransactOpts, permissions []KeystoneFeedsPermissionHandlerPermission) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)

	FilterFeedReceived(opts *bind.FilterOpts, feedId [][32]byte) (*KeystoneFeedsConsumerFeedReceivedIterator, error)

	WatchFeedReceived(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerFeedReceived, feedId [][32]byte) (event.Subscription, error)

	ParseFeedReceived(log types.Log) (*KeystoneFeedsConsumerFeedReceived, error)

	FilterMessageReceived(opts *bind.FilterOpts) (*KeystoneFeedsConsumerMessageReceivedIterator, error)

	WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerMessageReceived) (event.Subscription, error)

	ParseMessageReceived(log types.Log) (*KeystoneFeedsConsumerMessageReceived, error)

	FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeystoneFeedsConsumerOwnershipTransferRequestedIterator, error)

	WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferRequested(log types.Log) (*KeystoneFeedsConsumerOwnershipTransferRequested, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KeystoneFeedsConsumerOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*KeystoneFeedsConsumerOwnershipTransferred, error)

	FilterReportPermissionSet(opts *bind.FilterOpts, reportId [][32]byte) (*KeystoneFeedsConsumerReportPermissionSetIterator, error)

	WatchReportPermissionSet(opts *bind.WatchOpts, sink chan<- *KeystoneFeedsConsumerReportPermissionSet, reportId [][32]byte) (event.Subscription, error)

	ParseReportPermissionSet(log types.Log) (*KeystoneFeedsConsumerReportPermissionSet, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
