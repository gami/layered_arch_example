package user

import "context"

type Repository interface {
	FindByID(ctx context.Context, id uint64) (*User, error)
	Create(ctx context.Context, u *User) (uint64, error)
}
