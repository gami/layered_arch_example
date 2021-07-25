package di

import (
	"app/domain/profile"
	"app/domain/user"
	"app/infra/repository"
)

func InjectUserRepository() user.Repository {
	return repository.NewUser(
		InjectUserDB(),
	)
}

func InjectProfileRepository() profile.Repository {
	return repository.NewProfile(
		InjectUserDB(),
	)
}
