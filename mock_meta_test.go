// Code generated by MockGen. DO NOT EDIT.
// Source: meta.go

// Package dbman is a generated GoMock package.
package dbman

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockquerier is a mock of querier interface
type Mockquerier struct {
	ctrl     *gomock.Controller
	recorder *MockquerierMockRecorder
}

// MockquerierMockRecorder is the mock recorder for Mockquerier
type MockquerierMockRecorder struct {
	mock *Mockquerier
}

// NewMockquerier creates a new mock instance
func NewMockquerier(ctrl *gomock.Controller) *Mockquerier {
	mock := &Mockquerier{ctrl: ctrl}
	mock.recorder = &MockquerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockquerier) EXPECT() *MockquerierMockRecorder {
	return m.recorder
}

// PingContext mocks base method
func (m *Mockquerier) PingContext(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingContext", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingContext indicates an expected call of PingContext
func (mr *MockquerierMockRecorder) PingContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingContext", reflect.TypeOf((*Mockquerier)(nil).PingContext), arg0)
}

// Query mocks base method
func (m *Mockquerier) Query(arg0 string, arg1 ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockquerierMockRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*Mockquerier)(nil).Query), varargs...)
}

// Stats mocks base method
func (m *Mockquerier) Stats() sql.DBStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(sql.DBStats)
	return ret0
}

// Stats indicates an expected call of Stats
func (mr *MockquerierMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*Mockquerier)(nil).Stats))
}

// Close mocks base method
func (m *Mockquerier) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockquerierMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*Mockquerier)(nil).Close))
}

// MockmetaQuerier is a mock of metaQuerier interface
type MockmetaQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockmetaQuerierMockRecorder
}

// MockmetaQuerierMockRecorder is the mock recorder for MockmetaQuerier
type MockmetaQuerierMockRecorder struct {
	mock *MockmetaQuerier
}

// NewMockmetaQuerier creates a new mock instance
func NewMockmetaQuerier(ctrl *gomock.Controller) *MockmetaQuerier {
	mock := &MockmetaQuerier{ctrl: ctrl}
	mock.recorder = &MockmetaQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmetaQuerier) EXPECT() *MockmetaQuerierMockRecorder {
	return m.recorder
}

// PingContext mocks base method
func (m *MockmetaQuerier) PingContext(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingContext", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingContext indicates an expected call of PingContext
func (mr *MockmetaQuerierMockRecorder) PingContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingContext", reflect.TypeOf((*MockmetaQuerier)(nil).PingContext), arg0)
}

// Query mocks base method
func (m *MockmetaQuerier) Query(arg0 string, arg1 ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (mr *MockmetaQuerierMockRecorder) Query(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockmetaQuerier)(nil).Query), varargs...)
}

// Stats mocks base method
func (m *MockmetaQuerier) Stats() sql.DBStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(sql.DBStats)
	return ret0
}

// Stats indicates an expected call of Stats
func (mr *MockmetaQuerierMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockmetaQuerier)(nil).Stats))
}

// Close mocks base method
func (m *MockmetaQuerier) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockmetaQuerierMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockmetaQuerier)(nil).Close))
}

// ListTables mocks base method
func (m *MockmetaQuerier) ListTables() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTables")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTables indicates an expected call of ListTables
func (mr *MockmetaQuerierMockRecorder) ListTables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTables", reflect.TypeOf((*MockmetaQuerier)(nil).ListTables))
}

// ListTablesInSchema mocks base method
func (m *MockmetaQuerier) ListTablesInSchema(arg0 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTablesInSchema", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTablesInSchema indicates an expected call of ListTablesInSchema
func (mr *MockmetaQuerierMockRecorder) ListTablesInSchema(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTablesInSchema", reflect.TypeOf((*MockmetaQuerier)(nil).ListTablesInSchema), arg0)
}

// ListSchemas mocks base method
func (m *MockmetaQuerier) ListSchemas() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchemas")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchemas indicates an expected call of ListSchemas
func (mr *MockmetaQuerierMockRecorder) ListSchemas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchemas", reflect.TypeOf((*MockmetaQuerier)(nil).ListSchemas))
}

// DescribeTable mocks base method
func (m *MockmetaQuerier) DescribeTable(arg0 string) (*TableSchema, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTable", arg0)
	ret0, _ := ret[0].(*TableSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTable indicates an expected call of DescribeTable
func (mr *MockmetaQuerierMockRecorder) DescribeTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTable", reflect.TypeOf((*MockmetaQuerier)(nil).DescribeTable), arg0)
}
