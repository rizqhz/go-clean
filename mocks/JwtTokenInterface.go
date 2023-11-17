// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v5"

	mock "github.com/stretchr/testify/mock"
)

// JwtTokenInterface is an autogenerated mock type for the JwtTokenInterface type
type JwtTokenInterface struct {
	mock.Mock
}

// ExtractToken provides a mock function with given fields: token
func (_m *JwtTokenInterface) ExtractToken(token *jwt.Token) (*int, error) {
	ret := _m.Called(token)

	var r0 *int
	var r1 error
	if rf, ok := ret.Get(0).(func(*jwt.Token) (*int, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(*jwt.Token) *int); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int)
		}
	}

	if rf, ok := ret.Get(1).(func(*jwt.Token) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateToken provides a mock function with given fields: id
func (_m *JwtTokenInterface) GenerateToken(id uint) *string {
	ret := _m.Called(id)

	var r0 *string
	if rf, ok := ret.Get(0).(func(uint) *string); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// NewJwtTokenInterface creates a new instance of JwtTokenInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJwtTokenInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JwtTokenInterface {
	mock := &JwtTokenInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
