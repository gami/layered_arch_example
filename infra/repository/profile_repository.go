package repository

import (
	"context"
	"database/sql"
	"fmt"

	"app/adapter/mysql"
	"app/domain/failure"
	"app/domain/profile"
	"app/domain/user"
	"app/gen/schema"
	"app/infra/repository/build"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Profile struct {
	db *sql.DB
}

// build error will occur if User does not implement user.Repository.
var _ profile.Repository = &Profile{}

func NewProfile(db *mysql.DB) profile.Repository {
	return &Profile{
		db: db.DB,
	}
}

// Conn returns usually wrapped *sql.DB connection. If the context is in transaction, this returns *sql.Tx.
func (r *Profile) conn(ctx context.Context) boil.ContextExecutor {
	tx, ok := GetTx(ctx)
	if !ok {
		return r.db
	}

	return tx
}

func (r *Profile) FindByUserID(ctx context.Context, userID user.ID) (*profile.Profile, error) {
	s, err := schema.Profiles(
		schema.UserWhere.ID.EQ(uint64(userID)),
	).One(ctx, r.conn(ctx))

	if errors.Is(err, sql.ErrNoRows) {
		return nil, failure.NotFound(fmt.Errorf("misssing profile user_id=%d", userID))
	} else if err != nil {
		return nil, errors.Wrapf(err, "failed to query profile user_id=%d", userID)
	}

	return build.DomainProfile(s), nil
}

func (r *Profile) Create(ctx context.Context, u *profile.Profile) (profile.ID, error) {
	s := &schema.Profile{
		Hobby: StringMayNull(u.Hobby),
	}

	err := s.Insert(ctx, r.conn(ctx), boil.Infer())
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert profile")
	}

	return profile.ID(s.ID), nil
}
