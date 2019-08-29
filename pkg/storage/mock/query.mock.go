// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/eventstore-lib/pkg/storage (interfaces: Query)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/caos/eventstore-lib/pkg/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockQuery is a mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// Condition mocks base method
func (m *MockQuery) Condition(arg0 string, arg1 models.Operation, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Condition", arg0, arg1, arg2)
}

// Condition indicates an expected call of Condition
func (mr *MockQueryMockRecorder) Condition(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Condition", reflect.TypeOf((*MockQuery)(nil).Condition), arg0, arg1, arg2)
}