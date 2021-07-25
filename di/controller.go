package di

import "app/controller"

func InjectController() *controller.Controller {
	return controller.NewController(
		InjectUserController(),
	)
}

func InjectUserController() *controller.User {
	return controller.NewUser(
		InjectUserUsecase(),
		InjectProfileUsecase(),
	)
}
