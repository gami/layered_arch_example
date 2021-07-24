package controller

import (
	"net/http"

	validator "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	api "github.com/gami/layered_arch_example/gen/openapi"
	middleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
)

func NewRouter() (*chi.Mux, error) {
	swagger, err := api.GetSwagger()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load swagger")
	}

	r := chi.NewRouter()
	r.Use(recoverer)
	r.Use(middleware.RequestID)
	r.Use(validator.OapiRequestValidator(swagger))

	return r, nil
}

func recoverer(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rv := recover(); rv != nil {
				// log
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
