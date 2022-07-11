// Code generated by MockGen. DO NOT EDIT.
// Source: utils_test.go

// Package handler_test is a generated GoMock package.
package handler_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	handler "github.com/RobertVillalba/tusd/pkg/handler"
	io "io"
	reflect "reflect"
)

// MockFullDataStore is a mock of FullDataStore interface
type MockFullDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockFullDataStoreMockRecorder
}

// MockFullDataStoreMockRecorder is the mock recorder for MockFullDataStore
type MockFullDataStoreMockRecorder struct {
	mock *MockFullDataStore
}

// NewMockFullDataStore creates a new mock instance
func NewMockFullDataStore(ctrl *gomock.Controller) *MockFullDataStore {
	mock := &MockFullDataStore{ctrl: ctrl}
	mock.recorder = &MockFullDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFullDataStore) EXPECT() *MockFullDataStoreMockRecorder {
	return m.recorder
}

// NewUpload mocks base method
func (m *MockFullDataStore) NewUpload(ctx context.Context, info handler.FileInfo) (handler.Upload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUpload", ctx, info)
	ret0, _ := ret[0].(handler.Upload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUpload indicates an expected call of NewUpload
func (mr *MockFullDataStoreMockRecorder) NewUpload(ctx, info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUpload", reflect.TypeOf((*MockFullDataStore)(nil).NewUpload), ctx, info)
}

// GetUpload mocks base method
func (m *MockFullDataStore) GetUpload(ctx context.Context, id string) (handler.Upload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpload", ctx, id)
	ret0, _ := ret[0].(handler.Upload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpload indicates an expected call of GetUpload
func (mr *MockFullDataStoreMockRecorder) GetUpload(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpload", reflect.TypeOf((*MockFullDataStore)(nil).GetUpload), ctx, id)
}

// AsTerminatableUpload mocks base method
func (m *MockFullDataStore) AsTerminatableUpload(upload handler.Upload) handler.TerminatableUpload {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsTerminatableUpload", upload)
	ret0, _ := ret[0].(handler.TerminatableUpload)
	return ret0
}

// AsTerminatableUpload indicates an expected call of AsTerminatableUpload
func (mr *MockFullDataStoreMockRecorder) AsTerminatableUpload(upload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsTerminatableUpload", reflect.TypeOf((*MockFullDataStore)(nil).AsTerminatableUpload), upload)
}

// AsConcatableUpload mocks base method
func (m *MockFullDataStore) AsConcatableUpload(upload handler.Upload) handler.ConcatableUpload {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsConcatableUpload", upload)
	ret0, _ := ret[0].(handler.ConcatableUpload)
	return ret0
}

// AsConcatableUpload indicates an expected call of AsConcatableUpload
func (mr *MockFullDataStoreMockRecorder) AsConcatableUpload(upload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsConcatableUpload", reflect.TypeOf((*MockFullDataStore)(nil).AsConcatableUpload), upload)
}

// AsLengthDeclarableUpload mocks base method
func (m *MockFullDataStore) AsLengthDeclarableUpload(upload handler.Upload) handler.LengthDeclarableUpload {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsLengthDeclarableUpload", upload)
	ret0, _ := ret[0].(handler.LengthDeclarableUpload)
	return ret0
}

// AsLengthDeclarableUpload indicates an expected call of AsLengthDeclarableUpload
func (mr *MockFullDataStoreMockRecorder) AsLengthDeclarableUpload(upload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsLengthDeclarableUpload", reflect.TypeOf((*MockFullDataStore)(nil).AsLengthDeclarableUpload), upload)
}

// MockFullUpload is a mock of FullUpload interface
type MockFullUpload struct {
	ctrl     *gomock.Controller
	recorder *MockFullUploadMockRecorder
}

// MockFullUploadMockRecorder is the mock recorder for MockFullUpload
type MockFullUploadMockRecorder struct {
	mock *MockFullUpload
}

// NewMockFullUpload creates a new mock instance
func NewMockFullUpload(ctrl *gomock.Controller) *MockFullUpload {
	mock := &MockFullUpload{ctrl: ctrl}
	mock.recorder = &MockFullUploadMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFullUpload) EXPECT() *MockFullUploadMockRecorder {
	return m.recorder
}

// WriteChunk mocks base method
func (m *MockFullUpload) WriteChunk(ctx context.Context, offset int64, src io.Reader) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteChunk", ctx, offset, src)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteChunk indicates an expected call of WriteChunk
func (mr *MockFullUploadMockRecorder) WriteChunk(ctx, offset, src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteChunk", reflect.TypeOf((*MockFullUpload)(nil).WriteChunk), ctx, offset, src)
}

// GetInfo mocks base method
func (m *MockFullUpload) GetInfo(ctx context.Context) (handler.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo", ctx)
	ret0, _ := ret[0].(handler.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo
func (mr *MockFullUploadMockRecorder) GetInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockFullUpload)(nil).GetInfo), ctx)
}

// GetReader mocks base method
func (m *MockFullUpload) GetReader(ctx context.Context) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReader", ctx)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReader indicates an expected call of GetReader
func (mr *MockFullUploadMockRecorder) GetReader(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReader", reflect.TypeOf((*MockFullUpload)(nil).GetReader), ctx)
}

// FinishUpload mocks base method
func (m *MockFullUpload) FinishUpload(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishUpload", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinishUpload indicates an expected call of FinishUpload
func (mr *MockFullUploadMockRecorder) FinishUpload(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishUpload", reflect.TypeOf((*MockFullUpload)(nil).FinishUpload), ctx)
}

// Terminate mocks base method
func (m *MockFullUpload) Terminate(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Terminate", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Terminate indicates an expected call of Terminate
func (mr *MockFullUploadMockRecorder) Terminate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terminate", reflect.TypeOf((*MockFullUpload)(nil).Terminate), ctx)
}

// DeclareLength mocks base method
func (m *MockFullUpload) DeclareLength(ctx context.Context, length int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclareLength", ctx, length)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeclareLength indicates an expected call of DeclareLength
func (mr *MockFullUploadMockRecorder) DeclareLength(ctx, length interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclareLength", reflect.TypeOf((*MockFullUpload)(nil).DeclareLength), ctx, length)
}

// ConcatUploads mocks base method
func (m *MockFullUpload) ConcatUploads(ctx context.Context, partialUploads []handler.Upload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConcatUploads", ctx, partialUploads)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConcatUploads indicates an expected call of ConcatUploads
func (mr *MockFullUploadMockRecorder) ConcatUploads(ctx, partialUploads interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConcatUploads", reflect.TypeOf((*MockFullUpload)(nil).ConcatUploads), ctx, partialUploads)
}

// MockFullLocker is a mock of FullLocker interface
type MockFullLocker struct {
	ctrl     *gomock.Controller
	recorder *MockFullLockerMockRecorder
}

// MockFullLockerMockRecorder is the mock recorder for MockFullLocker
type MockFullLockerMockRecorder struct {
	mock *MockFullLocker
}

// NewMockFullLocker creates a new mock instance
func NewMockFullLocker(ctrl *gomock.Controller) *MockFullLocker {
	mock := &MockFullLocker{ctrl: ctrl}
	mock.recorder = &MockFullLockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFullLocker) EXPECT() *MockFullLockerMockRecorder {
	return m.recorder
}

// NewLock mocks base method
func (m *MockFullLocker) NewLock(id string) (handler.Lock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewLock", id)
	ret0, _ := ret[0].(handler.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewLock indicates an expected call of NewLock
func (mr *MockFullLockerMockRecorder) NewLock(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewLock", reflect.TypeOf((*MockFullLocker)(nil).NewLock), id)
}

// MockFullLock is a mock of FullLock interface
type MockFullLock struct {
	ctrl     *gomock.Controller
	recorder *MockFullLockMockRecorder
}

// MockFullLockMockRecorder is the mock recorder for MockFullLock
type MockFullLockMockRecorder struct {
	mock *MockFullLock
}

// NewMockFullLock creates a new mock instance
func NewMockFullLock(ctrl *gomock.Controller) *MockFullLock {
	mock := &MockFullLock{ctrl: ctrl}
	mock.recorder = &MockFullLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFullLock) EXPECT() *MockFullLockMockRecorder {
	return m.recorder
}

// Lock mocks base method
func (m *MockFullLock) Lock() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock")
	ret0, _ := ret[0].(error)
	return ret0
}

// Lock indicates an expected call of Lock
func (mr *MockFullLockMockRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockFullLock)(nil).Lock))
}

// Unlock mocks base method
func (m *MockFullLock) Unlock() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock
func (mr *MockFullLockMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockFullLock)(nil).Unlock))
}
