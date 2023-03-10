// Code generated by MockGen. DO NOT EDIT.
// Source: file_storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockFileStorage is a mock of FileStorage interface
type MockFileStorage struct {
	ctrl     *gomock.Controller
	recorder *MockFileStorageMockRecorder
}

// MockFileStorageMockRecorder is the mock recorder for MockFileStorage
type MockFileStorageMockRecorder struct {
	mock *MockFileStorage
}

// NewMockFileStorage creates a new mock instance
func NewMockFileStorage(ctrl *gomock.Controller) *MockFileStorage {
	mock := &MockFileStorage{ctrl: ctrl}
	mock.recorder = &MockFileStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileStorage) EXPECT() *MockFileStorageMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *MockFileStorage) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockFileStorageMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockFileStorage)(nil).Ping), ctx)
}

// Read mocks base method
func (m *MockFileStorage) Read(ctx context.Context, path string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, path)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockFileStorageMockRecorder) Read(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileStorage)(nil).Read), ctx, path)
}

// Write mocks base method
func (m *MockFileStorage) Write(ctx context.Context, data io.Reader, size int64, path, contentType string, isPublicRead bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, data, size, path, contentType, isPublicRead)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockFileStorageMockRecorder) Write(ctx, data, size, path, contentType, isPublicRead interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileStorage)(nil).Write), ctx, data, size, path, contentType, isPublicRead)
}

// DeleteData mocks base method
func (m *MockFileStorage) DeleteData(ctx context.Context, address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteData", ctx, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteData indicates an expected call of DeleteData
func (mr *MockFileStorageMockRecorder) DeleteData(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteData", reflect.TypeOf((*MockFileStorage)(nil).DeleteData), ctx, address)
}
