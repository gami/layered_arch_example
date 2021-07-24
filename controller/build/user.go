package build

import (
	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/domain/user"
	api "github.com/gami/layered_arch_example/gen/openapi"
)

type User struct {
	user api.User
}

func NewUser(u *user.User) *User {
	return &User{
		user: fromDomainUser(u),
	}
}

func (b *User) WithProfile(p *profile.Profile) *User {
	b.user.Profile = fromDomainProfile(p)
	return b
}

func (b *User) Build() api.User {
	return b.user
}

func fromDomainUser(m *user.User) api.User {
	return api.User{
		Id:   uint64(m.ID),
		Name: m.Name,
	}
}
