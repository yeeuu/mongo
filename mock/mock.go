// Package mock_mongo is a generated GoMock package.
package mock

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/yeeuu/mongo (interfaces: Storage,QuerySet)

import (
	"reflect"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/mock/gomock"
	"github.com/yeeuu/mongo"
)

func TestSetup(t *testing.T) (*MockStorage, *MockQuerySet, func()) {
	ctl := gomock.NewController(t)
	store := NewMockStorage(ctl)
	qs := NewMockQuerySet(ctl)
	return store, qs, func() {
		ctl.Finish()
	}
}

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Query mocks base method
func (m *MockStorage) Query(arg0 mongo.Collection) mongo.QuerySet {
	ret := m.ctrl.Call(m, "Query", arg0)
	ret0, _ := ret[0].(mongo.QuerySet)
	return ret0
}

// Query indicates an expected call of Query
func (mr *MockStorageMockRecorder) Query(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockStorage)(nil).Query), arg0)
}

// Raw mocks base method
func (m *MockStorage) Raw(arg0 mongo.Collection, arg1 bson.M) mongo.QuerySet {
	ret := m.ctrl.Call(m, "Raw", arg0, arg1)
	ret0, _ := ret[0].(mongo.QuerySet)
	return ret0
}

// Raw indicates an expected call of Raw
func (mr *MockStorageMockRecorder) Raw(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockStorage)(nil).Raw), arg0, arg1)
}

// MockQuerySet is a mock of QuerySet interface
type MockQuerySet struct {
	ctrl     *gomock.Controller
	recorder *MockQuerySetMockRecorder
}

// MockQuerySetMockRecorder is the mock recorder for MockQuerySet
type MockQuerySetMockRecorder struct {
	mock *MockQuerySet
}

// NewMockQuerySet creates a new mock instance
func NewMockQuerySet(ctrl *gomock.Controller) *MockQuerySet {
	mock := &MockQuerySet{ctrl: ctrl}
	mock.recorder = &MockQuerySetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuerySet) EXPECT() *MockQuerySetMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockQuerySet) All(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockQuerySetMockRecorder) All(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockQuerySet)(nil).All), arg0)
}

// Count mocks base method
func (m *MockQuerySet) Count() (int, error) {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockQuerySetMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockQuerySet)(nil).Count))
}

// Insert mocks base method
func (m *MockQuerySet) Insert(arg0 ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockQuerySetMockRecorder) Insert(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockQuerySet)(nil).Insert), arg0...)
}

// Limit mocks base method
func (m *MockQuerySet) Limit(arg0 int) mongo.QuerySet {
	ret := m.ctrl.Call(m, "Limit", arg0)
	ret0, _ := ret[0].(mongo.QuerySet)
	return ret0
}

// Limit indicates an expected call of Limit
func (mr *MockQuerySetMockRecorder) Limit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockQuerySet)(nil).Limit), arg0)
}

// One mocks base method
func (m *MockQuerySet) One(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// One indicates an expected call of One
func (mr *MockQuerySetMockRecorder) One(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockQuerySet)(nil).One), arg0)
}

// Prefetch mocks base method
func (m *MockQuerySet) Prefetch(arg0 float64) mongo.QuerySet {
	ret := m.ctrl.Call(m, "Prefetch", arg0)
	ret0, _ := ret[0].(mongo.QuerySet)
	return ret0
}

// Prefetch indicates an expected call of Prefetch
func (mr *MockQuerySetMockRecorder) Prefetch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefetch", reflect.TypeOf((*MockQuerySet)(nil).Prefetch), arg0)
}

// Remove mocks base method
func (m *MockQuerySet) Remove() error {
	ret := m.ctrl.Call(m, "Remove")
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockQuerySetMockRecorder) Remove() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockQuerySet)(nil).Remove))
}

// RemoveAll mocks base method
func (m *MockQuerySet) RemoveAll() (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "RemoveAll")
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *MockQuerySetMockRecorder) RemoveAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockQuerySet)(nil).RemoveAll))
}

// Skip mocks base method
func (m *MockQuerySet) Skip(arg0 int) mongo.QuerySet {
	ret := m.ctrl.Call(m, "Skip", arg0)
	ret0, _ := ret[0].(mongo.QuerySet)
	return ret0
}

// Skip indicates an expected call of Skip
func (mr *MockQuerySetMockRecorder) Skip(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockQuerySet)(nil).Skip), arg0)
}

// Sort mocks base method
func (m *MockQuerySet) Sort(arg0 ...string) mongo.QuerySet {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sort", varargs...)
	ret0, _ := ret[0].(mongo.QuerySet)
	return ret0
}

// Sort indicates an expected call of Sort
func (mr *MockQuerySetMockRecorder) Sort(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockQuerySet)(nil).Sort), arg0...)
}

// Update mocks base method
func (m *MockQuerySet) Update(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockQuerySetMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockQuerySet)(nil).Update), arg0)
}

// UpdateAll mocks base method
func (m *MockQuerySet) UpdateAll(arg0 interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpdateAll", arg0)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll
func (mr *MockQuerySetMockRecorder) UpdateAll(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockQuerySet)(nil).UpdateAll), arg0)
}

// Upsert mocks base method
func (m *MockQuerySet) Upsert(arg0 interface{}) (*mgo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "Upsert", arg0)
	ret0, _ := ret[0].(*mgo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockQuerySetMockRecorder) Upsert(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockQuerySet)(nil).Upsert), arg0)
}
