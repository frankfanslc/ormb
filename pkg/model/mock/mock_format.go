// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/model/format.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// ValidateDirectory mocks base method
func (m *MockInterface) ValidateDirectory(rootPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDirectory", rootPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateDirectory indicates an expected call of ValidateDirectory
func (mr *MockInterfaceMockRecorder) ValidateDirectory(rootPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDirectory", reflect.TypeOf((*MockInterface)(nil).ValidateDirectory), rootPath)
}
