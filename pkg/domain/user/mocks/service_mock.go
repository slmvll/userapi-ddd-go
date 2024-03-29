// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	user "userapi-ddd-go/pkg/domain/user"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: _a0
func (_m *UserRepository) AddUser(_a0 user.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AddUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields:
func (_m *UserRepository) GetAllUsers() ([]user.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []user.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]user.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []user.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: _a0
func (_m *UserRepository) GetUser(_a0 uuid.UUID) (user.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (user.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) user.User); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
