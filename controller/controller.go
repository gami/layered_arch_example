package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gami/layered_arch_example/controller/build"
	"github.com/gami/layered_arch_example/domain/failure"
	"github.com/pkg/errors"
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
		respondError(w, errors.Wrap(err, "faile to respond_ok"))

		return
	}
}

func respondError(w http.ResponseWriter, err error) {
	var ae *failure.AppError
	if errors.As(err, &ae) {
		// TODO log info
		switch ae.Code {
		case failure.ErrInvalid:
			w.WriteHeader(400)
		case failure.ErrForbidden:
			w.WriteHeader(403)
		case failure.ErrNotFound:
			w.WriteHeader(404)
		case failure.ErrConflict:
			w.WriteHeader(409)
		default:
			// TOOD log unknown code
			w.WriteHeader(500)
		}

		_ = json.NewEncoder(w).Encode(build.Error(err))
		return
	}

	// TODO log fatal
	w.WriteHeader(500)
	_ = json.NewEncoder(w).Encode(build.Error(err))
}
