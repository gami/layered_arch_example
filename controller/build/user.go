package build

import (
	"app/domain/profile"
	"app/domain/user"
	api "app/gen/openapi"
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
