package di

import (
	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/domain/user"
	"github.com/gami/layered_arch_example/repository"
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
