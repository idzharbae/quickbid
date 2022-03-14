// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package ucmock is a generated GoMock package.
package ucmock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAttendanceUC is a mock of AttendanceUC interface
type MockAttendanceUC struct {
	ctrl     *gomock.Controller
	recorder *MockAttendanceUCMockRecorder
}

// MockAttendanceUCMockRecorder is the mock recorder for MockAttendanceUC
type MockAttendanceUCMockRecorder struct {
	mock *MockAttendanceUC
}

// NewMockAttendanceUC creates a new mock instance
func NewMockAttendanceUC(ctrl *gomock.Controller) *MockAttendanceUC {
	mock := &MockAttendanceUC{ctrl: ctrl}
	mock.recorder = &MockAttendanceUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAttendanceUC) EXPECT() *MockAttendanceUCMockRecorder {
	return m.recorder
}

// Attend mocks base method
func (m *MockAttendanceUC) Attend(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attend", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Attend indicates an expected call of Attend
func (mr *MockAttendanceUCMockRecorder) Attend(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attend", reflect.TypeOf((*MockAttendanceUC)(nil).Attend), ctx, name)
}
