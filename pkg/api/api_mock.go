// Code generated by MockGen. DO NOT EDIT.
// Source: api.go

// Package api is a generated GoMock package.
package api

import (
	context "context"
	json "encoding/json"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCerberus is a mock of Cerberus interface
type MockCerberus struct {
	ctrl     *gomock.Controller
	recorder *MockCerberusMockRecorder
}

// MockCerberusMockRecorder is the mock recorder for MockCerberus
type MockCerberusMockRecorder struct {
	mock *MockCerberus
}

// NewMockCerberus creates a new mock instance
func NewMockCerberus(ctrl *gomock.Controller) *MockCerberus {
	mock := &MockCerberus{ctrl: ctrl}
	mock.recorder = &MockCerberusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCerberus) EXPECT() *MockCerberusMockRecorder {
	return m.recorder
}

// SendPDV mocks base method
func (m *MockCerberus) SendPDV(ctx context.Context, data []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPDV", ctx, data)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendPDV indicates an expected call of SendPDV
func (mr *MockCerberusMockRecorder) SendPDV(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPDV", reflect.TypeOf((*MockCerberus)(nil).SendPDV), ctx, data)
}

// ReceivePDV mocks base method
func (m *MockCerberus) ReceivePDV(ctx context.Context, address string) (json.RawMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceivePDV", ctx, address)
	ret0, _ := ret[0].(json.RawMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceivePDV indicates an expected call of ReceivePDV
func (mr *MockCerberusMockRecorder) ReceivePDV(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivePDV", reflect.TypeOf((*MockCerberus)(nil).ReceivePDV), ctx, address)
}

// DoesPDVExist mocks base method
func (m *MockCerberus) DoesPDVExist(ctx context.Context, address string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesPDVExist", ctx, address)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoesPDVExist indicates an expected call of DoesPDVExist
func (mr *MockCerberusMockRecorder) DoesPDVExist(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesPDVExist", reflect.TypeOf((*MockCerberus)(nil).DoesPDVExist), ctx, address)
}
