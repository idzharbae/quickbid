// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package ucmock is a generated GoMock package.
package ucmock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	entity "github.com/idzharbae/quickbid/src/entity"
	requests "github.com/idzharbae/quickbid/src/requests"
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

// AttendBulk mocks base method
func (m *MockAttendanceUC) AttendBulk(ctx context.Context, names []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttendBulk", ctx, names)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttendBulk indicates an expected call of AttendBulk
func (mr *MockAttendanceUCMockRecorder) AttendBulk(ctx, names interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttendBulk", reflect.TypeOf((*MockAttendanceUC)(nil).AttendBulk), ctx, names)
}

// MockBidUC is a mock of BidUC interface
type MockBidUC struct {
	ctrl     *gomock.Controller
	recorder *MockBidUCMockRecorder
}

// MockBidUCMockRecorder is the mock recorder for MockBidUC
type MockBidUCMockRecorder struct {
	mock *MockBidUC
}

// NewMockBidUC creates a new mock instance
func NewMockBidUC(ctrl *gomock.Controller) *MockBidUC {
	mock := &MockBidUC{ctrl: ctrl}
	mock.recorder = &MockBidUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBidUC) EXPECT() *MockBidUCMockRecorder {
	return m.recorder
}

// ListUserBiddedProducts mocks base method
func (m *MockBidUC) ListUserBiddedProducts(ctx context.Context, req requests.ListUserBiddedProductsRequest) ([]entity.Bid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserBiddedProducts", ctx, req)
	ret0, _ := ret[0].([]entity.Bid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserBiddedProducts indicates an expected call of ListUserBiddedProducts
func (mr *MockBidUCMockRecorder) ListUserBiddedProducts(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserBiddedProducts", reflect.TypeOf((*MockBidUC)(nil).ListUserBiddedProducts), ctx, req)
}
