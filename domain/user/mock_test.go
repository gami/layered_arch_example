package user_test

import (
	"context"

	"app/domain/user"
)

type MockRepository struct {
	FindByIDFunc func(ctx context.Context, id user.ID) (*user.User, error)
	CreateFunc   func(ctx context.Context, u *user.User) (user.ID, error)
}

func (m *MockRepository) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	return m.FindByIDFunc(ctx, id)
}
func (m *MockRepository) Create(ctx context.Context, u *user.User) (user.ID, error) {
	return m.CreateFunc(ctx, u)
}

type MockMessenger struct {
	SendCreatedFunc    func(ctx context.Context, msg *user.User) error
	RecieveCreatedFunc func(ctx context.Context, f func(msg *user.User) error) error
}

func (m *MockMessenger) SendCreated(ctx context.Context, msg *user.User) error {
	return m.SendCreatedFunc(ctx, msg)
}

func (m *MockMessenger) RecieveCreated(ctx context.Context, f func(msg *user.User) error) error {
	return m.RecieveCreatedFunc(ctx, f)
}
