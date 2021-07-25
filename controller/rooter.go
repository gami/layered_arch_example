package controller

import (
	"log"
	"net/http"
	"runtime"

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

func recoverer(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// TODO
				log.Printf("[ERROR] %s\n", err)
				for depth := 0; ; depth++ {
					_, file, line, ok := runtime.Caller(depth)
					if !ok {
						break
					}
					log.Printf("======> %d: %v:%d", depth, file, line)
				}

				respondError(w, errors.New("panic"))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(f)
}
