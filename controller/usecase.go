package controller

import (
	"context"

	"app/domain/profile"
	"app/domain/user"
	"app/usecase/form"
)

type UserQuery interface {
	Find(ctx context.Context, id user.ID) (*user.User, error)
}

type UserUsecase interface {
	Create(ctx context.Context, input form.CreateUser) (uint64, error)
}

type ProfileQuery interface {
	FindByUserID(ctx context.Context, userID user.ID) (*profile.Profile, error)
}
