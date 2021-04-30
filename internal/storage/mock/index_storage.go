// Code generated by MockGen. DO NOT EDIT.
// Source: index_storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	storage "github.com/Decentr-net/cerberus/internal/storage"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIndexStorage is a mock of IndexStorage interface
type MockIndexStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIndexStorageMockRecorder
}

// MockIndexStorageMockRecorder is the mock recorder for MockIndexStorage
type MockIndexStorageMockRecorder struct {
	mock *MockIndexStorage
}

// NewMockIndexStorage creates a new mock instance
func NewMockIndexStorage(ctrl *gomock.Controller) *MockIndexStorage {
	mock := &MockIndexStorage{ctrl: ctrl}
	mock.recorder = &MockIndexStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexStorage) EXPECT() *MockIndexStorageMockRecorder {
	return m.recorder
}

// GetProfile mocks base method
func (m *MockIndexStorage) GetProfile(ctx context.Context, addr string) (*storage.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", ctx, addr)
	ret0, _ := ret[0].(*storage.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile
func (mr *MockIndexStorageMockRecorder) GetProfile(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockIndexStorage)(nil).GetProfile), ctx, addr)
}

// GetProfiles mocks base method
func (m *MockIndexStorage) GetProfiles(ctx context.Context, addr []string) ([]*storage.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfiles", ctx, addr)
	ret0, _ := ret[0].([]*storage.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfiles indicates an expected call of GetProfiles
func (mr *MockIndexStorageMockRecorder) GetProfiles(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfiles", reflect.TypeOf((*MockIndexStorage)(nil).GetProfiles), ctx, addr)
}

// SetProfile mocks base method
func (m *MockIndexStorage) SetProfile(ctx context.Context, p *storage.SetProfileParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProfile", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProfile indicates an expected call of SetProfile
func (mr *MockIndexStorageMockRecorder) SetProfile(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProfile", reflect.TypeOf((*MockIndexStorage)(nil).SetProfile), ctx, p)
}
