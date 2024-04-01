// Code generated by mockery v2.38.0. DO NOT EDIT.

package sms

import (
	models "github.com/pubgolf/pubgolf/api/internal/lib/models"
	mock "github.com/stretchr/testify/mock"
)

// MockMessenger is an autogenerated mock type for the Messenger type
type MockMessenger struct {
	mock.Mock
}

// CheckVerification provides a mock function with given fields: to, code
func (_m *MockMessenger) CheckVerification(to models.PhoneNum, code string) (bool, error) {
	ret := _m.Called(to, code)

	if len(ret) == 0 {
		panic("no return value specified for CheckVerification")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(models.PhoneNum, string) (bool, error)); ok {
		return rf(to, code)
	}
	if rf, ok := ret.Get(0).(func(models.PhoneNum, string) bool); ok {
		r0 = rf(to, code)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(models.PhoneNum, string) error); ok {
		r1 = rf(to, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendVerification provides a mock function with given fields: to
func (_m *MockMessenger) SendVerification(to models.PhoneNum) error {
	ret := _m.Called(to)

	if len(ret) == 0 {
		panic("no return value specified for SendVerification")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.PhoneNum) error); ok {
		r0 = rf(to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockMessenger creates a new instance of MockMessenger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessenger(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessenger {
	mock := &MockMessenger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
