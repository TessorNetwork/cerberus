// Code generated by MockGen. DO NOT EDIT.
// Source: index_storage.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	entities "github.com/TessorNetwork/cerberus/internal/entities"
	storage "github.com/TessorNetwork/cerberus/internal/storage"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
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

// InTx mocks base method
func (m *MockIndexStorage) InTx(ctx context.Context, f func(storage.IndexStorage) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InTx", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// InTx indicates an expected call of InTx
func (mr *MockIndexStorageMockRecorder) InTx(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InTx", reflect.TypeOf((*MockIndexStorage)(nil).InTx), ctx, f)
}

// SetHeight mocks base method
func (m *MockIndexStorage) SetHeight(ctx context.Context, height uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeight", ctx, height)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeight indicates an expected call of SetHeight
func (mr *MockIndexStorageMockRecorder) SetHeight(ctx, height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeight", reflect.TypeOf((*MockIndexStorage)(nil).SetHeight), ctx, height)
}

// GetHeight mocks base method
func (m *MockIndexStorage) GetHeight(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeight", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeight indicates an expected call of GetHeight
func (mr *MockIndexStorageMockRecorder) GetHeight(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeight", reflect.TypeOf((*MockIndexStorage)(nil).GetHeight), ctx)
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

// SetProfileBanned mocks base method
func (m *MockIndexStorage) SetProfileBanned(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProfileBanned", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProfileBanned indicates an expected call of SetProfileBanned
func (mr *MockIndexStorageMockRecorder) SetProfileBanned(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProfileBanned", reflect.TypeOf((*MockIndexStorage)(nil).SetProfileBanned), ctx, addr)
}

// IsProfileBanned mocks base method
func (m *MockIndexStorage) IsProfileBanned(ctx context.Context, addr string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsProfileBanned", ctx, addr)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsProfileBanned indicates an expected call of IsProfileBanned
func (mr *MockIndexStorageMockRecorder) IsProfileBanned(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsProfileBanned", reflect.TypeOf((*MockIndexStorage)(nil).IsProfileBanned), ctx, addr)
}

// DeleteProfile mocks base method
func (m *MockIndexStorage) DeleteProfile(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile
func (mr *MockIndexStorageMockRecorder) DeleteProfile(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockIndexStorage)(nil).DeleteProfile), ctx, addr)
}

// ListPDV mocks base method
func (m *MockIndexStorage) ListPDV(ctx context.Context, owner string, from uint64, limit uint16) ([]uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPDV", ctx, owner, from, limit)
	ret0, _ := ret[0].([]uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPDV indicates an expected call of ListPDV
func (mr *MockIndexStorageMockRecorder) ListPDV(ctx, owner, from, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPDV", reflect.TypeOf((*MockIndexStorage)(nil).ListPDV), ctx, owner, from, limit)
}

// DeletePDV mocks base method
func (m *MockIndexStorage) DeletePDV(ctx context.Context, owner string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePDV", ctx, owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePDV indicates an expected call of DeletePDV
func (mr *MockIndexStorageMockRecorder) DeletePDV(ctx, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePDV", reflect.TypeOf((*MockIndexStorage)(nil).DeletePDV), ctx, owner)
}

// GetPDVMeta mocks base method
func (m *MockIndexStorage) GetPDVMeta(ctx context.Context, address string, id uint64) (*entities.PDVMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPDVMeta", ctx, address, id)
	ret0, _ := ret[0].(*entities.PDVMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPDVMeta indicates an expected call of GetPDVMeta
func (mr *MockIndexStorageMockRecorder) GetPDVMeta(ctx, address, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPDVMeta", reflect.TypeOf((*MockIndexStorage)(nil).GetPDVMeta), ctx, address, id)
}

// SetPDVMeta mocks base method
func (m_2 *MockIndexStorage) SetPDVMeta(ctx context.Context, address string, id uint64, tx, device string, m *entities.PDVMeta) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SetPDVMeta", ctx, address, id, tx, device, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPDVMeta indicates an expected call of SetPDVMeta
func (mr *MockIndexStorageMockRecorder) SetPDVMeta(ctx, address, id, tx, device, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPDVMeta", reflect.TypeOf((*MockIndexStorage)(nil).SetPDVMeta), ctx, address, id, tx, device, m)
}

// GetPDVDelta mocks base method
func (m *MockIndexStorage) GetPDVDelta(ctx context.Context, address string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPDVDelta", ctx, address)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPDVDelta indicates an expected call of GetPDVDelta
func (mr *MockIndexStorageMockRecorder) GetPDVDelta(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPDVDelta", reflect.TypeOf((*MockIndexStorage)(nil).GetPDVDelta), ctx, address)
}

// GetPDVTotalDelta mocks base method
func (m *MockIndexStorage) GetPDVTotalDelta(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPDVTotalDelta", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPDVTotalDelta indicates an expected call of GetPDVTotalDelta
func (mr *MockIndexStorageMockRecorder) GetPDVTotalDelta(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPDVTotalDelta", reflect.TypeOf((*MockIndexStorage)(nil).GetPDVTotalDelta), ctx)
}

// GetPDVDeltaList mocks base method
func (m *MockIndexStorage) GetPDVDeltaList(ctx context.Context) ([]*storage.PDVDelta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPDVDeltaList", ctx)
	ret0, _ := ret[0].([]*storage.PDVDelta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPDVDeltaList indicates an expected call of GetPDVDeltaList
func (mr *MockIndexStorageMockRecorder) GetPDVDeltaList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPDVDeltaList", reflect.TypeOf((*MockIndexStorage)(nil).GetPDVDeltaList), ctx)
}

// CreateRewardsQueueItem mocks base method
func (m *MockIndexStorage) CreateRewardsQueueItem(ctx context.Context, addr string, reward int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRewardsQueueItem", ctx, addr, reward)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRewardsQueueItem indicates an expected call of CreateRewardsQueueItem
func (mr *MockIndexStorageMockRecorder) CreateRewardsQueueItem(ctx, addr, reward interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRewardsQueueItem", reflect.TypeOf((*MockIndexStorage)(nil).CreateRewardsQueueItem), ctx, addr, reward)
}

// GetRewardsQueueItemList mocks base method
func (m *MockIndexStorage) GetRewardsQueueItemList(ctx context.Context) ([]*storage.RewardsQueueItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRewardsQueueItemList", ctx)
	ret0, _ := ret[0].([]*storage.RewardsQueueItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRewardsQueueItemList indicates an expected call of GetRewardsQueueItemList
func (mr *MockIndexStorageMockRecorder) GetRewardsQueueItemList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRewardsQueueItemList", reflect.TypeOf((*MockIndexStorage)(nil).GetRewardsQueueItemList), ctx)
}

// DeleteRewardsQueueItem mocks base method
func (m *MockIndexStorage) DeleteRewardsQueueItem(ctx context.Context, addr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRewardsQueueItem", ctx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRewardsQueueItem indicates an expected call of DeleteRewardsQueueItem
func (mr *MockIndexStorageMockRecorder) DeleteRewardsQueueItem(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRewardsQueueItem", reflect.TypeOf((*MockIndexStorage)(nil).DeleteRewardsQueueItem), ctx, addr)
}

// GetPDVRewardsDistributedDate mocks base method
func (m *MockIndexStorage) GetPDVRewardsDistributedDate(ctx context.Context) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPDVRewardsDistributedDate", ctx)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPDVRewardsDistributedDate indicates an expected call of GetPDVRewardsDistributedDate
func (mr *MockIndexStorageMockRecorder) GetPDVRewardsDistributedDate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPDVRewardsDistributedDate", reflect.TypeOf((*MockIndexStorage)(nil).GetPDVRewardsDistributedDate), ctx)
}

// SetPDVRewardsDistributedDate mocks base method
func (m *MockIndexStorage) SetPDVRewardsDistributedDate(ctx context.Context, date time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPDVRewardsDistributedDate", ctx, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPDVRewardsDistributedDate indicates an expected call of SetPDVRewardsDistributedDate
func (mr *MockIndexStorageMockRecorder) SetPDVRewardsDistributedDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPDVRewardsDistributedDate", reflect.TypeOf((*MockIndexStorage)(nil).SetPDVRewardsDistributedDate), ctx, date)
}
