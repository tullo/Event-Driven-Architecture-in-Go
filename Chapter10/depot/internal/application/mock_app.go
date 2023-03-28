// Code generated by mockery v2.23.1. DO NOT EDIT.

package application

import (
	context "context"
	commands "eda-in-golang/depot/internal/application/commands"

	domain "eda-in-golang/depot/internal/domain"

	mock "github.com/stretchr/testify/mock"

	queries "eda-in-golang/depot/internal/application/queries"
)

// MockApp is an autogenerated mock type for the App type
type MockApp struct {
	mock.Mock
}

// AssignShoppingList provides a mock function with given fields: ctx, cmd
func (_m *MockApp) AssignShoppingList(ctx context.Context, cmd commands.AssignShoppingList) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.AssignShoppingList) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CancelShoppingList provides a mock function with given fields: ctx, cmd
func (_m *MockApp) CancelShoppingList(ctx context.Context, cmd commands.CancelShoppingList) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CancelShoppingList) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteShoppingList provides a mock function with given fields: ctx, cmd
func (_m *MockApp) CompleteShoppingList(ctx context.Context, cmd commands.CompleteShoppingList) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CompleteShoppingList) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateShoppingList provides a mock function with given fields: ctx, cmd
func (_m *MockApp) CreateShoppingList(ctx context.Context, cmd commands.CreateShoppingList) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.CreateShoppingList) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetShoppingList provides a mock function with given fields: ctx, query
func (_m *MockApp) GetShoppingList(ctx context.Context, query queries.GetShoppingList) (*domain.ShoppingList, error) {
	ret := _m.Called(ctx, query)

	var r0 *domain.ShoppingList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, queries.GetShoppingList) (*domain.ShoppingList, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, queries.GetShoppingList) *domain.ShoppingList); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ShoppingList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, queries.GetShoppingList) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitiateShopping provides a mock function with given fields: ctx, cmd
func (_m *MockApp) InitiateShopping(ctx context.Context, cmd commands.InitiateShopping) error {
	ret := _m.Called(ctx, cmd)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, commands.InitiateShopping) error); ok {
		r0 = rf(ctx, cmd)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockApp interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockApp creates a new instance of MockApp. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockApp(t mockConstructorTestingTNewMockApp) *MockApp {
	mock := &MockApp{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
