package build

import (
	"app/domain/profile"
	api "app/gen/openapi"
)

func fromDomainProfile(m *profile.Profile) api.Profile {
	return api.Profile{
		Hobby: &m.Hobby,
	}
}
