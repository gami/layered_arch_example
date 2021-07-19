package profile

import "context"

type Repository interface {
	FindByUserID(ctx context.Context, userID uint64) (*Profile, error)
	Create(ctx context.Context, p *Profile) (uint64, error)
}
