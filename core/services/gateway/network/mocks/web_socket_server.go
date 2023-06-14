// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// WebSocketServer is an autogenerated mock type for the WebSocketServer type
type WebSocketServer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *WebSocketServer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPort provides a mock function with given fields:
func (_m *WebSocketServer) GetPort() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *WebSocketServer) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewWebSocketServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewWebSocketServer creates a new instance of WebSocketServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWebSocketServer(t mockConstructorTestingTNewWebSocketServer) *WebSocketServer {
	mock := &WebSocketServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
