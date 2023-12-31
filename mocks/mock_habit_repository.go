// Code generated by MockGen. DO NOT EDIT.
// Source: go-api-backend/internal/repository (interfaces: HabitRepositoryInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "go-api-backend/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHabitRepositoryInterface is a mock of HabitRepositoryInterface interface.
type MockHabitRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHabitRepositoryInterfaceMockRecorder
}

// MockHabitRepositoryInterfaceMockRecorder is the mock recorder for MockHabitRepositoryInterface.
type MockHabitRepositoryInterfaceMockRecorder struct {
	mock *MockHabitRepositoryInterface
}

// NewMockHabitRepositoryInterface creates a new mock instance.
func NewMockHabitRepositoryInterface(ctrl *gomock.Controller) *MockHabitRepositoryInterface {
	mock := &MockHabitRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockHabitRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHabitRepositoryInterface) EXPECT() *MockHabitRepositoryInterfaceMockRecorder {
	return m.recorder
}

// InsertHabit mocks base method.
func (m *MockHabitRepositoryInterface) InsertHabit(arg0 models.Habit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHabit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertHabit indicates an expected call of InsertHabit.
func (mr *MockHabitRepositoryInterfaceMockRecorder) InsertHabit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHabit", reflect.TypeOf((*MockHabitRepositoryInterface)(nil).InsertHabit), arg0)
}
