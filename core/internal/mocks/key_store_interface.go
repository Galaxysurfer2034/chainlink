// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	big "math/big"

	accounts "github.com/ethereum/go-ethereum/accounts"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// KeyStoreInterface is an autogenerated mock type for the KeyStoreInterface type
type KeyStoreInterface struct {
	mock.Mock
}

// Accounts provides a mock function with given fields:
func (_m *KeyStoreInterface) Accounts() []accounts.Account {
	ret := _m.Called()

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func() []accounts.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	return r0
}

// Export provides a mock function with given fields: a, passphrase, newPassphrase
func (_m *KeyStoreInterface) Export(a accounts.Account, passphrase string, newPassphrase string) ([]byte, error) {
	ret := _m.Called(a, passphrase, newPassphrase)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(accounts.Account, string, string) []byte); ok {
		r0 = rf(a, passphrase, newPassphrase)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accounts.Account, string, string) error); ok {
		r1 = rf(a, passphrase, newPassphrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountByAddress provides a mock function with given fields: _a0
func (_m *KeyStoreInterface) GetAccountByAddress(_a0 common.Address) (accounts.Account, error) {
	ret := _m.Called(_a0)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(common.Address) accounts.Account); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccounts provides a mock function with given fields:
func (_m *KeyStoreInterface) GetAccounts() []accounts.Account {
	ret := _m.Called()

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func() []accounts.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	return r0
}

// HasAccountWithAddress provides a mock function with given fields: _a0
func (_m *KeyStoreInterface) HasAccountWithAddress(_a0 common.Address) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(common.Address) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HasAccounts provides a mock function with given fields:
func (_m *KeyStoreInterface) HasAccounts() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Import provides a mock function with given fields: keyJSON, passphrase, newPassphrase
func (_m *KeyStoreInterface) Import(keyJSON []byte, passphrase string, newPassphrase string) (accounts.Account, error) {
	ret := _m.Called(keyJSON, passphrase, newPassphrase)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func([]byte, string, string) accounts.Account); ok {
		r0 = rf(keyJSON, passphrase, newPassphrase)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string, string) error); ok {
		r1 = rf(keyJSON, passphrase, newPassphrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccount provides a mock function with given fields: passphrase
func (_m *KeyStoreInterface) NewAccount(passphrase string) (accounts.Account, error) {
	ret := _m.Called(passphrase)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(string) accounts.Account); ok {
		r0 = rf(passphrase)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(passphrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignTx provides a mock function with given fields: account, tx, chainID
func (_m *KeyStoreInterface) SignTx(account accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	ret := _m.Called(account, tx, chainID)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(accounts.Account, *types.Transaction, *big.Int) *types.Transaction); ok {
		r0 = rf(account, tx, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(accounts.Account, *types.Transaction, *big.Int) error); ok {
		r1 = rf(account, tx, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlock provides a mock function with given fields: phrase
func (_m *KeyStoreInterface) Unlock(phrase string) error {
	ret := _m.Called(phrase)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(phrase)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Wallets provides a mock function with given fields:
func (_m *KeyStoreInterface) Wallets() []accounts.Wallet {
	ret := _m.Called()

	var r0 []accounts.Wallet
	if rf, ok := ret.Get(0).(func() []accounts.Wallet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Wallet)
		}
	}

	return r0
}
