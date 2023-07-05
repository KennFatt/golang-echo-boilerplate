// Code generated by mockery v2.30.1. DO NOT EDIT.

package mock_http

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// QuotesService is an autogenerated mock type for the QuotesService type
type QuotesService struct {
	mock.Mock
}

// GetQuotes provides a mock function with given fields: ctx
func (_m *QuotesService) GetQuotes(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewQuotesService creates a new instance of QuotesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuotesService(t interface {
	mock.TestingT
	Cleanup(func())
}) *QuotesService {
	mock := &QuotesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
