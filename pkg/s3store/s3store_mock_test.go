// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/RobertVillalba/tusd/pkg/s3store (interfaces: S3API)

// Package s3store is a generated GoMock package.
package s3store

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	s3 "github.com/aws/aws-sdk-go/service/s3"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockS3API is a mock of S3API interface
type MockS3API struct {
	ctrl     *gomock.Controller
	recorder *MockS3APIMockRecorder
}

// MockS3APIMockRecorder is the mock recorder for MockS3API
type MockS3APIMockRecorder struct {
	mock *MockS3API
}

// NewMockS3API creates a new mock instance
func NewMockS3API(ctrl *gomock.Controller) *MockS3API {
	mock := &MockS3API{ctrl: ctrl}
	mock.recorder = &MockS3APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockS3API) EXPECT() *MockS3APIMockRecorder {
	return m.recorder
}

// AbortMultipartUploadWithContext mocks base method
func (m *MockS3API) AbortMultipartUploadWithContext(arg0 context.Context, arg1 *s3.AbortMultipartUploadInput, arg2 ...request.Option) (*s3.AbortMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortMultipartUploadWithContext", varargs...)
	ret0, _ := ret[0].(*s3.AbortMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbortMultipartUploadWithContext indicates an expected call of AbortMultipartUploadWithContext
func (mr *MockS3APIMockRecorder) AbortMultipartUploadWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortMultipartUploadWithContext", reflect.TypeOf((*MockS3API)(nil).AbortMultipartUploadWithContext), varargs...)
}

// CompleteMultipartUploadWithContext mocks base method
func (m *MockS3API) CompleteMultipartUploadWithContext(arg0 context.Context, arg1 *s3.CompleteMultipartUploadInput, arg2 ...request.Option) (*s3.CompleteMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteMultipartUploadWithContext", varargs...)
	ret0, _ := ret[0].(*s3.CompleteMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteMultipartUploadWithContext indicates an expected call of CompleteMultipartUploadWithContext
func (mr *MockS3APIMockRecorder) CompleteMultipartUploadWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteMultipartUploadWithContext", reflect.TypeOf((*MockS3API)(nil).CompleteMultipartUploadWithContext), varargs...)
}

// CreateMultipartUploadWithContext mocks base method
func (m *MockS3API) CreateMultipartUploadWithContext(arg0 context.Context, arg1 *s3.CreateMultipartUploadInput, arg2 ...request.Option) (*s3.CreateMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMultipartUploadWithContext", varargs...)
	ret0, _ := ret[0].(*s3.CreateMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMultipartUploadWithContext indicates an expected call of CreateMultipartUploadWithContext
func (mr *MockS3APIMockRecorder) CreateMultipartUploadWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultipartUploadWithContext", reflect.TypeOf((*MockS3API)(nil).CreateMultipartUploadWithContext), varargs...)
}

// DeleteObjectWithContext mocks base method
func (m *MockS3API) DeleteObjectWithContext(arg0 context.Context, arg1 *s3.DeleteObjectInput, arg2 ...request.Option) (*s3.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObjectWithContext", varargs...)
	ret0, _ := ret[0].(*s3.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjectWithContext indicates an expected call of DeleteObjectWithContext
func (mr *MockS3APIMockRecorder) DeleteObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectWithContext", reflect.TypeOf((*MockS3API)(nil).DeleteObjectWithContext), varargs...)
}

// DeleteObjectsWithContext mocks base method
func (m *MockS3API) DeleteObjectsWithContext(arg0 context.Context, arg1 *s3.DeleteObjectsInput, arg2 ...request.Option) (*s3.DeleteObjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObjectsWithContext", varargs...)
	ret0, _ := ret[0].(*s3.DeleteObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjectsWithContext indicates an expected call of DeleteObjectsWithContext
func (mr *MockS3APIMockRecorder) DeleteObjectsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectsWithContext", reflect.TypeOf((*MockS3API)(nil).DeleteObjectsWithContext), varargs...)
}

// GetObjectWithContext mocks base method
func (m *MockS3API) GetObjectWithContext(arg0 context.Context, arg1 *s3.GetObjectInput, arg2 ...request.Option) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectWithContext", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectWithContext indicates an expected call of GetObjectWithContext
func (mr *MockS3APIMockRecorder) GetObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectWithContext", reflect.TypeOf((*MockS3API)(nil).GetObjectWithContext), varargs...)
}

// ListPartsWithContext mocks base method
func (m *MockS3API) ListPartsWithContext(arg0 context.Context, arg1 *s3.ListPartsInput, arg2 ...request.Option) (*s3.ListPartsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPartsWithContext", varargs...)
	ret0, _ := ret[0].(*s3.ListPartsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPartsWithContext indicates an expected call of ListPartsWithContext
func (mr *MockS3APIMockRecorder) ListPartsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPartsWithContext", reflect.TypeOf((*MockS3API)(nil).ListPartsWithContext), varargs...)
}

// PutObjectWithContext mocks base method
func (m *MockS3API) PutObjectWithContext(arg0 context.Context, arg1 *s3.PutObjectInput, arg2 ...request.Option) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObjectWithContext", varargs...)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObjectWithContext indicates an expected call of PutObjectWithContext
func (mr *MockS3APIMockRecorder) PutObjectWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectWithContext", reflect.TypeOf((*MockS3API)(nil).PutObjectWithContext), varargs...)
}

// UploadPartCopyWithContext mocks base method
func (m *MockS3API) UploadPartCopyWithContext(arg0 context.Context, arg1 *s3.UploadPartCopyInput, arg2 ...request.Option) (*s3.UploadPartCopyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadPartCopyWithContext", varargs...)
	ret0, _ := ret[0].(*s3.UploadPartCopyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadPartCopyWithContext indicates an expected call of UploadPartCopyWithContext
func (mr *MockS3APIMockRecorder) UploadPartCopyWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPartCopyWithContext", reflect.TypeOf((*MockS3API)(nil).UploadPartCopyWithContext), varargs...)
}

// UploadPartWithContext mocks base method
func (m *MockS3API) UploadPartWithContext(arg0 context.Context, arg1 *s3.UploadPartInput, arg2 ...request.Option) (*s3.UploadPartOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadPartWithContext", varargs...)
	ret0, _ := ret[0].(*s3.UploadPartOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadPartWithContext indicates an expected call of UploadPartWithContext
func (mr *MockS3APIMockRecorder) UploadPartWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPartWithContext", reflect.TypeOf((*MockS3API)(nil).UploadPartWithContext), varargs...)
}
