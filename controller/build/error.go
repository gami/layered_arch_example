package build

import (
	api "app/gen/openapi"
)

func Error(err error) *api.Error {
	return &api.Error{
		Message: err.Error(),
	}
}
