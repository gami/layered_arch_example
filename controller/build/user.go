package build

import (
	"github.com/gami/layered_arch_example/domain/user"
	api "github.com/gami/layered_arch_example/gen/openapi"
)

func FromDomainUser(m *user.User) *api.User {
	return &api.User{
		Id: m.ID,
	}
}
