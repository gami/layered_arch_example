package user_test

import (
	"context"

	"github.com/gami/layered_arch_example/domain"
	"github.com/gami/layered_arch_example/domain/user"
)

type MockRepository struct {
	FindByIDFunc func(ctx context.Context, id uint64) (*user.User, error)
	CreateFunc   func(ctx context.Context, u *user.User) (uint64, error)
}

func (m *MockRepository) FindByID(ctx context.Context, id uint64) (*user.User, error) {
	return m.FindByIDFunc(ctx, id)
}
func (m *MockRepository) Create(ctx context.Context, u *user.User) (uint64, error) {
	return m.CreateFunc(ctx, u)
}

type MockMessenger struct {
	SendFunc        func(ctx context.Context, key string, msg domain.Message) error
	RecieveUserFunc func(ctx context.Context, key string, f func(msg domain.Message) error) error
}

func (m *MockMessenger) Send(ctx context.Context, key string, msg domain.Message) error {
	return m.SendFunc(ctx, key, msg)
}

func (m *MockMessenger) RecieveUser(ctx context.Context, key string, f func(msg domain.Message) error) error {
	return m.RecieveUserFunc(ctx, key, f)
}
