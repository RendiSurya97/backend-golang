// Code generated by MockGen. DO NOT EDIT.
// Source: internal/movie/ext_module/omdb.go

// Package mock_ext_module is a generated GoMock package.
package mock_ext_module

import (
	ext_module "github.com/backendgolang/searchapp/internal/movie/ext_module"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIOmdb is a mock of IOmdb interface
type MockIOmdb struct {
	ctrl     *gomock.Controller
	recorder *MockIOmdbMockRecorder
}

// MockIOmdbMockRecorder is the mock recorder for MockIOmdb
type MockIOmdbMockRecorder struct {
	mock *MockIOmdb
}

// NewMockIOmdb creates a new mock instance
func NewMockIOmdb(ctrl *gomock.Controller) *MockIOmdb {
	mock := &MockIOmdb{ctrl: ctrl}
	mock.recorder = &MockIOmdbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIOmdb) EXPECT() *MockIOmdbMockRecorder {
	return m.recorder
}

// GetMovieByTitle mocks base method
func (m *MockIOmdb) GetMovieByTitle(title string, page int64) (ext_module.MovieResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieByTitle", title, page)
	ret0, _ := ret[0].(ext_module.MovieResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieByTitle indicates an expected call of GetMovieByTitle
func (mr *MockIOmdbMockRecorder) GetMovieByTitle(title, page interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieByTitle", reflect.TypeOf((*MockIOmdb)(nil).GetMovieByTitle), title, page)
}
