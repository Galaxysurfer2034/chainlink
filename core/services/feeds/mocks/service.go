// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	feeds "github.com/smartcontractkit/chainlink/core/services/feeds"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CountManagerServices provides a mock function with given fields:
func (_m *Service) CountManagerServices() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetManagerService provides a mock function with given fields: id
func (_m *Service) GetManagerService(id int32) (*feeds.ManagerService, error) {
	ret := _m.Called(id)

	var r0 *feeds.ManagerService
	if rf, ok := ret.Get(0).(func(int32) *feeds.ManagerService); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*feeds.ManagerService)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListManagerServices provides a mock function with given fields:
func (_m *Service) ListManagerServices() ([]feeds.ManagerService, error) {
	ret := _m.Called()

	var r0 []feeds.ManagerService
	if rf, ok := ret.Get(0).(func() []feeds.ManagerService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]feeds.ManagerService)
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

// RegisterManagerService provides a mock function with given fields: ms
func (_m *Service) RegisterManagerService(ms *feeds.ManagerService) (int32, error) {
	ret := _m.Called(ms)

	var r0 int32
	if rf, ok := ret.Get(0).(func(*feeds.ManagerService) int32); ok {
		r0 = rf(ms)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*feeds.ManagerService) error); ok {
		r1 = rf(ms)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
