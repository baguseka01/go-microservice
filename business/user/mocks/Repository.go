package mocks

import (
	user "github.com/baguseka01/golang_microservice_hexagonal/business/user"

	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

// FindAllUser provides a mock function with given fields: skip, rowPerPage
func (_m *Repository) FindAllUser(skip int, rowPerPage int) ([]user.User, error) {
	ret := _m.Called(skip, rowPerPage)

	var r0 []user.User
	if rf, ok := ret.Get(0).(func(int, int) []user.User); ok {
		r0 = rf(skip, rowPerPage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(skip, rowPerPage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1

}

// FindUserByID provides a mock function with given fields: id
func (_m Repository) FindUserByID(id int) (*user.User, error) {
	ret := _m.Called(id)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(int) *user.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByUsernameAndPassword provides a mock function with given fields: username, password
func (_m Repository) FindUserByUsernameAndPassword(username string, password string) (*user.User, error) {
	ret := _m.Called(username, password)

	var r0 *user.User
	if rf, ok := ret.Get(0).(func(string, string) *user.User); ok {
		r0 = rf(username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: _a0
func (_m Repository) InsertUser(_a0 user.User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: _a0, currentVersion
func (_m Repository) UpdateUser(_a0 user.User, currentVersion int) error {
	ret := _m.Called(_a0, currentVersion)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.User, int) error); ok {
		r0 = rf(_a0, currentVersion)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}