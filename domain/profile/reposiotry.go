package profile

import (
	"context"

	"app/domain/user"
)

type Repository interface {
	FindByUserID(ctx context.Context, userID user.ID) (*Profile, error)
	Create(ctx context.Context, p *Profile) (ID, error)
}
