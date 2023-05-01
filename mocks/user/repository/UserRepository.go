// Code generated by mockery v2.20.0. DO NOT EDIT.

package mock_repository

import (
	context "context"
	ent "myapp/ent"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, txClient, newUser
func (_m *UserRepository) Create(ctx context.Context, txClient *ent.Client, newUser ent.User) (*ent.User, error) {
	ret := _m.Called(ctx, txClient, newUser)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, ent.User) (*ent.User, error)); ok {
		return rf(ctx, txClient, newUser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, ent.User) *ent.User); ok {
		r0 = rf(ctx, txClient, newUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Client, ent.User) error); ok {
		r1 = rf(ctx, txClient, newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, tx, id
func (_m *UserRepository) Delete(ctx context.Context, tx *ent.Tx, id uint64) (*ent.User, error) {
	ret := _m.Called(ctx, tx, id)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx, uint64) (*ent.User, error)); ok {
		return rf(ctx, tx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Tx, uint64) *ent.User); ok {
		r0 = rf(ctx, tx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Tx, uint64) error); ok {
		r1 = rf(ctx, tx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *UserRepository) GetAll(ctx context.Context) ([]*ent.User, error) {
	ret := _m.Called(ctx)

	var r0 []*ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*ent.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*ent.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *UserRepository) GetById(ctx context.Context, id uint64) (*ent.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*ent.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *ent.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SoftDelete provides a mock function with given fields: ctx, id
func (_m *UserRepository) SoftDelete(ctx context.Context, id uint64) (*ent.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*ent.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *ent.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, txClient, updateUser, id
func (_m *UserRepository) Update(ctx context.Context, txClient *ent.Client, updateUser ent.User, id uint64) (*ent.User, error) {
	ret := _m.Called(ctx, txClient, updateUser, id)

	var r0 *ent.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, ent.User, uint64) (*ent.User, error)); ok {
		return rf(ctx, txClient, updateUser, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ent.Client, ent.User, uint64) *ent.User); ok {
		r0 = rf(ctx, txClient, updateUser, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ent.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ent.Client, ent.User, uint64) error); ok {
		r1 = rf(ctx, txClient, updateUser, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
