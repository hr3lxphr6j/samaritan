// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/samaritan-proxy/samaritan/proc/redis/compressor (interfaces: Compressor)

// Package compressor is a generated GoMock package.
package compressor

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockCompressor is a mock of Compressor interface
type MockCompressor struct {
	ctrl     *gomock.Controller
	recorder *MockCompressorMockRecorder
}

// MockCompressorMockRecorder is the mock recorder for MockCompressor
type MockCompressorMockRecorder struct {
	mock *MockCompressor
}

// NewMockCompressor creates a new mock instance
func NewMockCompressor(ctrl *gomock.Controller) *MockCompressor {
	mock := &MockCompressor{ctrl: ctrl}
	mock.recorder = &MockCompressorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCompressor) EXPECT() *MockCompressorMockRecorder {
	return m.recorder
}

// Compress mocks base method
func (m *MockCompressor) Compress(arg0 io.Writer) io.WriteCloser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compress", arg0)
	ret0, _ := ret[0].(io.WriteCloser)
	return ret0
}

// Compress indicates an expected call of Compress
func (mr *MockCompressorMockRecorder) Compress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compress", reflect.TypeOf((*MockCompressor)(nil).Compress), arg0)
}

// Decompress mocks base method
func (m *MockCompressor) Decompress(arg0 io.Reader) io.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decompress", arg0)
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

// Decompress indicates an expected call of Decompress
func (mr *MockCompressorMockRecorder) Decompress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decompress", reflect.TypeOf((*MockCompressor)(nil).Decompress), arg0)
}
