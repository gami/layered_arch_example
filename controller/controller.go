package controller

import (
	"encoding/json"
	"net/http"

	"app/controller/build"
	"app/domain/failure"

	"github.com/pkg/errors"
)

type Controller struct {
	*User
	*Health
}

func NewController(u *User) *Controller {
	return &Controller{
		Health: NewHealth(),
		User:   u,
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
			w.WriteHeader(http.StatusBadRequest)
		case failure.ErrForbidden:
			w.WriteHeader(http.StatusForbidden)
		case failure.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		case failure.ErrConflict:
			w.WriteHeader(http.StatusConflict)
		default:
			// TOOD log unknown code
			w.WriteHeader(http.StatusInternalServerError)
		}

		_ = json.NewEncoder(w).Encode(build.Error(err))
		return
	}

	// TODO log fatal
	w.WriteHeader(500)
	_ = json.NewEncoder(w).Encode(build.Error(err))
}
