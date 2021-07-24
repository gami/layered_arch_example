package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gami/layered_arch_example/controller/build"
	"github.com/gami/layered_arch_example/domain/user"
	api "github.com/gami/layered_arch_example/gen/openapi"
)

type User struct {
	userQuery    UserQuery
	profileQuery ProfileQuery
	userUsecase  UserUsecase
}

func NewUser(
	userQuery UserQuery,
	profileQuery ProfileQuery,
	userUsecase UserUsecase,
) *User {
	return &User{
		userQuery:    userQuery,
		profileQuery: profileQuery,
		userUsecase:  userUsecase,
	}
}

// GetUser processes (GET /user/{user_id})
func (c *User) GetUser(w http.ResponseWriter, r *http.Request, userID uint64) {
	ctx := context.Background()
	u, err := c.userQuery.Find(ctx, user.ID(userID))

	if err != nil {
		respondError(w, err)

		return
	}

	profile, err := c.profileQuery.FindByUserID(ctx, u.ID)
	if err != nil {
		respondError(w, err)

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
		respondError(w, err)

		return
	}

	id, err := c.userUsecase.Create(ctx, build.ToCreateUser(body))

	if err != nil {
		respondError(w, err)

		return
	}

	res := build.Created(id)
	respondOK(w, res)
}
