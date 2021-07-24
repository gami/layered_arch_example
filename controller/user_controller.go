package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gami/layered_arch_example/controller/build"
	"github.com/gami/layered_arch_example/domain/profile"
	"github.com/gami/layered_arch_example/domain/user"
	api "github.com/gami/layered_arch_example/gen/openapi"
)

type User struct {
	user        user.Service
	profile     profile.Service
	userUsecase UserUsecase
}

func NewUser(
	user user.Service,
	profile profile.Service,
	userUsecase UserUsecase,
) *User {
	return &User{
		user:        user,
		profile:     profile,
		userUsecase: userUsecase,
	}
}

// GetUser processes (GET /user/{user_id})
func (c *User) GetUser(w http.ResponseWriter, r *http.Request, userId uint64) {
	ctx := context.Background()
	u, err := c.user.FindByID(ctx, user.ID(userId))
	if err != nil {
		respond500(w, err)
		return
	}

	profile, err := c.profile.FindByUserID(ctx, u.ID)
	if err != nil {
		respond500(w, err)
		return
	}

	res := build.NewUser(u).
		WithProfile(profile).
		Build()

	respondOK(w, res)
}

// CreateUser processes (POST /users)
func (c *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var body *api.CreateUserJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		respond500(w, err)
		return
	}

	id, err := c.userUsecase.Create(ctx, build.ToCreateUser(body))

	if err != nil {
		respond500(w, err)
		return
	}

	res := build.Created(id)
	respondOK(w, res)
}
