package controller

import (
	"context"
	"net/http"

	"github.com/gami/layered_arch_example/controller/build"
	"github.com/gami/layered_arch_example/domain/user"
)

type User struct {
	user user.Service
}

func NewUser(user user.Service) *User {
	return &User{
		user: user,
	}
}

// GetUser processes (GET /user/{user_id})
func (c *User) GetUser(w http.ResponseWriter, r *http.Request, userId uint64) {
	ctx := context.Background()
	user, err := c.user.FindByID(ctx, userId)
	if err != nil {
		respond500(w, err)
		return
	}

	res := build.FromDomainUser(user)
	respondOK(w, res)
}

// CreateUser processes (POST /users)
func (c *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	id, err := c.user.Create(ctx, &user.User{})
	if err != nil {
		respond500(w, err)
		return
	}

	res := build.Created(id)
	respondOK(w, res)
}
