package profile

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/gami/layered_arch_example/domain/user"
)

type Service interface {
	FindByUserID(ctx context.Context, userID user.ID) (*Profile, error)
	Create(ctx context.Context, u *Profile) (ID, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) FindByUserID(ctx context.Context, userId user.ID) (*Profile, error) {
	u, err := s.repo.FindByUserID(ctx, userId)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch profile user_id=%v", userId)
	}

	return u, nil
}

func (s *service) Create(ctx context.Context, p *Profile) (ID, error) {
	if err := p.Validate(); err != nil {
		return 0, err
	}

	id, err := s.repo.Create(ctx, p)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to create profile id=%v", id)
	}

	return id, nil
}
