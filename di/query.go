package di

import (
	"github.com/gami/layered_arch_example/controller"
	"github.com/gami/layered_arch_example/usecase/query"
)

func InjectUserQuery() controller.UserQuery {
	return query.NewUser(
		InjectUserService(),
	)
}

func InjectProfileQuery() controller.ProfileQuery {
	return query.NewProfile(
		InjectProfileService(),
	)
}
