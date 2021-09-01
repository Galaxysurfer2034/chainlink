// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	config "github.com/smartcontractkit/chainlink/core/chains/evm/config"
	bulletprooftxmanager "github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"

	eth "github.com/smartcontractkit/chainlink/core/services/eth"

	log "github.com/smartcontractkit/chainlink/core/services/log"

	logger "github.com/smartcontractkit/chainlink/core/logger"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/core/services/headtracker/types"
)

// Chain is an autogenerated mock type for the Chain type
type Chain struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *Chain) Client() eth.Client {
	ret := _m.Called()

	var r0 eth.Client
	if rf, ok := ret.Get(0).(func() eth.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(eth.Client)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Chain) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Config provides a mock function with given fields:
func (_m *Chain) Config() config.ChainScopedConfig {
	ret := _m.Called()

	var r0 config.ChainScopedConfig
	if rf, ok := ret.Get(0).(func() config.ChainScopedConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.ChainScopedConfig)
		}
	}

	return r0
}

// HeadBroadcaster provides a mock function with given fields:
func (_m *Chain) HeadBroadcaster() types.HeadBroadcaster {
	ret := _m.Called()

	var r0 types.HeadBroadcaster
	if rf, ok := ret.Get(0).(func() types.HeadBroadcaster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.HeadBroadcaster)
		}
	}

	return r0
}

// HeadTracker provides a mock function with given fields:
func (_m *Chain) HeadTracker() types.Tracker {
	ret := _m.Called()

	var r0 types.Tracker
	if rf, ok := ret.Get(0).(func() types.Tracker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Tracker)
		}
	}

	return r0
}

// Healthy provides a mock function with given fields:
func (_m *Chain) Healthy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Chain) ID() *big.Int {
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

// IsArbitrum provides a mock function with given fields:
func (_m *Chain) IsArbitrum() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsL2 provides a mock function with given fields:
func (_m *Chain) IsL2() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsOptimism provides a mock function with given fields:
func (_m *Chain) IsOptimism() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LogBroadcaster provides a mock function with given fields:
func (_m *Chain) LogBroadcaster() log.Broadcaster {
	ret := _m.Called()

	var r0 log.Broadcaster
	if rf, ok := ret.Get(0).(func() log.Broadcaster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Broadcaster)
		}
	}

	return r0
}

// Logger provides a mock function with given fields:
func (_m *Chain) Logger() *logger.Logger {
	ret := _m.Called()

	var r0 *logger.Logger
	if rf, ok := ret.Get(0).(func() *logger.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*logger.Logger)
		}
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *Chain) Ready() error {
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
func (_m *Chain) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxManager provides a mock function with given fields:
func (_m *Chain) TxManager() bulletprooftxmanager.TxManager {
	ret := _m.Called()

	var r0 bulletprooftxmanager.TxManager
	if rf, ok := ret.Get(0).(func() bulletprooftxmanager.TxManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bulletprooftxmanager.TxManager)
		}
	}

	return r0
}
