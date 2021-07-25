package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"app/controller/build"
	"app/domain/user"
	api "app/gen/openapi"
)

type User struct {
	user    UserUsecase
	profile ProfileUsecase
}

func NewUser(
	u UserUsecase,
	p ProfileUsecase,
) *User {
	return &User{
		user:    u,
		profile: p,
	}
}

// GetUser processes (GET /user/{user_id})
func (c *User) GetUser(w http.ResponseWriter, r *http.Request, userID uint64) {
	ctx := context.Background()
	u, err := c.user.Find(ctx, user.ID(userID))

	if err != nil {
		RespondError(w, err)

		return
	}

	profile, err := c.profile.FindByUserID(ctx, u.ID)
	if err != nil {
		RespondError(w, err)

		return
	}

	res := build.NewUser(u).
		WithProfile(profile).
		Build()

	RespondOK(w, res)
}

// CreateUser processes (POST /users)
func (c *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var body *api.CreateUserJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		RespondError(w, err)

		return
	}

	id, err := c.user.Create(ctx, build.ToCreateUser(body))

	if err != nil {
		RespondError(w, err)

		return
	}

	res := build.Created(id)
	RespondOK(w, res)
}
