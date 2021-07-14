package di

import (
	"github.com/gami/layered_arch_example/domain/user"
	"github.com/gami/layered_arch_example/repository"
)

func InjectUserRepository() user.Repository {
	return repository.NewUser(
		InjectUserDB(),
	)
}
