package user

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/gami/layered_arch_example/domain"
)

type Service interface {
	FindByID(ctx context.Context, id ID) (*User, error)
	Create(ctx context.Context, u *User) (ID, error)
}

type service struct {
	repo Repository
	msgs domain.Messenger
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) FindByID(ctx context.Context, id ID) (*User, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch user id=%v", id)
	}

	return u, nil
}

func (s *service) Create(ctx context.Context, u *User) (ID, error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}

	id, err := s.repo.Create(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to create user id=%v", id)
	}

	u.ID = id

	err = s.msgs.Send(ctx, domain.KeyUserCreated, u)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to send user_created message id=%v", id)
	}

	return id, nil
}
