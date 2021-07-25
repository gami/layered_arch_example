package query

import (
	"context"

	"app/domain/profile"
	"app/domain/user"
	"app/usecase"
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
