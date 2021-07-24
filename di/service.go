package di

import (
	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/domain/user"
	"github.com/gami/layered_arch_example/usecase"
)

func InjectUserService() usecase.UserService {
	return user.NewService(
		InjectUserRepository(),
	)
}

func InjectProfileService() usecase.ProfileService {
	return profile.NewService(
		InjectProfileRepository(),
	)
}
