package usecase

import (
	"context"

	"app/domain/profile"
	"app/domain/user"
)

type Profile struct {
	service ProfileService
}

func NewProfile(p ProfileService) *Profile {
	return &Profile{
		service: p,
	}
}

func (s *Profile) FindByUserID(ctx context.Context, userID user.ID) (*profile.Profile, error) {
	return s.service.FindByUserID(ctx, userID)
}
