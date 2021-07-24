package profile

import (
	"context"

	"github.com/gami/layered_arch_example/domain/user"
)

type Repository interface {
	FindByUserID(ctx context.Context, userID user.ID) (*Profile, error)
	Create(ctx context.Context, p *Profile) (ID, error)
}
