package build

import (
	"app/domain/user"
	"app/gen/schema"
)

func DomainUser(s *schema.User) *user.User {
	return &user.User{
		ID:   user.ID(s.ID),
		Name: s.Name,
	}
}

func DomainUsers(ss []*schema.User) []*user.User {
	ds := make([]*user.User, 0, len(ss))

	for _, s := range ss {
		ds = append(ds, DomainUser(s))
	}

	return ds
}
