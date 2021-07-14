package user

import (
	"context"

	"github.com/friendsofgo/errors"
)

type Service interface {
	FindByID(ctx context.Context, id uint64) (*User, error)
	Create(ctx context.Context, u *User) (uint64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) FindByID(ctx context.Context, id uint64) (*User, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch user id=%v", id)
	}

	return u, nil
}

func (s *service) Create(ctx context.Context, u *User) (uint64, error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}

	id, err := s.repo.Create(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to create user id=%v", id)
	}

	return id, nil
}
