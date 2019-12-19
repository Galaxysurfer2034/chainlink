// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "chainlink/core/store/models"
import services "chainlink/core/services"

// DeviationCheckerFactory is an autogenerated mock type for the DeviationCheckerFactory type
type DeviationCheckerFactory struct {
	mock.Mock
}

// New provides a mock function with given fields: _a0, _a1
func (_m *DeviationCheckerFactory) New(_a0 models.Initiator, _a1 services.RunManager) (services.DeviationChecker, error) {
	ret := _m.Called(_a0, _a1)

	var r0 services.DeviationChecker
	if rf, ok := ret.Get(0).(func(models.Initiator, services.RunManager) services.DeviationChecker); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(services.DeviationChecker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Initiator, services.RunManager) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
