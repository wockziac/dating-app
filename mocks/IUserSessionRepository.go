// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	core "datingapp/application/core"

	mock "github.com/stretchr/testify/mock"
)

// IUserSessionRepository is an autogenerated mock type for the IUserSessionRepository type
type IUserSessionRepository struct {
	mock.Mock
}

// GetUserSessionsByPhone provides a mock function with given fields: phoneNumber
func (_m *IUserSessionRepository) GetUserSessionsByPhone(phoneNumber string) ([]core.UserSession, error) {
	ret := _m.Called(phoneNumber)

	if len(ret) == 0 {
		panic("no return value specified for GetUserSessionsByPhone")
	}

	var r0 []core.UserSession
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]core.UserSession, error)); ok {
		return rf(phoneNumber)
	}
	if rf, ok := ret.Get(0).(func(string) []core.UserSession); ok {
		r0 = rf(phoneNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.UserSession)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phoneNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitiateLogin provides a mock function with given fields: session
func (_m *IUserSessionRepository) InitiateLogin(session core.UserSession) error {
	ret := _m.Called(session)

	if len(ret) == 0 {
		panic("no return value specified for InitiateLogin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(core.UserSession) error); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: session
func (_m *IUserSessionRepository) Login(session core.UserSession) (core.UserSession, error) {
	ret := _m.Called(session)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 core.UserSession
	var r1 error
	if rf, ok := ret.Get(0).(func(core.UserSession) (core.UserSession, error)); ok {
		return rf(session)
	}
	if rf, ok := ret.Get(0).(func(core.UserSession) core.UserSession); ok {
		r0 = rf(session)
	} else {
		r0 = ret.Get(0).(core.UserSession)
	}

	if rf, ok := ret.Get(1).(func(core.UserSession) error); ok {
		r1 = rf(session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIUserSessionRepository creates a new instance of IUserSessionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserSessionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserSessionRepository {
	mock := &IUserSessionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
