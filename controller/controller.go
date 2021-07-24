package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gami/layered_arch_example/controller/build"
)

type Controller struct {
	*User
}

func NewController(u *User) *Controller {
	return &Controller{
		User: u,
	}
}

func respondOK(w http.ResponseWriter, result interface{}) {
	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(build.Error(err))

		return
	}
}

func respond500(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	_ = json.NewEncoder(w).Encode(build.Error(err))
}
