package build

import (
	api "github.com/gami/layered_arch_example/gen/openapi"
)

func Error(err error) *api.Error {
	return &api.Error{
		Message: err.Error(),
	}
}
