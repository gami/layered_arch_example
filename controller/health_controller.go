package controller

import (
	"net/http"

	api "github.com/gami/layered_arch_example/gen/openapi"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

// GetHealth processes (GET /health)
func (c *Health) GetHealth(w http.ResponseWriter, r *http.Request) {
	res := api.OK{
		Message: "OK",
	}

	respondOK(w, res)
}
