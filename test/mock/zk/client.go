// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/userpro/gohbase/zk (interfaces: Client)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	zk "github.com/userpro/gohbase/zk"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// LocateResource mocks base method
func (m *MockClient) LocateResource(arg0 zk.ResourceName) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocateResource", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocateResource indicates an expected call of LocateResource
func (mr *MockClientMockRecorder) LocateResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocateResource", reflect.TypeOf((*MockClient)(nil).LocateResource), arg0)
}
