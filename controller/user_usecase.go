package controller

import (
	"context"

	"github.com/gami/layered_arch_example/usecase"
)

type CreateUser interface {
	Do(ctx context.Context, input *usecase.CreateUserInput) (uint64, error)
}
