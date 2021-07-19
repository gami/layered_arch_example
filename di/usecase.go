package di

import (
	"github.com/gami/layered_arch_example/controller"
	"github.com/gami/layered_arch_example/usecase"
)

func InjectCreateUserUsecase() controller.CreateUser {
	return usecase.NewCreateUser(
		InjectTx(),
		InjectUserService(),
		InjectProfileService(),
	)
}
