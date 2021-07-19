package build

import (
	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/gen/schema"
)

func DomainProfile(s *schema.Profile) *profile.Profile {
	return &profile.Profile{
		ID:    s.ID,
		Hobby: s.Hobby.String,
	}
}
