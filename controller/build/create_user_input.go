package build

import (
	api "github.com/gami/layered_arch_example/gen/openapi"
	"github.com/gami/layered_arch_example/usecase"
)

func ToCreateUserInput(r *api.CreateUserJSONRequestBody) *usecase.CreateUserInput {
	return &usecase.CreateUserInput{
		Name:  r.Name,
		Hobby: *r.Hobby,
	}
}
