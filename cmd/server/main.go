package main

import (
	"fmt"
	"log"
	"net/http"

	"app/controller/rooter"
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
	rt, err := rooter.NewRouter()
	if err != nil {
		log.Fatal(err)
	}

	api.HandlerFromMux(
		di.InjectController(),
		rt,
	)

	return &http.Server{
		Handler: rt,
	}
}
