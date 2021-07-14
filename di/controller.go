package di

import "github.com/gami/layered_arch_example/controller"

func InjectUserController() *controller.User {
	return controller.NewUser(
		InjectUserService(),
	)
}
