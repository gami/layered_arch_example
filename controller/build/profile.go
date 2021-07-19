package build

import (
	"github.com/gami/layered_arch_example/domain/profile"
	api "github.com/gami/layered_arch_example/gen/openapi"
)

func fromDomainProfile(m *profile.Profile) api.Profile {
	return api.Profile{
		Hobby: &m.Hobby,
	}
}
