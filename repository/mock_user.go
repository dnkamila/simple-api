// Automatically generated by MockGen. DO NOT EDIT!
// Source: repository/user.go

package repository

import (
	gomock "github.com/golang/mock/gomock"
	. "simple-api/models"
)

// Mock of UserRepositoryInterface interface
type MockUserRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *_MockUserRepositoryInterfaceRecorder
}

// Recorder for MockUserRepositoryInterface (not exported)
type _MockUserRepositoryInterfaceRecorder struct {
	mock *MockUserRepositoryInterface
}

func NewMockUserRepositoryInterface(ctrl *gomock.Controller) *MockUserRepositoryInterface {
	mock := &MockUserRepositoryInterface{ctrl: ctrl}
	mock.recorder = &_MockUserRepositoryInterfaceRecorder{mock}
	return mock
}

func (_m *MockUserRepositoryInterface) EXPECT() *_MockUserRepositoryInterfaceRecorder {
	return _m.recorder
}

func (_m *MockUserRepositoryInterface) CreateUser(user *User) (*User, error) {
	ret := _m.ctrl.Call(_m, "CreateUser", user)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockUserRepositoryInterfaceRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateUser", arg0)
}

func (_m *MockUserRepositoryInterface) GetUserById(user *User) (*User, error) {
	ret := _m.ctrl.Call(_m, "GetUserById", user)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockUserRepositoryInterfaceRecorder) GetUserById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserById", arg0)
}

func (_m *MockUserRepositoryInterface) GetUserByUsername(user *User) (*User, error) {
	ret := _m.ctrl.Call(_m, "GetUserByUsername", user)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockUserRepositoryInterfaceRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUserByUsername", arg0)
}

func (_m *MockUserRepositoryInterface) UpdateUserPasswordById(user *User) (*User, error) {
	ret := _m.ctrl.Call(_m, "UpdateUserPasswordById", user)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockUserRepositoryInterfaceRecorder) UpdateUserPasswordById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateUserPasswordById", arg0)
}

func (_m *MockUserRepositoryInterface) UpdateUserPasswordByUsername(user *User) (*User, error) {
	ret := _m.ctrl.Call(_m, "UpdateUserPasswordByUsername", user)
	ret0, _ := ret[0].(*User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockUserRepositoryInterfaceRecorder) UpdateUserPasswordByUsername(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateUserPasswordByUsername", arg0)
}

func (_m *MockUserRepositoryInterface) DeleteUserById(user *User) error {
	ret := _m.ctrl.Call(_m, "DeleteUserById", user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUserRepositoryInterfaceRecorder) DeleteUserById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUserById", arg0)
}

func (_m *MockUserRepositoryInterface) DeleteUserByUsername(user *User) error {
	ret := _m.ctrl.Call(_m, "DeleteUserByUsername", user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockUserRepositoryInterfaceRecorder) DeleteUserByUsername(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteUserByUsername", arg0)
}