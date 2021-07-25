package build

import (
	api "app/gen/openapi"
)

func Created(id uint64) *api.Created {
	return &api.Created{
		Id:      id,
		Message: "Created",
	}
}
