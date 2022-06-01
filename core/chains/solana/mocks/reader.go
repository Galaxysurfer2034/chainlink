// Code generated by mockery v2.13.0-beta.1. DO NOT EDIT.

package mocks

import (
	context "context"

	rpc "github.com/gagliardetto/solana-go/rpc"
	mock "github.com/stretchr/testify/mock"

	solana "github.com/gagliardetto/solana-go"
)

// Reader is an autogenerated mock type for the Reader type
type Reader struct {
	mock.Mock
}

// Balance provides a mock function with given fields: addr
func (_m *Reader) Balance(addr solana.PublicKey) (uint64, error) {
	ret := _m.Called(addr)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(solana.PublicKey) uint64); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(solana.PublicKey) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainID provides a mock function with given fields:
func (_m *Reader) ChainID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountInfoWithOpts provides a mock function with given fields: ctx, addr, opts
func (_m *Reader) GetAccountInfoWithOpts(ctx context.Context, addr solana.PublicKey, opts *rpc.GetAccountInfoOpts) (*rpc.GetAccountInfoResult, error) {
	ret := _m.Called(ctx, addr, opts)

	var r0 *rpc.GetAccountInfoResult
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, *rpc.GetAccountInfoOpts) *rpc.GetAccountInfoResult); ok {
		r0 = rf(ctx, addr, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetAccountInfoResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, *rpc.GetAccountInfoOpts) error); ok {
		r1 = rf(ctx, addr, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFeeForMessage provides a mock function with given fields: msg
func (_m *Reader) GetFeeForMessage(msg string) (uint64, error) {
	ret := _m.Called(msg)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(string) uint64); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LatestBlockhash provides a mock function with given fields:
func (_m *Reader) LatestBlockhash() (*rpc.GetLatestBlockhashResult, error) {
	ret := _m.Called()

	var r0 *rpc.GetLatestBlockhashResult
	if rf, ok := ret.Get(0).(func() *rpc.GetLatestBlockhashResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetLatestBlockhashResult)
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

// SlotHeight provides a mock function with given fields:
func (_m *Reader) SlotHeight() (uint64, error) {
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

type NewReaderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewReader creates a new instance of Reader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReader(t NewReaderT) *Reader {
	mock := &Reader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
