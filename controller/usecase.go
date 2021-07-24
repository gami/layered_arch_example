package controller

import (
	"context"

	"github.com/gami/layered_arch_example/usecase/form"
)

type UserUsecase interface {
	Create(ctx context.Context, input form.CreateUser) (uint64, error)
}
