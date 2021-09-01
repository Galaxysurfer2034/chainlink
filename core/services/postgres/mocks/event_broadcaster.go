// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	postgres "github.com/smartcontractkit/chainlink/core/services/postgres"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// EventBroadcaster is an autogenerated mock type for the EventBroadcaster type
type EventBroadcaster struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *EventBroadcaster) Close() error {
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
func (_m *EventBroadcaster) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Notify provides a mock function with given fields: channel, payload
func (_m *EventBroadcaster) Notify(channel string, payload string) error {
	ret := _m.Called(channel, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(channel, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotifyInsideGormTx provides a mock function with given fields: tx, channel, payload
func (_m *EventBroadcaster) NotifyInsideGormTx(tx *gorm.DB, channel string, payload string) error {
	ret := _m.Called(tx, channel, payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, string) error); ok {
		r0 = rf(tx, channel, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *EventBroadcaster) Ready() error {
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
func (_m *EventBroadcaster) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: channel, payloadFilter
func (_m *EventBroadcaster) Subscribe(channel string, payloadFilter string) (postgres.Subscription, error) {
	ret := _m.Called(channel, payloadFilter)

	var r0 postgres.Subscription
	if rf, ok := ret.Get(0).(func(string, string) postgres.Subscription); ok {
		r0 = rf(channel, payloadFilter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(postgres.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(channel, payloadFilter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
