package di

import "github.com/gami/layered_arch_example/domain/user"

func InjectUserService() user.Service {
	return user.NewService(
		InjectUserRepository(),
	)
}
