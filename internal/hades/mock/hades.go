// Code generated by MockGen. DO NOT EDIT.
// Source: hades.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	hades "github.com/Decentr-net/cerberus/internal/hades"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHades is a mock of Hades interface
type MockHades struct {
	ctrl     *gomock.Controller
	recorder *MockHadesMockRecorder
}

// MockHadesMockRecorder is the mock recorder for MockHades
type MockHadesMockRecorder struct {
	mock *MockHades
}

// NewMockHades creates a new mock instance
func NewMockHades(ctrl *gomock.Controller) *MockHades {
	mock := &MockHades{ctrl: ctrl}
	mock.recorder = &MockHadesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHades) EXPECT() *MockHadesMockRecorder {
	return m.recorder
}

// AntiFraud mocks base method
func (m *MockHades) AntiFraud(ctx context.Context, req *hades.AntiFraudRequest) (*hades.AntiFraudResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AntiFraud", ctx, req)
	ret0, _ := ret[0].(*hades.AntiFraudResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AntiFraud indicates an expected call of AntiFraud
func (mr *MockHadesMockRecorder) AntiFraud(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AntiFraud", reflect.TypeOf((*MockHades)(nil).AntiFraud), ctx, req)
}