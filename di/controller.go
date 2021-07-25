package di

import "app/controller"

func InjectUserController() *controller.User {
	return controller.NewUser(
		InjectUserQuery(),
		InjectProfileQuery(),
		InjectUserUsecase(),
	)
}
