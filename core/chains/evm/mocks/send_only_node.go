// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	big "math/big"

	context "context"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/ethereum/go-ethereum/rpc"

	types "github.com/ethereum/go-ethereum/core/types"
)

// SendOnlyNode is an autogenerated mock type for the SendOnlyNode type
type SendOnlyNode struct {
	mock.Mock
}

// BatchCallContext provides a mock function with given fields: ctx, b
func (_m *SendOnlyNode) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	ret := _m.Called(ctx, b)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []rpc.BatchElem) error); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChainID provides a mock function with given fields:
func (_m *SendOnlyNode) ChainID() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *SendOnlyNode) Close() {
	_m.Called()
}

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *SendOnlyNode) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ret := _m.Called(ctx, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *SendOnlyNode) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *SendOnlyNode) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type NewSendOnlyNodeT interface {
	mock.TestingT
	Cleanup(func())
}

// NewSendOnlyNode creates a new instance of SendOnlyNode. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSendOnlyNode(t NewSendOnlyNodeT) *SendOnlyNode {
	mock := &SendOnlyNode{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
