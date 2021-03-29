// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mic90/metrics-api/persistance (interfaces: Storage)

// Package mock_persistance is a generated GoMock package.
package mock_persistance

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	metrics "github.com/mic90/metrics-api/metrics"
	data "github.com/mic90/metrics-api/metrics/data"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddData mocks base method.
func (m *MockStorage) AddData(arg0 metrics.Descriptor, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddData", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddData indicates an expected call of AddData.
func (mr *MockStorageMockRecorder) AddData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddData", reflect.TypeOf((*MockStorage)(nil).AddData), arg0, arg1)
}

// AddMetric mocks base method.
func (m *MockStorage) AddMetric(arg0 metrics.Descriptor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMetric", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMetric indicates an expected call of AddMetric.
func (mr *MockStorageMockRecorder) AddMetric(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMetric", reflect.TypeOf((*MockStorage)(nil).AddMetric), arg0)
}

// GetData mocks base method.
func (m *MockStorage) GetData(arg0 metrics.Descriptor, arg1, arg2 time.Time) ([]data.Point, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", arg0, arg1, arg2)
	ret0, _ := ret[0].([]data.Point)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockStorageMockRecorder) GetData(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockStorage)(nil).GetData), arg0, arg1, arg2)
}

// GetMetrics mocks base method.
func (m *MockStorage) GetMetrics() []metrics.Descriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetrics")
	ret0, _ := ret[0].([]metrics.Descriptor)
	return ret0
}

// GetMetrics indicates an expected call of GetMetrics.
func (mr *MockStorageMockRecorder) GetMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetrics", reflect.TypeOf((*MockStorage)(nil).GetMetrics))
}

// RemoveMetric mocks base method.
func (m *MockStorage) RemoveMetric(arg0 metrics.Descriptor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMetric", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMetric indicates an expected call of RemoveMetric.
func (mr *MockStorageMockRecorder) RemoveMetric(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMetric", reflect.TypeOf((*MockStorage)(nil).RemoveMetric), arg0)
}
