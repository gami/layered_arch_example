package di

import (
	"app/controller"
	"app/usecase/query"
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
