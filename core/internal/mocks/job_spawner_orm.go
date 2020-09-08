// Code generated by mockery v2.1.0. DO NOT EDIT.

package mocks

import (
	job "github.com/smartcontractkit/chainlink/core/services/pipeline"
	mock "github.com/stretchr/testify/mock"

	models "github.com/smartcontractkit/chainlink/core/store/models"
)

// JobSpawnerORM is an autogenerated mock type for the JobSpawnerORM type
type JobSpawnerORM struct {
	mock.Mock
}

// JobsAsInterfaces provides a mock function with given fields: fn
func (_m *JobSpawnerORM) JobsAsInterfaces(fn func(job.JobSpec) bool) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(func(job.JobSpec) bool) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertErrorFor provides a mock function with given fields: jobID, err
func (_m *JobSpawnerORM) UpsertErrorFor(jobID *models.ID, err string) {
	_m.Called(jobID, err)
}
