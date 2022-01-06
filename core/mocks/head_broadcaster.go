// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	context "context"

	eth "github.com/smartcontractkit/chainlink/core/services/eth"
	httypes "github.com/smartcontractkit/chainlink/core/services/headtracker/types"
	mock "github.com/stretchr/testify/mock"
)

// HeadBroadcaster is an autogenerated mock type for the HeadBroadcaster type
type HeadBroadcaster struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *HeadBroadcaster) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *HeadBroadcaster) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *HeadBroadcaster) OnNewLongestChain(ctx context.Context, head *eth.Head) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *HeadBroadcaster) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *HeadBroadcaster) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: callback
func (_m *HeadBroadcaster) Subscribe(callback httypes.HeadTrackable) (*eth.Head, func()) {
	ret := _m.Called(callback)

	var r0 *eth.Head
	if rf, ok := ret.Get(0).(func(httypes.HeadTrackable) *eth.Head); ok {
		r0 = rf(callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eth.Head)
		}
	}

	var r1 func()
	if rf, ok := ret.Get(1).(func(httypes.HeadTrackable) func()); ok {
		r1 = rf(callback)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	return r0, r1
}
