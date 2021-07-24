package usecase

import (
	"context"

	"github.com/friendsofgo/errors"
	"github.com/gami/layered_arch_example/domain"
	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/domain/user"
	"github.com/gami/layered_arch_example/usecase/form"
)

type User struct {
	tx      domain.Tx
	user    user.Service
	profile profile.Service
}

func NewUser(tx domain.Tx, user user.Service, profile profile.Service) *User {
	return &User{
		tx:      tx,
		user:    user,
		profile: profile,
	}
}

func (s *User) Create(ctx context.Context, input form.CreateUser) (uint64, error) {
	u := &user.User{
		Name: input.Name,
	}

	res, err := s.tx.Transact(ctx, func(c context.Context) (interface{}, error) {
		id, err := s.user.Create(ctx, u)
		if err != nil {
			return 0, errors.Wrap(err, "failed to create user")
		}

		p := &profile.Profile{
			UserID: u.ID,
		}

		_, err = s.profile.Create(ctx, p)
		if err != nil {
			return 0, errors.Wrapf(err, "failed to create profile user_id=%d", id)
		}

		return id, nil
	})

	if err != nil {
		return 0, nil
	}

	return res.(uint64), nil
}
