// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package repomock is a generated GoMock package.
package repomock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	src "github.com/idzharbae/quickbid/src"
	db "github.com/idzharbae/quickbid/src/bridge/db"
	entity "github.com/idzharbae/quickbid/src/entity"
	reflect "reflect"
)

// MockAttendanceReaderRepo is a mock of AttendanceReaderRepo interface
type MockAttendanceReaderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAttendanceReaderRepoMockRecorder
}

// MockAttendanceReaderRepoMockRecorder is the mock recorder for MockAttendanceReaderRepo
type MockAttendanceReaderRepoMockRecorder struct {
	mock *MockAttendanceReaderRepo
}

// NewMockAttendanceReaderRepo creates a new mock instance
func NewMockAttendanceReaderRepo(ctrl *gomock.Controller) *MockAttendanceReaderRepo {
	mock := &MockAttendanceReaderRepo{ctrl: ctrl}
	mock.recorder = &MockAttendanceReaderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAttendanceReaderRepo) EXPECT() *MockAttendanceReaderRepoMockRecorder {
	return m.recorder
}

// GetByName mocks base method
func (m *MockAttendanceReaderRepo) GetByName(ctx context.Context, name string) (entity.Attendance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", ctx, name)
	ret0, _ := ret[0].(entity.Attendance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockAttendanceReaderRepoMockRecorder) GetByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockAttendanceReaderRepo)(nil).GetByName), ctx, name)
}

// WithTx mocks base method
func (m *MockAttendanceReaderRepo) WithTx(arg0 db.Tx) src.AttendanceReaderRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", arg0)
	ret0, _ := ret[0].(src.AttendanceReaderRepo)
	return ret0
}

// WithTx indicates an expected call of WithTx
func (mr *MockAttendanceReaderRepoMockRecorder) WithTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockAttendanceReaderRepo)(nil).WithTx), arg0)
}

// MockAttendanceWriterRepo is a mock of AttendanceWriterRepo interface
type MockAttendanceWriterRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAttendanceWriterRepoMockRecorder
}

// MockAttendanceWriterRepoMockRecorder is the mock recorder for MockAttendanceWriterRepo
type MockAttendanceWriterRepoMockRecorder struct {
	mock *MockAttendanceWriterRepo
}

// NewMockAttendanceWriterRepo creates a new mock instance
func NewMockAttendanceWriterRepo(ctrl *gomock.Controller) *MockAttendanceWriterRepo {
	mock := &MockAttendanceWriterRepo{ctrl: ctrl}
	mock.recorder = &MockAttendanceWriterRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAttendanceWriterRepo) EXPECT() *MockAttendanceWriterRepoMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockAttendanceWriterRepo) Insert(ctx context.Context, attendance entity.Attendance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, attendance)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockAttendanceWriterRepoMockRecorder) Insert(ctx, attendance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAttendanceWriterRepo)(nil).Insert), ctx, attendance)
}

// WithTx mocks base method
func (m *MockAttendanceWriterRepo) WithTx(arg0 db.Tx) src.AttendanceWriterRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", arg0)
	ret0, _ := ret[0].(src.AttendanceWriterRepo)
	return ret0
}

// WithTx indicates an expected call of WithTx
func (mr *MockAttendanceWriterRepoMockRecorder) WithTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockAttendanceWriterRepo)(nil).WithTx), arg0)
}

// MockBidReaderRepo is a mock of BidReaderRepo interface
type MockBidReaderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockBidReaderRepoMockRecorder
}

// MockBidReaderRepoMockRecorder is the mock recorder for MockBidReaderRepo
type MockBidReaderRepoMockRecorder struct {
	mock *MockBidReaderRepo
}

// NewMockBidReaderRepo creates a new mock instance
func NewMockBidReaderRepo(ctrl *gomock.Controller) *MockBidReaderRepo {
	mock := &MockBidReaderRepo{ctrl: ctrl}
	mock.recorder = &MockBidReaderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBidReaderRepo) EXPECT() *MockBidReaderRepoMockRecorder {
	return m.recorder
}

// ListUserBiddedProducts mocks base method
func (m *MockBidReaderRepo) ListUserBiddedProducts(ctx context.Context, userID, page, limit int) ([]entity.Bid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserBiddedProducts", ctx, userID, page, limit)
	ret0, _ := ret[0].([]entity.Bid)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserBiddedProducts indicates an expected call of ListUserBiddedProducts
func (mr *MockBidReaderRepoMockRecorder) ListUserBiddedProducts(ctx, userID, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserBiddedProducts", reflect.TypeOf((*MockBidReaderRepo)(nil).ListUserBiddedProducts), ctx, userID, page, limit)
}

// WithTx mocks base method
func (m *MockBidReaderRepo) WithTx(arg0 db.Tx) src.BidReaderRepo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", arg0)
	ret0, _ := ret[0].(src.BidReaderRepo)
	return ret0
}

// WithTx indicates an expected call of WithTx
func (mr *MockBidReaderRepoMockRecorder) WithTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockBidReaderRepo)(nil).WithTx), arg0)
}
