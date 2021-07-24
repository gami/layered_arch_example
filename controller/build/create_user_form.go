package build

import (
	api "github.com/gami/layered_arch_example/gen/openapi"
	"github.com/gami/layered_arch_example/usecase/form"
)

func ToCreateUser(r *api.CreateUserJSONRequestBody) form.CreateUser {
	return form.CreateUser{
		Name:  r.Name,
		Hobby: *r.Hobby,
	}
}
