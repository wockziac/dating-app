// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	core "datingapp/application/core"

	mock "github.com/stretchr/testify/mock"
)

// IProfileUseCase is an autogenerated mock type for the IProfileUseCase type
type IProfileUseCase struct {
	mock.Mock
}

// RegisterProfile provides a mock function with given fields: profile
func (_m *IProfileUseCase) RegisterProfile(profile core.Profile) error {
	ret := _m.Called(profile)

	if len(ret) == 0 {
		panic("no return value specified for RegisterProfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Profile) error); ok {
		r0 = rf(profile)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIProfileUseCase creates a new instance of IProfileUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProfileUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProfileUseCase {
	mock := &IProfileUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}