package repository

import (
	"context"
	"database/sql"
	"fmt"

	"app/adapter/mysql"
	"app/domain/failure"
	"app/domain/user"
	"app/gen/schema"
	"app/infra/repository/build"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type User struct {
	db *mysql.DB
}

// build error will occur if User does not implement user.Repository.
var _ user.Repository = &User{}

func NewUser(db *mysql.DB) *User {
	return &User{
		db: db,
	}
}

// Conn returns usually wrapped *sql.DB connection. If the context is in transaction, this returns *sql.Tx.
func (r *User) conn(ctx context.Context) boil.ContextExecutor {
	tx, ok := GetTx(ctx)
	if !ok {
		return r.db.DB
	}

	return tx
}

func (r *User) FindByID(ctx context.Context, id user.ID) (*user.User, error) {
	s, err := schema.Users(
		schema.UserWhere.ID.EQ(uint64(id)),
	).One(ctx, r.conn(ctx))

	if errors.Is(err, sql.ErrNoRows) {
		return nil, failure.NotFound(fmt.Errorf("misssing user id=%d", id))
	} else if err != nil {
		return nil, errors.Wrapf(err, "failed to query user id=%d", id)
	}

	return build.DomainUser(s), nil
}

func (r *User) Create(ctx context.Context, u *user.User) (user.ID, error) {
	s := &schema.User{
		Name: u.Name,
	}

	err := s.Insert(ctx, r.conn(ctx), boil.Infer())
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert user")
	}

	return user.ID(s.ID), nil
}

func (r *User) FindAllByIDs(ctx context.Context, ids []user.ID) ([]*user.User, error) {
	s, err := schema.Users().All(ctx, r.conn(ctx))

	if err != nil {
		return nil, errors.Wrapf(err, "failed to query users")
	}

	return build.DomainUsers(s), nil
}
