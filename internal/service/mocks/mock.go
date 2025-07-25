// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"
	models "vk-server-task/internal/models"
	service "vk-server-task/internal/service"

	gomock "github.com/golang/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuth) CreateUser(ctx context.Context, login, password string) (*models.User, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, login, password)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthMockRecorder) CreateUser(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuth)(nil).CreateUser), ctx, login, password)
}

// LoginUser mocks base method.
func (m *MockAuth) LoginUser(ctx context.Context, login, password string) (*models.User, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", ctx, login, password)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockAuthMockRecorder) LoginUser(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockAuth)(nil).LoginUser), ctx, login, password)
}

// ParseToken mocks base method.
func (m *MockAuth) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuth)(nil).ParseToken), accessToken)
}

// MockAds is a mock of Ads interface.
type MockAds struct {
	ctrl     *gomock.Controller
	recorder *MockAdsMockRecorder
}

// MockAdsMockRecorder is the mock recorder for MockAds.
type MockAdsMockRecorder struct {
	mock *MockAds
}

// NewMockAds creates a new mock instance.
func NewMockAds(ctrl *gomock.Controller) *MockAds {
	mock := &MockAds{ctrl: ctrl}
	mock.recorder = &MockAdsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAds) EXPECT() *MockAdsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAds) Create(ctx context.Context, userId int, params service.CreateRequest) (*models.Ads, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userId, params)
	ret0, _ := ret[0].(*models.Ads)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAdsMockRecorder) Create(ctx, userId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAds)(nil).Create), ctx, userId, params)
}

// Get mocks base method.
func (m *MockAds) Get(ctx context.Context, userId int, params models.AdsParams) ([]models.AdsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userId, params)
	ret0, _ := ret[0].([]models.AdsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAdsMockRecorder) Get(ctx, userId, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAds)(nil).Get), ctx, userId, params)
}
