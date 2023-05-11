// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	context "context"

	commontypes "github.com/smartcontractkit/chainlink/v2/common/types"

	logger "github.com/smartcontractkit/chainlink/v2/core/logger"

	mock "github.com/stretchr/testify/mock"

	types "github.com/smartcontractkit/chainlink/v2/common/txmgr/types"
)

// TxAttemptBuilder is an autogenerated mock type for the TxAttemptBuilder type
type TxAttemptBuilder[CHAIN_ID types.ID, HEAD commontypes.Head[TX_HASH], ADDR commontypes.Hashable, TX_HASH commontypes.Hashable, BLOCK_HASH commontypes.Hashable, R types.ChainReceipt[TX_HASH, BLOCK_HASH], SEQ types.Sequence, FEE types.Fee, ADD interface{}] struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HealthReport provides a mock function with given fields:
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) HealthReport() map[string]error {
	ret := _m.Called()

	var r0 map[string]error
	if rf, ok := ret.Get(0).(func() map[string]error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]error)
		}
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewBumpTxAttempt provides a mock function with given fields: ctx, tx, previousAttempt, priorAttempts, lggr
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) NewBumpTxAttempt(ctx context.Context, tx types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], previousAttempt types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], priorAttempts []types.PriorAttempt[FEE, TX_HASH], lggr logger.Logger) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, bool, error) {
	ret := _m.Called(ctx, tx, previousAttempt, priorAttempts, lggr)

	var r0 types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]
	var r1 FEE
	var r2 uint32
	var r3 bool
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], []types.PriorAttempt[FEE, TX_HASH], logger.Logger) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, bool, error)); ok {
		return rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], []types.PriorAttempt[FEE, TX_HASH], logger.Logger) types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]); ok {
		r0 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r0 = ret.Get(0).(types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD])
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], []types.PriorAttempt[FEE, TX_HASH], logger.Logger) FEE); ok {
		r1 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r1 = ret.Get(1).(FEE)
	}

	if rf, ok := ret.Get(2).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], []types.PriorAttempt[FEE, TX_HASH], logger.Logger) uint32); ok {
		r2 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r2 = ret.Get(2).(uint32)
	}

	if rf, ok := ret.Get(3).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], []types.PriorAttempt[FEE, TX_HASH], logger.Logger) bool); ok {
		r3 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r3 = ret.Get(3).(bool)
	}

	if rf, ok := ret.Get(4).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], []types.PriorAttempt[FEE, TX_HASH], logger.Logger) error); ok {
		r4 = rf(ctx, tx, previousAttempt, priorAttempts, lggr)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// NewCustomTxAttempt provides a mock function with given fields: tx, fee, gasLimit, txType, lggr
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) NewCustomTxAttempt(tx types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], fee FEE, gasLimit uint32, txType int, lggr logger.Logger) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], bool, error) {
	ret := _m.Called(tx, fee, gasLimit, txType, lggr)

	var r0 types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, int, logger.Logger) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], bool, error)); ok {
		return rf(tx, fee, gasLimit, txType, lggr)
	}
	if rf, ok := ret.Get(0).(func(types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, int, logger.Logger) types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]); ok {
		r0 = rf(tx, fee, gasLimit, txType, lggr)
	} else {
		r0 = ret.Get(0).(types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD])
	}

	if rf, ok := ret.Get(1).(func(types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, int, logger.Logger) bool); ok {
		r1 = rf(tx, fee, gasLimit, txType, lggr)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, int, logger.Logger) error); ok {
		r2 = rf(tx, fee, gasLimit, txType, lggr)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewEmptyTxAttempt provides a mock function with given fields: seq, feeLimit, fee, fromAddress
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) NewEmptyTxAttempt(seq SEQ, feeLimit uint32, fee FEE, fromAddress ADDR) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], error) {
	ret := _m.Called(seq, feeLimit, fee, fromAddress)

	var r0 types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]
	var r1 error
	if rf, ok := ret.Get(0).(func(SEQ, uint32, FEE, ADDR) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], error)); ok {
		return rf(seq, feeLimit, fee, fromAddress)
	}
	if rf, ok := ret.Get(0).(func(SEQ, uint32, FEE, ADDR) types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]); ok {
		r0 = rf(seq, feeLimit, fee, fromAddress)
	} else {
		r0 = ret.Get(0).(types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD])
	}

	if rf, ok := ret.Get(1).(func(SEQ, uint32, FEE, ADDR) error); ok {
		r1 = rf(seq, feeLimit, fee, fromAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTxAttempt provides a mock function with given fields: ctx, tx, lggr, opts
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) NewTxAttempt(ctx context.Context, tx types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], lggr logger.Logger, opts ...types.Opt) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, bool, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, tx, lggr)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]
	var r1 FEE
	var r2 uint32
	var r3 bool
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, ...types.Opt) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, bool, error)); ok {
		return rf(ctx, tx, lggr, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, ...types.Opt) types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]); ok {
		r0 = rf(ctx, tx, lggr, opts...)
	} else {
		r0 = ret.Get(0).(types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD])
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, ...types.Opt) FEE); ok {
		r1 = rf(ctx, tx, lggr, opts...)
	} else {
		r1 = ret.Get(1).(FEE)
	}

	if rf, ok := ret.Get(2).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, ...types.Opt) uint32); ok {
		r2 = rf(ctx, tx, lggr, opts...)
	} else {
		r2 = ret.Get(2).(uint32)
	}

	if rf, ok := ret.Get(3).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, ...types.Opt) bool); ok {
		r3 = rf(ctx, tx, lggr, opts...)
	} else {
		r3 = ret.Get(3).(bool)
	}

	if rf, ok := ret.Get(4).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, ...types.Opt) error); ok {
		r4 = rf(ctx, tx, lggr, opts...)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// NewTxAttemptWithType provides a mock function with given fields: ctx, tx, lggr, txType, opts
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) NewTxAttemptWithType(ctx context.Context, tx types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], lggr logger.Logger, txType int, opts ...types.Opt) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, bool, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, tx, lggr, txType)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]
	var r1 FEE
	var r2 uint32
	var r3 bool
	var r4 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, int, ...types.Opt) (types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], FEE, uint32, bool, error)); ok {
		return rf(ctx, tx, lggr, txType, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, int, ...types.Opt) types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD]); ok {
		r0 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r0 = ret.Get(0).(types.TxAttempt[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD])
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, int, ...types.Opt) FEE); ok {
		r1 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r1 = ret.Get(1).(FEE)
	}

	if rf, ok := ret.Get(2).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, int, ...types.Opt) uint32); ok {
		r2 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r2 = ret.Get(2).(uint32)
	}

	if rf, ok := ret.Get(3).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, int, ...types.Opt) bool); ok {
		r3 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r3 = ret.Get(3).(bool)
	}

	if rf, ok := ret.Get(4).(func(context.Context, types.Tx[CHAIN_ID, ADDR, TX_HASH, BLOCK_HASH, R, FEE, ADD], logger.Logger, int, ...types.Opt) error); ok {
		r4 = rf(ctx, tx, lggr, txType, opts...)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// OnNewLongestChain provides a mock function with given fields: ctx, head
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) OnNewLongestChain(ctx context.Context, head HEAD) {
	_m.Called(ctx, head)
}

// Ready provides a mock function with given fields:
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) Ready() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTxAttemptBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewTxAttemptBuilder creates a new instance of TxAttemptBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTxAttemptBuilder[CHAIN_ID types.ID, HEAD commontypes.Head[TX_HASH], ADDR commontypes.Hashable, TX_HASH commontypes.Hashable, BLOCK_HASH commontypes.Hashable, R types.ChainReceipt[TX_HASH, BLOCK_HASH], SEQ types.Sequence, FEE types.Fee, ADD interface{}](t mockConstructorTestingTNewTxAttemptBuilder) *TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD] {
	mock := &TxAttemptBuilder[CHAIN_ID, HEAD, ADDR, TX_HASH, BLOCK_HASH, R, SEQ, FEE, ADD]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
