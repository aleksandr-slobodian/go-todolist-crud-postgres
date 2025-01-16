package store

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUsers struct {
	mock.Mock
}

func (m *MockUsers) Create(ctx context.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

type MockTodos struct {
	mock.Mock
}

func (m *MockTodos) GetByID(ctx context.Context, id int64) (*Todo, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*Todo), args.Error(1)
}

func (m *MockTodos) Create(ctx context.Context, todo *Todo) error {
	args := m.Called(ctx, todo)
	return args.Error(0)
}

func (m *MockTodos) Update(ctx context.Context, todo *Todo) error {
	args := m.Called(ctx, todo)
	return args.Error(0)
}

func (m *MockTodos) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTodos) GetTodos(ctx context.Context, params TodosQueryParams) ([]*Todo, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]*Todo), args.Error(1)
}
