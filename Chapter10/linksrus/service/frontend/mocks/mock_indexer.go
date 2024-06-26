// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ibiscum/Hands-On-Software-Engineering-with-Golang/Chapter06/textindexer/index (interfaces: Iterator)

// Package mocks is a generated GoMock package.
package mocks

import (
	index "github.com/ibiscum/Hands-On-Software-Engineering-with-Golang/Chapter06/textindexer/index"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIterator is a mock of Iterator interface
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIterator)(nil).Close))
}

// Document mocks base method
func (m *MockIterator) Document() *index.Document {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Document")
	ret0, _ := ret[0].(*index.Document)
	return ret0
}

// Document indicates an expected call of Document
func (mr *MockIteratorMockRecorder) Document() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Document", reflect.TypeOf((*MockIterator)(nil).Document))
}

// Error mocks base method
func (m *MockIterator) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockIteratorMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockIterator)(nil).Error))
}

// Next mocks base method
func (m *MockIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next))
}

// TotalCount mocks base method
func (m *MockIterator) TotalCount() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalCount")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// TotalCount indicates an expected call of TotalCount
func (mr *MockIteratorMockRecorder) TotalCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalCount", reflect.TypeOf((*MockIterator)(nil).TotalCount))
}
