// Code generated by mockery v2.0.3. DO NOT EDIT.

package mocks

import (
	big "math/big"

	assets "github.com/smartcontractkit/chainlink/core/assets"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	eth "github.com/smartcontractkit/chainlink/core/services/eth"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Call provides a mock function with given fields: result, method, args
func (_m *Client) Call(result interface{}, method string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, result, method)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) error); ok {
		r0 = rf(result, method, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBlockByNumber provides a mock function with given fields: hex
func (_m *Client) GetBlockByNumber(hex string) (models.Block, error) {
	ret := _m.Called(hex)

	var r0 models.Block
	if rf, ok := ret.Get(0).(func(string) models.Block); ok {
		r0 = rf(hex)
	} else {
		r0 = ret.Get(0).(models.Block)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(hex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockHeight provides a mock function with given fields:
func (_m *Client) GetBlockHeight() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChainID provides a mock function with given fields:
func (_m *Client) GetChainID() (*big.Int, error) {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetERC20Balance provides a mock function with given fields: address, contractAddress
func (_m *Client) GetERC20Balance(address common.Address, contractAddress common.Address) (*big.Int, error) {
	ret := _m.Called(address, contractAddress)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) *big.Int); ok {
		r0 = rf(address, contractAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address) error); ok {
		r1 = rf(address, contractAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEthBalance provides a mock function with given fields: address
func (_m *Client) GetEthBalance(address common.Address) (*assets.Eth, error) {
	ret := _m.Called(address)

	var r0 *assets.Eth
	if rf, ok := ret.Get(0).(func(common.Address) *assets.Eth); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Eth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestBlock provides a mock function with given fields:
func (_m *Client) GetLatestBlock() (models.Block, error) {
	ret := _m.Called()

	var r0 models.Block
	if rf, ok := ret.Get(0).(func() models.Block); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Block)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogs provides a mock function with given fields: q
func (_m *Client) GetLogs(q ethereum.FilterQuery) ([]models.Log, error) {
	ret := _m.Called(q)

	var r0 []models.Log
	if rf, ok := ret.Get(0).(func(ethereum.FilterQuery) []models.Log); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ethereum.FilterQuery) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: address
func (_m *Client) GetNonce(address common.Address) (uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxReceipt provides a mock function with given fields: hash
func (_m *Client) GetTxReceipt(hash common.Hash) (*models.TxReceipt, error) {
	ret := _m.Called(hash)

	var r0 *models.TxReceipt
	if rf, ok := ret.Get(0).(func(common.Hash) *models.TxReceipt); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.TxReceipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GethClient provides a mock function with given fields: _a0
func (_m *Client) GethClient(_a0 func(eth.GethClient) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(eth.GethClient) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RPCClient provides a mock function with given fields: _a0
func (_m *Client) RPCClient(_a0 func(eth.RPCClient) error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(eth.RPCClient) error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendRawTx provides a mock function with given fields: bytes
func (_m *Client) SendRawTx(bytes []byte) (common.Hash, error) {
	ret := _m.Called(bytes)

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func([]byte) common.Hash); ok {
		r0 = rf(bytes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(bytes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) Subscribe(_a0 context.Context, _a1 interface{}, _a2 ...interface{}) (eth.Subscription, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 eth.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...interface{}) eth.Subscription); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToLogs provides a mock function with given fields: ctx, channel, q
func (_m *Client) SubscribeToLogs(ctx context.Context, channel chan<- models.Log, q ethereum.FilterQuery) (eth.Subscription, error) {
	ret := _m.Called(ctx, channel, q)

	var r0 eth.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, chan<- models.Log, ethereum.FilterQuery) eth.Subscription); ok {
		r0 = rf(ctx, channel, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- models.Log, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, channel, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeToNewHeads provides a mock function with given fields: ctx, channel
func (_m *Client) SubscribeToNewHeads(ctx context.Context, channel chan<- types.Header) (eth.Subscription, error) {
	ret := _m.Called(ctx, channel)

	var r0 eth.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, chan<- types.Header) eth.Subscription); ok {
		r0 = rf(ctx, channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, chan<- types.Header) error); ok {
		r1 = rf(ctx, channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
