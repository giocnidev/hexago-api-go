// Code generated by MockGen. DO NOT EDIT.
// Source: internal/port/product_usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/giocnidev/api-hexagonal/internal/core/model"
	gomock "go.uber.org/mock/gomock"
)

// MockCreateProductUseCase is a mock of CreateProductUseCase interface.
type MockCreateProductUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCreateProductUseCaseMockRecorder
}

// MockCreateProductUseCaseMockRecorder is the mock recorder for MockCreateProductUseCase.
type MockCreateProductUseCaseMockRecorder struct {
	mock *MockCreateProductUseCase
}

// NewMockCreateProductUseCase creates a new mock instance.
func NewMockCreateProductUseCase(ctrl *gomock.Controller) *MockCreateProductUseCase {
	mock := &MockCreateProductUseCase{ctrl: ctrl}
	mock.recorder = &MockCreateProductUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateProductUseCase) EXPECT() *MockCreateProductUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockCreateProductUseCase) Execute(arg0 model.Product) (*model.Product, *model.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(*model.Error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockCreateProductUseCaseMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockCreateProductUseCase)(nil).Execute), arg0)
}

// MockGetProductById is a mock of GetProductById interface.
type MockGetProductById struct {
	ctrl     *gomock.Controller
	recorder *MockGetProductByIdMockRecorder
}

// MockGetProductByIdMockRecorder is the mock recorder for MockGetProductById.
type MockGetProductByIdMockRecorder struct {
	mock *MockGetProductById
}

// NewMockGetProductById creates a new mock instance.
func NewMockGetProductById(ctrl *gomock.Controller) *MockGetProductById {
	mock := &MockGetProductById{ctrl: ctrl}
	mock.recorder = &MockGetProductByIdMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetProductById) EXPECT() *MockGetProductByIdMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockGetProductById) Execute(arg0 string) (*model.Product, *model.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(*model.Product)
	ret1, _ := ret[1].(*model.Error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockGetProductByIdMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockGetProductById)(nil).Execute), arg0)
}