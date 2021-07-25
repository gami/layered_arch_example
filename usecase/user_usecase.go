package usecase

import (
	"context"

	"app/domain"
	"app/domain/profile"
	"app/domain/user"
	"app/usecase/form"

	"github.com/friendsofgo/errors"
)

type User struct {
	tx      domain.Tx
	user    UserService
	profile ProfileService
}

func NewUser(tx domain.Tx, u UserService, p ProfileService) *User {
	return &User{
		tx:      tx,
		user:    u,
		profile: p,
	}
}

func (s *User) Find(ctx context.Context, id user.ID) (*user.User, error) {
	return s.user.FindByID(ctx, id)
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
