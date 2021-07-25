package main

import (
	"fmt"
	"log"
	"net/http"

	"app/controller"
	"app/di"

	api "app/gen/openapi"
)

func main() {
	port := 8080
	s := Server()
	s.Addr = fmt.Sprintf("0.0.0.0:%d", port)
	log.Fatal(s.ListenAndServe())
}

func Server() *http.Server {
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

	return &http.Server{
		Handler: rt,
	}
}
