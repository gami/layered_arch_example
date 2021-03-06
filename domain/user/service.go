package user

import (
	"context"

	"app/domain/failure"

	"github.com/friendsofgo/errors"
)

type Service struct {
	repo Repository
	msgs Messenger
}

func NewService(repo Repository, msgs Messenger) *Service {
	return &Service{
		repo: repo,
		msgs: msgs,
	}
}

func (s *Service) FindByID(ctx context.Context, id ID) (*User, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) FindAllByIDs(ctx context.Context, ids []ID) ([]*User, error) {
	return s.repo.FindAllByIDs(ctx, ids)
}

func (s *Service) Create(ctx context.Context, u *User) (ID, error) {
	if err := u.Validate(); err != nil {
		return 0, failure.Invalid(err)
	}

	id, err := s.repo.Create(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to create user id=%v", id)
	}

	u.ID = id

	err = s.msgs.SendCreated(ctx, u)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to send user_created message id=%v", id)
	}

	return id, nil
}
