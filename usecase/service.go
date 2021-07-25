package usecase

import (
	"context"

	"app/domain/profile"
	"app/domain/user"
)

type UserService interface {
	FindByID(ctx context.Context, id user.ID) (*user.User, error)
	Create(ctx context.Context, u *user.User) (user.ID, error)
}

type ProfileService interface {
	FindByUserID(ctx context.Context, userID user.ID) (*profile.Profile, error)
	Create(ctx context.Context, u *profile.Profile) (profile.ID, error)
}
