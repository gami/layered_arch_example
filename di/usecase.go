package di

import (
	"app/controller"
	"app/usecase"
)

func InjectUserUsecase() controller.UserUsecase {
	return usecase.NewUser(
		InjectTx(),
		InjectUserService(),
		InjectProfileService(),
	)
}
