package profile

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/gami/layered_arch_example/domain/failure"
	"github.com/gami/layered_arch_example/domain/user"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindByUserID(ctx context.Context, userID user.ID) (*Profile, error) {
	u, err := s.repo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch profile user_id=%v", userID)
	}

	return u, nil
}

func (s *Service) Create(ctx context.Context, p *Profile) (ID, error) {
	if err := p.Validate(); err != nil {
		return 0, failure.Invalid(err)
	}

	id, err := s.repo.Create(ctx, p)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to create profile id=%v", id)
	}

	return id, nil
}
