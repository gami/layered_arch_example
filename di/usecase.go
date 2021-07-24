package di

import (
	"github.com/gami/layered_arch_example/controller"
	"github.com/gami/layered_arch_example/usecase"
)

func InjectUserUsecase() controller.UserUsecase {
	return usecase.NewUser(
		InjectTx(),
		InjectUserService(),
		InjectProfileService(),
	)
}
