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
