package controller

import (
	"context"

	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/domain/user"
	"github.com/gami/layered_arch_example/usecase/form"
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
