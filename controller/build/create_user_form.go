package build

import (
	api "app/gen/openapi"
	"app/usecase/form"
)

func ToCreateUser(r *api.CreateUserJSONRequestBody) form.CreateUser {
	return form.CreateUser{
		Name:  r.Name,
		Hobby: *r.Hobby,
	}
}
