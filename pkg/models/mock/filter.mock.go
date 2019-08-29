// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/eventstore-lib/pkg/models (interfaces: Filter,EventFilter)

// Package mock is a generated GoMock package.
package mock

import (
	models "github.com/caos/eventstore-lib/pkg/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFilter is a mock of Filter interface
type MockFilter struct {
	ctrl     *gomock.Controller
	recorder *MockFilterMockRecorder
}

// MockFilterMockRecorder is the mock recorder for MockFilter
type MockFilterMockRecorder struct {
	mock *MockFilter
}

// NewMockFilter creates a new mock instance
func NewMockFilter(ctrl *gomock.Controller) *MockFilter {
	mock := &MockFilter{ctrl: ctrl}
	mock.recorder = &MockFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilter) EXPECT() *MockFilterMockRecorder {
	return m.recorder
}

// GetField mocks base method
func (m *MockFilter) GetField() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetField")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetField indicates an expected call of GetField
func (mr *MockFilterMockRecorder) GetField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetField", reflect.TypeOf((*MockFilter)(nil).GetField))
}

// GetOperation mocks base method
func (m *MockFilter) GetOperation() models.Operation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation")
	ret0, _ := ret[0].(models.Operation)
	return ret0
}

// GetOperation indicates an expected call of GetOperation
func (mr *MockFilterMockRecorder) GetOperation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockFilter)(nil).GetOperation))
}

// GetValue mocks base method
func (m *MockFilter) GetValue() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetValue indicates an expected call of GetValue
func (mr *MockFilterMockRecorder) GetValue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockFilter)(nil).GetValue))
}

// MockEventFilter is a mock of EventFilter interface
type MockEventFilter struct {
	ctrl     *gomock.Controller
	recorder *MockEventFilterMockRecorder
}

// MockEventFilterMockRecorder is the mock recorder for MockEventFilter
type MockEventFilterMockRecorder struct {
	mock *MockEventFilter
}

// NewMockEventFilter creates a new mock instance
func NewMockEventFilter(ctrl *gomock.Controller) *MockEventFilter {
	mock := &MockEventFilter{ctrl: ctrl}
	mock.recorder = &MockEventFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventFilter) EXPECT() *MockEventFilterMockRecorder {
	return m.recorder
}

// GetAggregate mocks base method
func (m *MockEventFilter) GetAggregate() models.Aggregate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAggregate")
	ret0, _ := ret[0].(models.Aggregate)
	return ret0
}

// GetAggregate indicates an expected call of GetAggregate
func (mr *MockEventFilterMockRecorder) GetAggregate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAggregate", reflect.TypeOf((*MockEventFilter)(nil).GetAggregate))
}

// GetFilters mocks base method
func (m *MockEventFilter) GetFilters() []models.Filter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilters")
	ret0, _ := ret[0].([]models.Filter)
	return ret0
}

// GetFilters indicates an expected call of GetFilters
func (mr *MockEventFilterMockRecorder) GetFilters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilters", reflect.TypeOf((*MockEventFilter)(nil).GetFilters))
}