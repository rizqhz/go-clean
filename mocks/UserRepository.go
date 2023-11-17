// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	entity "github.com/rizghz/clean/module/user/entity"
	mock "github.com/stretchr/testify/mock"

	transfer "github.com/rizghz/clean/module/user/transfer"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *UserRepository) Create(data *entity.User) (bool, *entity.User) {
	ret := _m.Called(data)

	var r0 bool
	var r1 *entity.User
	if rf, ok := ret.Get(0).(func(*entity.User) (bool, *entity.User)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*entity.User) bool); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*entity.User) *entity.User); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*entity.User)
		}
	}

	return r0, r1
}

// Find provides a mock function with given fields: data
func (_m *UserRepository) Find(data *transfer.UserRequestBody) *entity.User {
	ret := _m.Called(data)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(*transfer.UserRequestBody) *entity.User); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	return r0
}

// Get provides a mock function with given fields:
func (_m *UserRepository) Get() []*entity.User {
	ret := _m.Called()

	var r0 []*entity.User
	if rf, ok := ret.Get(0).(func() []*entity.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.User)
		}
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
