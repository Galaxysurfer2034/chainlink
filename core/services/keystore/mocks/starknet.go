// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	starkkey "github.com/smartcontractkit/chainlink/core/services/keystore/keys/starkkey"
	mock "github.com/stretchr/testify/mock"
)

// StarkNet is an autogenerated mock type for the StarkNet type
type StarkNet struct {
	mock.Mock
}

// Add provides a mock function with given fields: key
func (_m *StarkNet) Add(key starkkey.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(starkkey.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields:
func (_m *StarkNet) Create() (starkkey.Key, error) {
	ret := _m.Called()

	var r0 starkkey.Key
	if rf, ok := ret.Get(0).(func() starkkey.Key); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(starkkey.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *StarkNet) Delete(id string) (starkkey.Key, error) {
	ret := _m.Called(id)

	var r0 starkkey.Key
	if rf, ok := ret.Get(0).(func(string) starkkey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(starkkey.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields:
func (_m *StarkNet) EnsureKey() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *StarkNet) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *StarkNet) Get(id string) (starkkey.Key, error) {
	ret := _m.Called(id)

	var r0 starkkey.Key
	if rf, ok := ret.Get(0).(func(string) starkkey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(starkkey.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *StarkNet) GetAll() ([]starkkey.Key, error) {
	ret := _m.Called()

	var r0 []starkkey.Key
	if rf, ok := ret.Get(0).(func() []starkkey.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]starkkey.Key)
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

// Import provides a mock function with given fields: keyJSON, password
func (_m *StarkNet) Import(keyJSON []byte, password string) (starkkey.Key, error) {
	ret := _m.Called(keyJSON, password)

	var r0 starkkey.Key
	if rf, ok := ret.Get(0).(func([]byte, string) starkkey.Key); ok {
		r0 = rf(keyJSON, password)
	} else {
		r0 = ret.Get(0).(starkkey.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStarkNet interface {
	mock.TestingT
	Cleanup(func())
}

// NewStarkNet creates a new instance of StarkNet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStarkNet(t mockConstructorTestingTNewStarkNet) *StarkNet {
	mock := &StarkNet{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
