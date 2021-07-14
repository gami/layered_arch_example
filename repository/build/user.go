package build

import (
	"github.com/gami/layered_arch_example/domain/user"
	"github.com/gami/layered_arch_example/gen/schema"
)

func DomainUser(s *schema.User) *user.User {
	return &user.User{
		ID: s.ID,
	}
}
