// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package chain_reader_tester

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

type InnerTestStruct struct {
	IntVal int64
	S      string
}

type MidLevelTestStruct struct {
	FixedBytes [2]byte
	Inner      InnerTestStruct
}

type TestStruct struct {
	Field          int32
	DifferentField string
	OracleId       uint8
	OracleIds      [32]uint8
	Account        common.Address
	Accounts       []common.Address
	BigField       *big.Int
	NestedStruct   MidLevelTestStruct
}

var ChainReaderTesterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"structInnerTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"indexed\":false,\"internalType\":\"structMidLevelTestStruct\",\"name\":\"nestedStruct\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"}],\"name\":\"Triggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"fieldHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"}],\"name\":\"TriggeredEventWithDynamicTopic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field1\",\"type\":\"int32\"},{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field2\",\"type\":\"int32\"},{\"indexed\":true,\"internalType\":\"int32\",\"name\":\"field3\",\"type\":\"int32\"}],\"name\":\"TriggeredWithFourTopics\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"field1\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint8[32]\",\"name\":\"field2\",\"type\":\"uint8[32]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"field3\",\"type\":\"bytes32\"}],\"name\":\"TriggeredWithFourTopicsWithHashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"structInnerTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"structMidLevelTestStruct\",\"name\":\"nestedStruct\",\"type\":\"tuple\"}],\"name\":\"addTestStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAlterablePrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDifferentPrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getElementAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"Field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"DifferentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"OracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"OracleIds\",\"type\":\"uint8[32]\"},{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"BigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"structInnerTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"structMidLevelTestStruct\",\"name\":\"NestedStruct\",\"type\":\"tuple\"}],\"internalType\":\"structTestStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrimitiveValue\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSliceValue\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"structInnerTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"structMidLevelTestStruct\",\"name\":\"nestedStruct\",\"type\":\"tuple\"}],\"name\":\"returnSeen\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"Field\",\"type\":\"int32\"},{\"internalType\":\"string\",\"name\":\"DifferentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"OracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"OracleIds\",\"type\":\"uint8[32]\"},{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"Accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"BigField\",\"type\":\"int192\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"structInnerTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"structMidLevelTestStruct\",\"name\":\"NestedStruct\",\"type\":\"tuple\"}],\"internalType\":\"structTestStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"setAlterablePrimitiveValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field\",\"type\":\"int32\"},{\"components\":[{\"internalType\":\"bytes2\",\"name\":\"FixedBytes\",\"type\":\"bytes2\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"IntVal\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"S\",\"type\":\"string\"}],\"internalType\":\"structInnerTestStruct\",\"name\":\"Inner\",\"type\":\"tuple\"}],\"internalType\":\"structMidLevelTestStruct\",\"name\":\"nestedStruct\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"differentField\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"oracleId\",\"type\":\"uint8\"},{\"internalType\":\"uint8[32]\",\"name\":\"oracleIds\",\"type\":\"uint8[32]\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"int192\",\"name\":\"bigField\",\"type\":\"int192\"}],\"name\":\"triggerEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"}],\"name\":\"triggerEventWithDynamicTopic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"field1\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"field2\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"field3\",\"type\":\"int32\"}],\"name\":\"triggerWithFourTopics\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field1\",\"type\":\"string\"},{\"internalType\":\"uint8[32]\",\"name\":\"field2\",\"type\":\"uint8[32]\"},{\"internalType\":\"bytes32\",\"name\":\"field3\",\"type\":\"bytes32\"}],\"name\":\"triggerWithFourTopicsWithHashed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600180548082018255600082905260048082047fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6908101805460086003958616810261010090810a8088026001600160401b0391820219909416939093179093558654808801909755848704909301805496909516909202900a918202910219909216919091179055611a7b806100a96000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c8063a90e199811610081578063ef4e1ced1161005b578063ef4e1ced146101de578063f6f871c8146101e5578063fbe9fbf6146101f857600080fd5b8063a90e19981461019b578063ab5e0b38146101ae578063dbfd7332146101cb57600080fd5b80636104b97a116100b25780636104b97a1461012a578063679004a41461013d5780636c9a43b61461015257600080fd5b80632c45576f146100d95780633272b66c1461010257806349eac2ac14610117575b600080fd5b6100ec6100e7366004610ca5565b61020a565b6040516100f99190610e0e565b60405180910390f35b610115610110366004610f4d565b6104e5565b005b610115610125366004611062565b61053a565b610115610138366004611154565b61083d565b610145610894565b6040516100f99190611243565b610115610160366004611291565b600280547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff92909216919091179055565b6101156101a93660046113c5565b610920565b6107c65b60405167ffffffffffffffff90911681526020016100f9565b6101156101d936600461147a565b61097a565b60036101b2565b6100ec6101f3366004611062565b6109b7565b60025467ffffffffffffffff166101b2565b610212610ac2565b600061021f6001846114bd565b8154811061022f5761022f6114f7565b6000918252602091829020604080516101008101909152600a90920201805460030b8252600181018054929391929184019161026a90611526565b80601f016020809104026020016040519081016040528092919081815260200182805461029690611526565b80156102e35780601f106102b8576101008083540402835291602001916102e3565b820191906000526020600020905b8154815290600101906020018083116102c657829003601f168201915b5050509183525050600282015460ff166020808301919091526040805161040081018083529190930192916003850191826000855b825461010083900a900460ff1681526020600192830181810494850194909303909202910180841161031857505050928452505050600482015473ffffffffffffffffffffffffffffffffffffffff1660208083019190915260058301805460408051828502810185018252828152940193928301828280156103d157602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116103a6575b5050509183525050600682015460170b6020808301919091526040805180820182526007808601805460f01b7fffff0000000000000000000000000000000000000000000000000000000000001683528351808501855260088801805490930b8152600988018054959097019693959194868301949193928401919061045690611526565b80601f016020809104026020016040519081016040528092919081815260200182805461048290611526565b80156104cf5780601f106104a4576101008083540402835291602001916104cf565b820191906000526020600020905b8154815290600101906020018083116104b257829003601f168201915b5050509190925250505090525090525092915050565b81816040516104f5929190611573565b60405180910390207f3d969732b1bbbb9f1d7eb9f3f14e4cb50a74d950b3ef916a397b85dfbab93c67838360405161052e9291906115cc565b60405180910390a25050565b60006040518061010001604052808c60030b81526020018b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060ff8a166020808301919091526040805161040081810183529190930192918b9183908390808284376000920191909152505050815273ffffffffffffffffffffffffffffffffffffffff8816602080830191909152604080518883028181018401835289825291909301929189918991829190850190849080828437600092019190915250505090825250601785900b602082015260400161062c84611622565b905281546001808201845560009384526020938490208351600a9093020180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff909316929092178255928201519192909190820190610692908261170f565b5060408201516002820180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff90921691909117905560608201516106e09060038301906020610b11565b5060808201516004820180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691909117905560a08201518051610747916005840191602090910190610ba4565b5060c08201516006820180547fffffffffffffffff0000000000000000000000000000000000000000000000001677ffffffffffffffffffffffffffffffffffffffffffffffff90921691909117905560e082015180516007830180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660f09290921c91909117815560208083015180516008860180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff90921691909117815591810151909190600986019061082a908261170f565b5050505050505050505050505050505050565b8960030b7fd9efd9ec8d76dc6f59023b0aaa46f526c8248b77378282121ab309fc66dd5d528a8a8a8a8a8a8a8a8a604051610880999897969594939291906118a4565b60405180910390a250505050505050505050565b6060600180548060200260200160405190810160405280929190818152602001828054801561091657602002820191906000526020600020906000905b82829054906101000a900467ffffffffffffffff1667ffffffffffffffff16815260200190600801906020826007010492830192600103820291508084116108d15790505b5050505050905090565b808260405161092f9190611a16565b6040518091039020846040516109459190611a52565b604051908190038120907f7220e4dbe4e9d0ed5f71acd022bc89c26748ac6784f2c548bc17bb8e52af34b090600090a4505050565b8060030b8260030b8460030b7f91c80dc390f3d041b3a04b0099b19634499541ea26972250986ee4b24a12fac560405160405180910390a4505050565b6109bf610ac2565b6040518061010001604052808c60030b81526020018b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060ff8a166020808301919091526040805161040081810183529190930192918b9183908390808284376000920191909152505050815273ffffffffffffffffffffffffffffffffffffffff8816602080830191909152604080518883028181018401835289825291909301929189918991829190850190849080828437600092019190915250505090825250601785900b6020820152604001610aaf84611622565b905290505b9a9950505050505050505050565b6040805161010081018252600080825260606020830181905292820152908101610aea610c1e565b8152600060208201819052606060408301819052820152608001610b0c610c3d565b905290565b600183019183908215610b945791602002820160005b83821115610b6557835183826101000a81548160ff021916908360ff1602179055509260200192600101602081600001049283019260010302610b27565b8015610b925782816101000a81549060ff0219169055600101602081600001049283019260010302610b65565b505b50610ba0929150610c90565b5090565b828054828255906000526020600020908101928215610b94579160200282015b82811115610b9457825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190610bc4565b6040518061040001604052806020906020820280368337509192915050565b604051806040016040528060007dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152602001610b0c6040518060400160405280600060070b8152602001606081525090565b5b80821115610ba05760008155600101610c91565b600060208284031215610cb757600080fd5b5035919050565b60005b83811015610cd9578181015183820152602001610cc1565b50506000910152565b60008151808452610cfa816020860160208601610cbe565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8060005b6020808210610d3f5750610d56565b825160ff1685529384019390910190600101610d30565b50505050565b600081518084526020808501945080840160005b83811015610da257815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101610d70565b509495945050505050565b7fffff00000000000000000000000000000000000000000000000000000000000081511682526000602082015160406020850152805160070b60408501526020810151905060406060850152610e066080850182610ce2565b949350505050565b60208152610e2260208201835160030b9052565b600060208301516104e0806040850152610e40610500850183610ce2565b91506040850151610e56606086018260ff169052565b506060850151610e696080860182610d2c565b50608085015173ffffffffffffffffffffffffffffffffffffffff1661048085015260a08501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085840381016104a0870152610ec68483610d5c565b935060c08701519150610edf6104c087018360170b9052565b60e0870151915080868503018387015250610efa8382610dad565b9695505050505050565b60008083601f840112610f1657600080fd5b50813567ffffffffffffffff811115610f2e57600080fd5b602083019150836020828501011115610f4657600080fd5b9250929050565b60008060208385031215610f6057600080fd5b823567ffffffffffffffff811115610f7757600080fd5b610f8385828601610f04565b90969095509350505050565b8035600381900b8114610fa157600080fd5b919050565b803560ff81168114610fa157600080fd5b806104008101831015610fc957600080fd5b92915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610fa157600080fd5b60008083601f84011261100557600080fd5b50813567ffffffffffffffff81111561101d57600080fd5b6020830191508360208260051b8501011115610f4657600080fd5b8035601781900b8114610fa157600080fd5b60006040828403121561105c57600080fd5b50919050565b6000806000806000806000806000806104e08b8d03121561108257600080fd5b61108b8b610f8f565b995060208b013567ffffffffffffffff808211156110a857600080fd5b6110b48e838f01610f04565b909b5099508991506110c860408e01610fa6565b98506110d78e60608f01610fb7565b97506110e66104608e01610fcf565b96506104808d01359150808211156110fd57600080fd5b6111098e838f01610ff3565b909650945084915061111e6104a08e01611038565b93506104c08d013591508082111561113557600080fd5b506111428d828e0161104a565b9150509295989b9194979a5092959850565b6000806000806000806000806000806104e08b8d03121561117457600080fd5b61117d8b610f8f565b995060208b013567ffffffffffffffff8082111561119a57600080fd5b6111a68e838f0161104a565b9a5060408d01359150808211156111bc57600080fd5b6111c88e838f01610f04565b909a5098508891506111dc60608e01610fa6565b97506111eb8e60808f01610fb7565b96506111fa6104808e01610fcf565b95506104a08d013591508082111561121157600080fd5b5061121e8d828e01610ff3565b909450925061123290506104c08c01611038565b90509295989b9194979a5092959850565b6020808252825182820181905260009190848201906040850190845b8181101561128557835167ffffffffffffffff168352928401929184019160010161125f565b50909695505050505050565b6000602082840312156112a357600080fd5b813567ffffffffffffffff811681146112bb57600080fd5b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611314576113146112c2565b60405290565b600082601f83011261132b57600080fd5b813567ffffffffffffffff80821115611346576113466112c2565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561138c5761138c6112c2565b816040528381528660208588010111156113a557600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600061044084860312156113db57600080fd5b833567ffffffffffffffff808211156113f357600080fd5b6113ff8783880161131a565b94506020915086603f87011261141457600080fd5b60405161040081018181108382111715611430576114306112c2565b60405290508061042087018881111561144857600080fd5b8388015b8181101561146a5761145d81610fa6565b845292840192840161144c565b5095989097509435955050505050565b60008060006060848603121561148f57600080fd5b61149884610f8f565b92506114a660208501610f8f565b91506114b460408501610f8f565b90509250925092565b81810381811115610fc9577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c9082168061153a57607f821691505b60208210810361105c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b602081526000610e06602083018486611583565b80357fffff00000000000000000000000000000000000000000000000000000000000081168114610fa157600080fd5b8035600781900b8114610fa157600080fd5b60006040823603121561163457600080fd5b61163c6112f1565b611645836115e0565b8152602083013567ffffffffffffffff8082111561166257600080fd5b81850191506040823603121561167757600080fd5b61167f6112f1565b61168883611610565b815260208301358281111561169c57600080fd5b6116a83682860161131a565b60208301525080602085015250505080915050919050565b601f82111561170a57600081815260208120601f850160051c810160208610156116e75750805b601f850160051c820191505b81811015611706578281556001016116f3565b5050505b505050565b815167ffffffffffffffff811115611729576117296112c2565b61173d816117378454611526565b846116c0565b602080601f831160018114611790576000841561175a5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611706565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156117dd578886015182559484019460019091019084016117be565b508582101561181957878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b8060005b602080821061183c5750610d56565b60ff61184784610fa6565b16855293840193919091019060010161182d565b8183526000602080850194508260005b85811015610da25773ffffffffffffffffffffffffffffffffffffffff61189183610fcf565b168752958201959082019060010161186b565b60006104c08083527fffff0000000000000000000000000000000000000000000000000000000000006118d68d6115e0565b16818401525060208b01357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc18c360301811261191157600080fd5b60406104e08401528b0161192481611610565b60070b61050084015260208101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe182360301811261196257600080fd5b0160208101903567ffffffffffffffff81111561197e57600080fd5b80360382131561198d57600080fd5b60406105208501526119a461054085018284611583565b91505082810360208401526119ba818b8d611583565b60ff8a16604085015290506119d26060840189611829565b73ffffffffffffffffffffffffffffffffffffffff8716610460840152828103610480840152611a0381868861185b565b915050610ab46104a083018460170b9052565b60008183825b6020808210611a2b5750611a42565b825160ff1684529283019290910190600101611a1c565b5050506104008201905092915050565b60008251611a64818460208701610cbe565b919091019291505056fea164736f6c6343000813000a",
}

var ChainReaderTesterABI = ChainReaderTesterMetaData.ABI

var ChainReaderTesterBin = ChainReaderTesterMetaData.Bin

func DeployChainReaderTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChainReaderTester, error) {
	parsed, err := ChainReaderTesterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChainReaderTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChainReaderTester{address: address, abi: *parsed, ChainReaderTesterCaller: ChainReaderTesterCaller{contract: contract}, ChainReaderTesterTransactor: ChainReaderTesterTransactor{contract: contract}, ChainReaderTesterFilterer: ChainReaderTesterFilterer{contract: contract}}, nil
}

type ChainReaderTester struct {
	address common.Address
	abi     abi.ABI
	ChainReaderTesterCaller
	ChainReaderTesterTransactor
	ChainReaderTesterFilterer
}

type ChainReaderTesterCaller struct {
	contract *bind.BoundContract
}

type ChainReaderTesterTransactor struct {
	contract *bind.BoundContract
}

type ChainReaderTesterFilterer struct {
	contract *bind.BoundContract
}

type ChainReaderTesterSession struct {
	Contract     *ChainReaderTester
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type ChainReaderTesterCallerSession struct {
	Contract *ChainReaderTesterCaller
	CallOpts bind.CallOpts
}

type ChainReaderTesterTransactorSession struct {
	Contract     *ChainReaderTesterTransactor
	TransactOpts bind.TransactOpts
}

type ChainReaderTesterRaw struct {
	Contract *ChainReaderTester
}

type ChainReaderTesterCallerRaw struct {
	Contract *ChainReaderTesterCaller
}

type ChainReaderTesterTransactorRaw struct {
	Contract *ChainReaderTesterTransactor
}

func NewChainReaderTester(address common.Address, backend bind.ContractBackend) (*ChainReaderTester, error) {
	abi, err := abi.JSON(strings.NewReader(ChainReaderTesterABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindChainReaderTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainReaderTester{address: address, abi: abi, ChainReaderTesterCaller: ChainReaderTesterCaller{contract: contract}, ChainReaderTesterTransactor: ChainReaderTesterTransactor{contract: contract}, ChainReaderTesterFilterer: ChainReaderTesterFilterer{contract: contract}}, nil
}

func NewChainReaderTesterCaller(address common.Address, caller bind.ContractCaller) (*ChainReaderTesterCaller, error) {
	contract, err := bindChainReaderTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainReaderTesterCaller{contract: contract}, nil
}

func NewChainReaderTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainReaderTesterTransactor, error) {
	contract, err := bindChainReaderTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainReaderTesterTransactor{contract: contract}, nil
}

func NewChainReaderTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainReaderTesterFilterer, error) {
	contract, err := bindChainReaderTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainReaderTesterFilterer{contract: contract}, nil
}

func bindChainReaderTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChainReaderTesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_ChainReaderTester *ChainReaderTesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainReaderTester.Contract.ChainReaderTesterCaller.contract.Call(opts, result, method, params...)
}

func (_ChainReaderTester *ChainReaderTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.ChainReaderTesterTransactor.contract.Transfer(opts)
}

func (_ChainReaderTester *ChainReaderTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.ChainReaderTesterTransactor.contract.Transact(opts, method, params...)
}

func (_ChainReaderTester *ChainReaderTesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainReaderTester.Contract.contract.Call(opts, result, method, params...)
}

func (_ChainReaderTester *ChainReaderTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.contract.Transfer(opts)
}

func (_ChainReaderTester *ChainReaderTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.contract.Transact(opts, method, params...)
}

func (_ChainReaderTester *ChainReaderTesterCaller) GetAlterablePrimitiveValue(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ChainReaderTester.contract.Call(opts, &out, "getAlterablePrimitiveValue")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_ChainReaderTester *ChainReaderTesterSession) GetAlterablePrimitiveValue() (uint64, error) {
	return _ChainReaderTester.Contract.GetAlterablePrimitiveValue(&_ChainReaderTester.CallOpts)
}

func (_ChainReaderTester *ChainReaderTesterCallerSession) GetAlterablePrimitiveValue() (uint64, error) {
	return _ChainReaderTester.Contract.GetAlterablePrimitiveValue(&_ChainReaderTester.CallOpts)
}

func (_ChainReaderTester *ChainReaderTesterCaller) GetDifferentPrimitiveValue(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ChainReaderTester.contract.Call(opts, &out, "getDifferentPrimitiveValue")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_ChainReaderTester *ChainReaderTesterSession) GetDifferentPrimitiveValue() (uint64, error) {
	return _ChainReaderTester.Contract.GetDifferentPrimitiveValue(&_ChainReaderTester.CallOpts)
}

func (_ChainReaderTester *ChainReaderTesterCallerSession) GetDifferentPrimitiveValue() (uint64, error) {
	return _ChainReaderTester.Contract.GetDifferentPrimitiveValue(&_ChainReaderTester.CallOpts)
}

func (_ChainReaderTester *ChainReaderTesterCaller) GetElementAtIndex(opts *bind.CallOpts, i *big.Int) (TestStruct, error) {
	var out []interface{}
	err := _ChainReaderTester.contract.Call(opts, &out, "getElementAtIndex", i)

	if err != nil {
		return *new(TestStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(TestStruct)).(*TestStruct)

	return out0, err

}

func (_ChainReaderTester *ChainReaderTesterSession) GetElementAtIndex(i *big.Int) (TestStruct, error) {
	return _ChainReaderTester.Contract.GetElementAtIndex(&_ChainReaderTester.CallOpts, i)
}

func (_ChainReaderTester *ChainReaderTesterCallerSession) GetElementAtIndex(i *big.Int) (TestStruct, error) {
	return _ChainReaderTester.Contract.GetElementAtIndex(&_ChainReaderTester.CallOpts, i)
}

func (_ChainReaderTester *ChainReaderTesterCaller) GetPrimitiveValue(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ChainReaderTester.contract.Call(opts, &out, "getPrimitiveValue")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

func (_ChainReaderTester *ChainReaderTesterSession) GetPrimitiveValue() (uint64, error) {
	return _ChainReaderTester.Contract.GetPrimitiveValue(&_ChainReaderTester.CallOpts)
}

func (_ChainReaderTester *ChainReaderTesterCallerSession) GetPrimitiveValue() (uint64, error) {
	return _ChainReaderTester.Contract.GetPrimitiveValue(&_ChainReaderTester.CallOpts)
}

func (_ChainReaderTester *ChainReaderTesterCaller) GetSliceValue(opts *bind.CallOpts) ([]uint64, error) {
	var out []interface{}
	err := _ChainReaderTester.contract.Call(opts, &out, "getSliceValue")

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

func (_ChainReaderTester *ChainReaderTesterSession) GetSliceValue() ([]uint64, error) {
	return _ChainReaderTester.Contract.GetSliceValue(&_ChainReaderTester.CallOpts)
}

func (_ChainReaderTester *ChainReaderTesterCallerSession) GetSliceValue() ([]uint64, error) {
	return _ChainReaderTester.Contract.GetSliceValue(&_ChainReaderTester.CallOpts)
}

func (_ChainReaderTester *ChainReaderTesterCaller) ReturnSeen(opts *bind.CallOpts, field int32, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int, nestedStruct MidLevelTestStruct) (TestStruct, error) {
	var out []interface{}
	err := _ChainReaderTester.contract.Call(opts, &out, "returnSeen", field, differentField, oracleId, oracleIds, account, accounts, bigField, nestedStruct)

	if err != nil {
		return *new(TestStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(TestStruct)).(*TestStruct)

	return out0, err

}

func (_ChainReaderTester *ChainReaderTesterSession) ReturnSeen(field int32, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int, nestedStruct MidLevelTestStruct) (TestStruct, error) {
	return _ChainReaderTester.Contract.ReturnSeen(&_ChainReaderTester.CallOpts, field, differentField, oracleId, oracleIds, account, accounts, bigField, nestedStruct)
}

func (_ChainReaderTester *ChainReaderTesterCallerSession) ReturnSeen(field int32, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int, nestedStruct MidLevelTestStruct) (TestStruct, error) {
	return _ChainReaderTester.Contract.ReturnSeen(&_ChainReaderTester.CallOpts, field, differentField, oracleId, oracleIds, account, accounts, bigField, nestedStruct)
}

func (_ChainReaderTester *ChainReaderTesterTransactor) AddTestStruct(opts *bind.TransactOpts, field int32, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int, nestedStruct MidLevelTestStruct) (*types.Transaction, error) {
	return _ChainReaderTester.contract.Transact(opts, "addTestStruct", field, differentField, oracleId, oracleIds, account, accounts, bigField, nestedStruct)
}

func (_ChainReaderTester *ChainReaderTesterSession) AddTestStruct(field int32, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int, nestedStruct MidLevelTestStruct) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.AddTestStruct(&_ChainReaderTester.TransactOpts, field, differentField, oracleId, oracleIds, account, accounts, bigField, nestedStruct)
}

func (_ChainReaderTester *ChainReaderTesterTransactorSession) AddTestStruct(field int32, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int, nestedStruct MidLevelTestStruct) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.AddTestStruct(&_ChainReaderTester.TransactOpts, field, differentField, oracleId, oracleIds, account, accounts, bigField, nestedStruct)
}

func (_ChainReaderTester *ChainReaderTesterTransactor) SetAlterablePrimitiveValue(opts *bind.TransactOpts, value uint64) (*types.Transaction, error) {
	return _ChainReaderTester.contract.Transact(opts, "setAlterablePrimitiveValue", value)
}

func (_ChainReaderTester *ChainReaderTesterSession) SetAlterablePrimitiveValue(value uint64) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.SetAlterablePrimitiveValue(&_ChainReaderTester.TransactOpts, value)
}

func (_ChainReaderTester *ChainReaderTesterTransactorSession) SetAlterablePrimitiveValue(value uint64) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.SetAlterablePrimitiveValue(&_ChainReaderTester.TransactOpts, value)
}

func (_ChainReaderTester *ChainReaderTesterTransactor) TriggerEvent(opts *bind.TransactOpts, field int32, nestedStruct MidLevelTestStruct, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int) (*types.Transaction, error) {
	return _ChainReaderTester.contract.Transact(opts, "triggerEvent", field, nestedStruct, differentField, oracleId, oracleIds, account, accounts, bigField)
}

func (_ChainReaderTester *ChainReaderTesterSession) TriggerEvent(field int32, nestedStruct MidLevelTestStruct, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.TriggerEvent(&_ChainReaderTester.TransactOpts, field, nestedStruct, differentField, oracleId, oracleIds, account, accounts, bigField)
}

func (_ChainReaderTester *ChainReaderTesterTransactorSession) TriggerEvent(field int32, nestedStruct MidLevelTestStruct, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.TriggerEvent(&_ChainReaderTester.TransactOpts, field, nestedStruct, differentField, oracleId, oracleIds, account, accounts, bigField)
}

func (_ChainReaderTester *ChainReaderTesterTransactor) TriggerEventWithDynamicTopic(opts *bind.TransactOpts, field string) (*types.Transaction, error) {
	return _ChainReaderTester.contract.Transact(opts, "triggerEventWithDynamicTopic", field)
}

func (_ChainReaderTester *ChainReaderTesterSession) TriggerEventWithDynamicTopic(field string) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.TriggerEventWithDynamicTopic(&_ChainReaderTester.TransactOpts, field)
}

func (_ChainReaderTester *ChainReaderTesterTransactorSession) TriggerEventWithDynamicTopic(field string) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.TriggerEventWithDynamicTopic(&_ChainReaderTester.TransactOpts, field)
}

func (_ChainReaderTester *ChainReaderTesterTransactor) TriggerWithFourTopics(opts *bind.TransactOpts, field1 int32, field2 int32, field3 int32) (*types.Transaction, error) {
	return _ChainReaderTester.contract.Transact(opts, "triggerWithFourTopics", field1, field2, field3)
}

func (_ChainReaderTester *ChainReaderTesterSession) TriggerWithFourTopics(field1 int32, field2 int32, field3 int32) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.TriggerWithFourTopics(&_ChainReaderTester.TransactOpts, field1, field2, field3)
}

func (_ChainReaderTester *ChainReaderTesterTransactorSession) TriggerWithFourTopics(field1 int32, field2 int32, field3 int32) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.TriggerWithFourTopics(&_ChainReaderTester.TransactOpts, field1, field2, field3)
}

func (_ChainReaderTester *ChainReaderTesterTransactor) TriggerWithFourTopicsWithHashed(opts *bind.TransactOpts, field1 string, field2 [32]uint8, field3 [32]byte) (*types.Transaction, error) {
	return _ChainReaderTester.contract.Transact(opts, "triggerWithFourTopicsWithHashed", field1, field2, field3)
}

func (_ChainReaderTester *ChainReaderTesterSession) TriggerWithFourTopicsWithHashed(field1 string, field2 [32]uint8, field3 [32]byte) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.TriggerWithFourTopicsWithHashed(&_ChainReaderTester.TransactOpts, field1, field2, field3)
}

func (_ChainReaderTester *ChainReaderTesterTransactorSession) TriggerWithFourTopicsWithHashed(field1 string, field2 [32]uint8, field3 [32]byte) (*types.Transaction, error) {
	return _ChainReaderTester.Contract.TriggerWithFourTopicsWithHashed(&_ChainReaderTester.TransactOpts, field1, field2, field3)
}

type ChainReaderTesterTriggeredIterator struct {
	Event *ChainReaderTesterTriggered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChainReaderTesterTriggeredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainReaderTesterTriggered)
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
		it.Event = new(ChainReaderTesterTriggered)
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

func (it *ChainReaderTesterTriggeredIterator) Error() error {
	return it.fail
}

func (it *ChainReaderTesterTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChainReaderTesterTriggered struct {
	Field          int32
	NestedStruct   MidLevelTestStruct
	DifferentField string
	OracleId       uint8
	OracleIds      [32]uint8
	Account        common.Address
	Accounts       []common.Address
	BigField       *big.Int
	Raw            types.Log
}

func (_ChainReaderTester *ChainReaderTesterFilterer) FilterTriggered(opts *bind.FilterOpts, field []int32) (*ChainReaderTesterTriggeredIterator, error) {

	var fieldRule []interface{}
	for _, fieldItem := range field {
		fieldRule = append(fieldRule, fieldItem)
	}

	logs, sub, err := _ChainReaderTester.contract.FilterLogs(opts, "Triggered", fieldRule)
	if err != nil {
		return nil, err
	}
	return &ChainReaderTesterTriggeredIterator{contract: _ChainReaderTester.contract, event: "Triggered", logs: logs, sub: sub}, nil
}

func (_ChainReaderTester *ChainReaderTesterFilterer) WatchTriggered(opts *bind.WatchOpts, sink chan<- *ChainReaderTesterTriggered, field []int32) (event.Subscription, error) {

	var fieldRule []interface{}
	for _, fieldItem := range field {
		fieldRule = append(fieldRule, fieldItem)
	}

	logs, sub, err := _ChainReaderTester.contract.WatchLogs(opts, "Triggered", fieldRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChainReaderTesterTriggered)
				if err := _ChainReaderTester.contract.UnpackLog(event, "Triggered", log); err != nil {
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

func (_ChainReaderTester *ChainReaderTesterFilterer) ParseTriggered(log types.Log) (*ChainReaderTesterTriggered, error) {
	event := new(ChainReaderTesterTriggered)
	if err := _ChainReaderTester.contract.UnpackLog(event, "Triggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ChainReaderTesterTriggeredEventWithDynamicTopicIterator struct {
	Event *ChainReaderTesterTriggeredEventWithDynamicTopic

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChainReaderTesterTriggeredEventWithDynamicTopicIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainReaderTesterTriggeredEventWithDynamicTopic)
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
		it.Event = new(ChainReaderTesterTriggeredEventWithDynamicTopic)
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

func (it *ChainReaderTesterTriggeredEventWithDynamicTopicIterator) Error() error {
	return it.fail
}

func (it *ChainReaderTesterTriggeredEventWithDynamicTopicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChainReaderTesterTriggeredEventWithDynamicTopic struct {
	FieldHash common.Hash
	Field     string
	Raw       types.Log
}

func (_ChainReaderTester *ChainReaderTesterFilterer) FilterTriggeredEventWithDynamicTopic(opts *bind.FilterOpts, fieldHash []string) (*ChainReaderTesterTriggeredEventWithDynamicTopicIterator, error) {

	var fieldHashRule []interface{}
	for _, fieldHashItem := range fieldHash {
		fieldHashRule = append(fieldHashRule, fieldHashItem)
	}

	logs, sub, err := _ChainReaderTester.contract.FilterLogs(opts, "TriggeredEventWithDynamicTopic", fieldHashRule)
	if err != nil {
		return nil, err
	}
	return &ChainReaderTesterTriggeredEventWithDynamicTopicIterator{contract: _ChainReaderTester.contract, event: "TriggeredEventWithDynamicTopic", logs: logs, sub: sub}, nil
}

func (_ChainReaderTester *ChainReaderTesterFilterer) WatchTriggeredEventWithDynamicTopic(opts *bind.WatchOpts, sink chan<- *ChainReaderTesterTriggeredEventWithDynamicTopic, fieldHash []string) (event.Subscription, error) {

	var fieldHashRule []interface{}
	for _, fieldHashItem := range fieldHash {
		fieldHashRule = append(fieldHashRule, fieldHashItem)
	}

	logs, sub, err := _ChainReaderTester.contract.WatchLogs(opts, "TriggeredEventWithDynamicTopic", fieldHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChainReaderTesterTriggeredEventWithDynamicTopic)
				if err := _ChainReaderTester.contract.UnpackLog(event, "TriggeredEventWithDynamicTopic", log); err != nil {
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

func (_ChainReaderTester *ChainReaderTesterFilterer) ParseTriggeredEventWithDynamicTopic(log types.Log) (*ChainReaderTesterTriggeredEventWithDynamicTopic, error) {
	event := new(ChainReaderTesterTriggeredEventWithDynamicTopic)
	if err := _ChainReaderTester.contract.UnpackLog(event, "TriggeredEventWithDynamicTopic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ChainReaderTesterTriggeredWithFourTopicsIterator struct {
	Event *ChainReaderTesterTriggeredWithFourTopics

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChainReaderTesterTriggeredWithFourTopicsIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainReaderTesterTriggeredWithFourTopics)
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
		it.Event = new(ChainReaderTesterTriggeredWithFourTopics)
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

func (it *ChainReaderTesterTriggeredWithFourTopicsIterator) Error() error {
	return it.fail
}

func (it *ChainReaderTesterTriggeredWithFourTopicsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChainReaderTesterTriggeredWithFourTopics struct {
	Field1 int32
	Field2 int32
	Field3 int32
	Raw    types.Log
}

func (_ChainReaderTester *ChainReaderTesterFilterer) FilterTriggeredWithFourTopics(opts *bind.FilterOpts, field1 []int32, field2 []int32, field3 []int32) (*ChainReaderTesterTriggeredWithFourTopicsIterator, error) {

	var field1Rule []interface{}
	for _, field1Item := range field1 {
		field1Rule = append(field1Rule, field1Item)
	}
	var field2Rule []interface{}
	for _, field2Item := range field2 {
		field2Rule = append(field2Rule, field2Item)
	}
	var field3Rule []interface{}
	for _, field3Item := range field3 {
		field3Rule = append(field3Rule, field3Item)
	}

	logs, sub, err := _ChainReaderTester.contract.FilterLogs(opts, "TriggeredWithFourTopics", field1Rule, field2Rule, field3Rule)
	if err != nil {
		return nil, err
	}
	return &ChainReaderTesterTriggeredWithFourTopicsIterator{contract: _ChainReaderTester.contract, event: "TriggeredWithFourTopics", logs: logs, sub: sub}, nil
}

func (_ChainReaderTester *ChainReaderTesterFilterer) WatchTriggeredWithFourTopics(opts *bind.WatchOpts, sink chan<- *ChainReaderTesterTriggeredWithFourTopics, field1 []int32, field2 []int32, field3 []int32) (event.Subscription, error) {

	var field1Rule []interface{}
	for _, field1Item := range field1 {
		field1Rule = append(field1Rule, field1Item)
	}
	var field2Rule []interface{}
	for _, field2Item := range field2 {
		field2Rule = append(field2Rule, field2Item)
	}
	var field3Rule []interface{}
	for _, field3Item := range field3 {
		field3Rule = append(field3Rule, field3Item)
	}

	logs, sub, err := _ChainReaderTester.contract.WatchLogs(opts, "TriggeredWithFourTopics", field1Rule, field2Rule, field3Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChainReaderTesterTriggeredWithFourTopics)
				if err := _ChainReaderTester.contract.UnpackLog(event, "TriggeredWithFourTopics", log); err != nil {
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

func (_ChainReaderTester *ChainReaderTesterFilterer) ParseTriggeredWithFourTopics(log types.Log) (*ChainReaderTesterTriggeredWithFourTopics, error) {
	event := new(ChainReaderTesterTriggeredWithFourTopics)
	if err := _ChainReaderTester.contract.UnpackLog(event, "TriggeredWithFourTopics", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type ChainReaderTesterTriggeredWithFourTopicsWithHashedIterator struct {
	Event *ChainReaderTesterTriggeredWithFourTopicsWithHashed

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *ChainReaderTesterTriggeredWithFourTopicsWithHashedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainReaderTesterTriggeredWithFourTopicsWithHashed)
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
		it.Event = new(ChainReaderTesterTriggeredWithFourTopicsWithHashed)
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

func (it *ChainReaderTesterTriggeredWithFourTopicsWithHashedIterator) Error() error {
	return it.fail
}

func (it *ChainReaderTesterTriggeredWithFourTopicsWithHashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type ChainReaderTesterTriggeredWithFourTopicsWithHashed struct {
	Field1 common.Hash
	Field2 [32]uint8
	Field3 [32]byte
	Raw    types.Log
}

func (_ChainReaderTester *ChainReaderTesterFilterer) FilterTriggeredWithFourTopicsWithHashed(opts *bind.FilterOpts, field1 []string, field2 [][32]uint8, field3 [][32]byte) (*ChainReaderTesterTriggeredWithFourTopicsWithHashedIterator, error) {

	var field1Rule []interface{}
	for _, field1Item := range field1 {
		field1Rule = append(field1Rule, field1Item)
	}
	var field2Rule []interface{}
	for _, field2Item := range field2 {
		field2Rule = append(field2Rule, field2Item)
	}
	var field3Rule []interface{}
	for _, field3Item := range field3 {
		field3Rule = append(field3Rule, field3Item)
	}

	logs, sub, err := _ChainReaderTester.contract.FilterLogs(opts, "TriggeredWithFourTopicsWithHashed", field1Rule, field2Rule, field3Rule)
	if err != nil {
		return nil, err
	}
	return &ChainReaderTesterTriggeredWithFourTopicsWithHashedIterator{contract: _ChainReaderTester.contract, event: "TriggeredWithFourTopicsWithHashed", logs: logs, sub: sub}, nil
}

func (_ChainReaderTester *ChainReaderTesterFilterer) WatchTriggeredWithFourTopicsWithHashed(opts *bind.WatchOpts, sink chan<- *ChainReaderTesterTriggeredWithFourTopicsWithHashed, field1 []string, field2 [][32]uint8, field3 [][32]byte) (event.Subscription, error) {

	var field1Rule []interface{}
	for _, field1Item := range field1 {
		field1Rule = append(field1Rule, field1Item)
	}
	var field2Rule []interface{}
	for _, field2Item := range field2 {
		field2Rule = append(field2Rule, field2Item)
	}
	var field3Rule []interface{}
	for _, field3Item := range field3 {
		field3Rule = append(field3Rule, field3Item)
	}

	logs, sub, err := _ChainReaderTester.contract.WatchLogs(opts, "TriggeredWithFourTopicsWithHashed", field1Rule, field2Rule, field3Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(ChainReaderTesterTriggeredWithFourTopicsWithHashed)
				if err := _ChainReaderTester.contract.UnpackLog(event, "TriggeredWithFourTopicsWithHashed", log); err != nil {
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

func (_ChainReaderTester *ChainReaderTesterFilterer) ParseTriggeredWithFourTopicsWithHashed(log types.Log) (*ChainReaderTesterTriggeredWithFourTopicsWithHashed, error) {
	event := new(ChainReaderTesterTriggeredWithFourTopicsWithHashed)
	if err := _ChainReaderTester.contract.UnpackLog(event, "TriggeredWithFourTopicsWithHashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_ChainReaderTester *ChainReaderTester) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _ChainReaderTester.abi.Events["Triggered"].ID:
		return _ChainReaderTester.ParseTriggered(log)
	case _ChainReaderTester.abi.Events["TriggeredEventWithDynamicTopic"].ID:
		return _ChainReaderTester.ParseTriggeredEventWithDynamicTopic(log)
	case _ChainReaderTester.abi.Events["TriggeredWithFourTopics"].ID:
		return _ChainReaderTester.ParseTriggeredWithFourTopics(log)
	case _ChainReaderTester.abi.Events["TriggeredWithFourTopicsWithHashed"].ID:
		return _ChainReaderTester.ParseTriggeredWithFourTopicsWithHashed(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (ChainReaderTesterTriggered) Topic() common.Hash {
	return common.HexToHash("0xd9efd9ec8d76dc6f59023b0aaa46f526c8248b77378282121ab309fc66dd5d52")
}

func (ChainReaderTesterTriggeredEventWithDynamicTopic) Topic() common.Hash {
	return common.HexToHash("0x3d969732b1bbbb9f1d7eb9f3f14e4cb50a74d950b3ef916a397b85dfbab93c67")
}

func (ChainReaderTesterTriggeredWithFourTopics) Topic() common.Hash {
	return common.HexToHash("0x91c80dc390f3d041b3a04b0099b19634499541ea26972250986ee4b24a12fac5")
}

func (ChainReaderTesterTriggeredWithFourTopicsWithHashed) Topic() common.Hash {
	return common.HexToHash("0x7220e4dbe4e9d0ed5f71acd022bc89c26748ac6784f2c548bc17bb8e52af34b0")
}

func (_ChainReaderTester *ChainReaderTester) Address() common.Address {
	return _ChainReaderTester.address
}

type ChainReaderTesterInterface interface {
	GetAlterablePrimitiveValue(opts *bind.CallOpts) (uint64, error)

	GetDifferentPrimitiveValue(opts *bind.CallOpts) (uint64, error)

	GetElementAtIndex(opts *bind.CallOpts, i *big.Int) (TestStruct, error)

	GetPrimitiveValue(opts *bind.CallOpts) (uint64, error)

	GetSliceValue(opts *bind.CallOpts) ([]uint64, error)

	ReturnSeen(opts *bind.CallOpts, field int32, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int, nestedStruct MidLevelTestStruct) (TestStruct, error)

	AddTestStruct(opts *bind.TransactOpts, field int32, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int, nestedStruct MidLevelTestStruct) (*types.Transaction, error)

	SetAlterablePrimitiveValue(opts *bind.TransactOpts, value uint64) (*types.Transaction, error)

	TriggerEvent(opts *bind.TransactOpts, field int32, nestedStruct MidLevelTestStruct, differentField string, oracleId uint8, oracleIds [32]uint8, account common.Address, accounts []common.Address, bigField *big.Int) (*types.Transaction, error)

	TriggerEventWithDynamicTopic(opts *bind.TransactOpts, field string) (*types.Transaction, error)

	TriggerWithFourTopics(opts *bind.TransactOpts, field1 int32, field2 int32, field3 int32) (*types.Transaction, error)

	TriggerWithFourTopicsWithHashed(opts *bind.TransactOpts, field1 string, field2 [32]uint8, field3 [32]byte) (*types.Transaction, error)

	FilterTriggered(opts *bind.FilterOpts, field []int32) (*ChainReaderTesterTriggeredIterator, error)

	WatchTriggered(opts *bind.WatchOpts, sink chan<- *ChainReaderTesterTriggered, field []int32) (event.Subscription, error)

	ParseTriggered(log types.Log) (*ChainReaderTesterTriggered, error)

	FilterTriggeredEventWithDynamicTopic(opts *bind.FilterOpts, fieldHash []string) (*ChainReaderTesterTriggeredEventWithDynamicTopicIterator, error)

	WatchTriggeredEventWithDynamicTopic(opts *bind.WatchOpts, sink chan<- *ChainReaderTesterTriggeredEventWithDynamicTopic, fieldHash []string) (event.Subscription, error)

	ParseTriggeredEventWithDynamicTopic(log types.Log) (*ChainReaderTesterTriggeredEventWithDynamicTopic, error)

	FilterTriggeredWithFourTopics(opts *bind.FilterOpts, field1 []int32, field2 []int32, field3 []int32) (*ChainReaderTesterTriggeredWithFourTopicsIterator, error)

	WatchTriggeredWithFourTopics(opts *bind.WatchOpts, sink chan<- *ChainReaderTesterTriggeredWithFourTopics, field1 []int32, field2 []int32, field3 []int32) (event.Subscription, error)

	ParseTriggeredWithFourTopics(log types.Log) (*ChainReaderTesterTriggeredWithFourTopics, error)

	FilterTriggeredWithFourTopicsWithHashed(opts *bind.FilterOpts, field1 []string, field2 [][32]uint8, field3 [][32]byte) (*ChainReaderTesterTriggeredWithFourTopicsWithHashedIterator, error)

	WatchTriggeredWithFourTopicsWithHashed(opts *bind.WatchOpts, sink chan<- *ChainReaderTesterTriggeredWithFourTopicsWithHashed, field1 []string, field2 [][32]uint8, field3 [][32]byte) (event.Subscription, error)

	ParseTriggeredWithFourTopicsWithHashed(log types.Log) (*ChainReaderTesterTriggeredWithFourTopicsWithHashed, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
