package repository

import (
	"context"
	"database/sql"

	"github.com/gami/layered_arch_example/adapter/mysql"
	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/gen/schema"
	"github.com/gami/layered_arch_example/repository/build"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Profile struct {
	db *sql.DB
}

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

func (r *Profile) FindByUserID(ctx context.Context, id uint64) (*profile.Profile, error) {
	s, err := schema.Profiles(
		schema.UserWhere.ID.EQ(id),
	).One(ctx, r.conn(ctx))

	if err != nil {
		return nil, errors.Wrapf(err, "failed to query profile id=%v", id)
	}
	return build.DomainProfile(s), nil
}

func (r *Profile) Create(ctx context.Context, u *profile.Profile) (uint64, error) {
	s := &schema.Profile{
		Hobby: StringMayNull(u.Hobby),
	}
	err := s.Insert(ctx, r.conn(ctx), boil.Infer())
	if err != nil {
		return 0, errors.Wrap(err, "failed to insert profile")
	}
	return s.ID, nil
}
