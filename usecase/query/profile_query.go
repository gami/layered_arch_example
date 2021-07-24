package query

import (
	"context"

	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/domain/user"
	"github.com/gami/layered_arch_example/usecase"
)

type Profile struct {
	service usecase.ProfileService
}

func NewProfile(p usecase.ProfileService) *Profile {
	return &Profile{
		service: p,
	}
}

func (q *Profile) FindByUserID(ctx context.Context, userID user.ID) (*profile.Profile, error) {
	return q.service.FindByUserID(ctx, userID)
}
