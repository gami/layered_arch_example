package repository

import (
	"context"
	"database/sql"

	"github.com/gami/layered_arch_example/domain/user"
	"github.com/gami/layered_arch_example/gen/schema"
	"github.com/gami/layered_arch_example/mysql"
	"github.com/gami/layered_arch_example/repository/build"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type User struct {
	db *sql.DB
}

func NewUser(db *mysql.DB) user.Repository {
	return &User{
		db: db.DB,
	}
}

func (r *User) FindByID(ctx context.Context, id uint64) (*user.User, error) {
	s, err := schema.Users(
		schema.UserWhere.ID.EQ(id),
	).One(ctx, r.db)

	if err != nil {
		return nil, errors.Wrapf(err, "failed to query user id=%v", id)
	}
	return build.DomainUser(s), nil
}

func (r *User) Create(ctx context.Context, u *user.User) (uint64, error) {
	s := &schema.User{}
	err := s.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert user")
	}
	return s.ID, nil
}
