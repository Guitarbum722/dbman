// Code generated by MockGen. DO NOT EDIT.
// Source: state.go

// Package main is a generated GoMock package.
package main

import (
	dbman "dabbertorres.dev/dbman"
	gomock "github.com/golang/mock/gomock"
	ssh "golang.org/x/crypto/ssh"
	reflect "reflect"
)

// MockdbManager is a mock of dbManager interface
type MockdbManager struct {
	ctrl     *gomock.Controller
	recorder *MockdbManagerMockRecorder
}

// MockdbManagerMockRecorder is the mock recorder for MockdbManager
type MockdbManagerMockRecorder struct {
	mock *MockdbManager
}

// NewMockdbManager creates a new mock instance
func NewMockdbManager(ctrl *gomock.Controller) *MockdbManager {
	mock := &MockdbManager{ctrl: ctrl}
	mock.recorder = &MockdbManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockdbManager) EXPECT() *MockdbManagerMockRecorder {
	return m.recorder
}

// CurrentName mocks base method
func (m *MockdbManager) CurrentName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentName")
	ret0, _ := ret[0].(string)
	return ret0
}

// CurrentName indicates an expected call of CurrentName
func (mr *MockdbManagerMockRecorder) CurrentName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentName", reflect.TypeOf((*MockdbManager)(nil).CurrentName))
}

// ListConnections mocks base method
func (m *MockdbManager) ListConnections() ([]string, []bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConnections")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].([]bool)
	return ret0, ret1
}

// ListConnections indicates an expected call of ListConnections
func (mr *MockdbManagerMockRecorder) ListConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnections", reflect.TypeOf((*MockdbManager)(nil).ListConnections))
}

// SwitchConnection mocks base method
func (m *MockdbManager) SwitchConnection(connName string, prompter ssh.KeyboardInteractiveChallenge) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwitchConnection", connName, prompter)
	ret0, _ := ret[0].(error)
	return ret0
}

// SwitchConnection indicates an expected call of SwitchConnection
func (mr *MockdbManagerMockRecorder) SwitchConnection(connName, prompter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwitchConnection", reflect.TypeOf((*MockdbManager)(nil).SwitchConnection), connName, prompter)
}

// ListTables mocks base method
func (m *MockdbManager) ListTables(schema string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTables", schema)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTables indicates an expected call of ListTables
func (mr *MockdbManagerMockRecorder) ListTables(schema interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTables", reflect.TypeOf((*MockdbManager)(nil).ListTables), schema)
}

// ListSchemas mocks base method
func (m *MockdbManager) ListSchemas() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchemas")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchemas indicates an expected call of ListSchemas
func (mr *MockdbManagerMockRecorder) ListSchemas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchemas", reflect.TypeOf((*MockdbManager)(nil).ListSchemas))
}

// DescribeTable mocks base method
func (m *MockdbManager) DescribeTable(name string) (*dbman.TableSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTable", name)
	ret0, _ := ret[0].(*dbman.TableSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTable indicates an expected call of DescribeTable
func (mr *MockdbManagerMockRecorder) DescribeTable(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTable", reflect.TypeOf((*MockdbManager)(nil).DescribeTable), name)
}

// Query mocks base method
func (m *MockdbManager) Query(script string) (*dbman.QueryResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", script)
	ret0, _ := ret[0].(*dbman.QueryResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockdbManagerMockRecorder) Query(script interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockdbManager)(nil).Query), script)
}
