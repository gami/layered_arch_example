package query

import (
	"context"

	"app/domain/user"
	"app/usecase"
)

type User struct {
	service usecase.UserService
}

func NewUser(u usecase.UserService) *User {
	return &User{
		service: u,
	}
}

func (q *User) Find(ctx context.Context, id user.ID) (*user.User, error) {
	return q.service.FindByID(ctx, id)
}
