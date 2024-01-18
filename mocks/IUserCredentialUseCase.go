// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	core "datingapp/application/core"

	mock "github.com/stretchr/testify/mock"
)

// IUserCredentialUseCase is an autogenerated mock type for the IUserCredentialUseCase type
type IUserCredentialUseCase struct {
	mock.Mock
}

// RegisterUserCredential provides a mock function with given fields: request
func (_m *IUserCredentialUseCase) RegisterUserCredential(request core.UserCredentialRequest) error {
	ret := _m.Called(request)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUserCredential")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(core.UserCredentialRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIUserCredentialUseCase creates a new instance of IUserCredentialUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserCredentialUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserCredentialUseCase {
	mock := &IUserCredentialUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}