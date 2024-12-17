// Code generated by MockGen. DO NOT EDIT.
// Source: operations.go
//
// Generated by this command:
//
//	mockgen -source=operations.go -package=repository -destination=operations_mock.go
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockOperations is a mock of Operations interface.
type MockOperations struct {
	ctrl     *gomock.Controller
	recorder *MockOperationsMockRecorder
	isgomock struct{}
}

// MockOperationsMockRecorder is the mock recorder for MockOperations.
type MockOperationsMockRecorder struct {
	mock *MockOperations
}

// NewMockOperations creates a new mock instance.
func NewMockOperations(ctrl *gomock.Controller) *MockOperations {
	mock := &MockOperations{ctrl: ctrl}
	mock.recorder = &MockOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperations) EXPECT() *MockOperationsMockRecorder {
	return m.recorder
}

// InsertOne mocks base method.
func (m *MockOperations) InsertOne(arg0 context.Context, arg1 any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", arg0, arg1)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockOperationsMockRecorder) InsertOne(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockOperations)(nil).InsertOne), arg0, arg1)
}
