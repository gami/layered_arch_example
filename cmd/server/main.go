package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gami/layered_arch_example/controller"
	"github.com/gami/layered_arch_example/di"
	api "github.com/gami/layered_arch_example/gen/openapi"
)

func main() {
	port := 8080

	rt, err := controller.NewRouter()
	if err != nil {
		log.Fatal(err)
	}

	api.HandlerFromMux(
		controller.NewController(
			di.InjectUserController(),
		),
		rt,
	)

	s := &http.Server{
		Handler: rt,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}

	log.Fatal(s.ListenAndServe())
}
