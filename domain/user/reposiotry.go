package user

import "context"

type Repository interface {
	FindByID(ctx context.Context, id ID) (*User, error)
	FindAllByIDs(ctx context.Context, ids []ID) ([]*User, error)
	Create(ctx context.Context, u *User) (ID, error)
}
