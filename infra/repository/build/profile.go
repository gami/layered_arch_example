package build

import (
	"app/domain/profile"
	"app/gen/schema"
)

func DomainProfile(s *schema.Profile) *profile.Profile {
	return &profile.Profile{
		ID:    s.ID,
		Hobby: s.Hobby.String,
	}
}
