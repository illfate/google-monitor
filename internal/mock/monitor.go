// Code generated by MockGen. DO NOT EDIT.
// Source: internal/monitor/monitor.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	monitor "github.com/illfate/google-monitor/internal/monitor"
)

// MockGoogleClient is a mock of GoogleClient interface.
type MockGoogleClient struct {
	ctrl     *gomock.Controller
	recorder *MockGoogleClientMockRecorder
}

// MockGoogleClientMockRecorder is the mock recorder for MockGoogleClient.
type MockGoogleClientMockRecorder struct {
	mock *MockGoogleClient
}

// NewMockGoogleClient creates a new mock instance.
func NewMockGoogleClient(ctrl *gomock.Controller) *MockGoogleClient {
	mock := &MockGoogleClient{ctrl: ctrl}
	mock.recorder = &MockGoogleClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGoogleClient) EXPECT() *MockGoogleClientMockRecorder {
	return m.recorder
}

// MakeGetRequest mocks base method.
func (m *MockGoogleClient) MakeGetRequest(ctx context.Context) (monitor.RequestResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeGetRequest", ctx)
	ret0, _ := ret[0].(monitor.RequestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeGetRequest indicates an expected call of MakeGetRequest.
func (mr *MockGoogleClientMockRecorder) MakeGetRequest(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeGetRequest", reflect.TypeOf((*MockGoogleClient)(nil).MakeGetRequest), ctx)
}

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// InsertRequestRes mocks base method.
func (m *MockRepository) InsertRequestRes(ctx context.Context, res monitor.RequestResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRequestRes", ctx, res)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertRequestRes indicates an expected call of InsertRequestRes.
func (mr *MockRepositoryMockRecorder) InsertRequestRes(ctx, res interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRequestRes", reflect.TypeOf((*MockRepository)(nil).InsertRequestRes), ctx, res)
}
