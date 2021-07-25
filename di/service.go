package di

import (
	"app/domain/profile"
	"app/domain/user"
	"app/usecase"
)

func InjectUserService() usecase.UserService {
	return user.NewService(
		InjectUserRepository(),
		InjectUserMessenger(),
	)
}

func InjectProfileService() usecase.ProfileService {
	return profile.NewService(
		InjectProfileRepository(),
	)
}
