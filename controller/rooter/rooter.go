package rooter

import (
	api "app/gen/openapi"

	validator "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
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
	r.Use(middleware.Compress(5))
	r.Use(middleware.RequestID)
	r.Use(validator.OapiRequestValidator(swagger))

	return r, nil
}
