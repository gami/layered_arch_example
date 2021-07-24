package user

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/gami/layered_arch_example/domain"
)

type Service struct {
	repo Repository
	msgs domain.Messenger
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindByID(ctx context.Context, id ID) (*User, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch user id=%v", id)
	}

	return u, nil
}

func (s *Service) Create(ctx context.Context, u *User) (ID, error) {
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
