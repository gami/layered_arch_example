package controller

import (
	"context"

	"app/domain/profile"
	"app/domain/user"
	"app/usecase/form"
)

type UserUsecase interface {
	Find(ctx context.Context, id user.ID) (*user.User, error)
	Create(ctx context.Context, input form.CreateUser) (uint64, error)
}

type ProfileUsecase interface {
	FindByUserID(ctx context.Context, userID user.ID) (*profile.Profile, error)
}
