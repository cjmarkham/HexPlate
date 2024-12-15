package main

import (
	"fmt"
	"github.com/cjmarkham/hexplate/cmd/api"
	api2 "github.com/cjmarkham/hexplate/internal/api"
	"github.com/cjmarkham/hexplate/internal/api/forum"
	"github.com/go-chi/chi/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	handlers := injectApp()
	start(handlers)
}

func start(handlers forum.Handlers) {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		ErrorHandler: api2.HandleValidationError,
	}))

	hexForumHandler := api.NewStrictHandlerWithOptions(handlers, nil, api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  api2.HandleResponseError,
		ResponseErrorHandlerFunc: api2.HandleResponseError,
	})
	api.HandlerFromMux(hexForumHandler, r)

	server := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	log.Print("Starting server on port 8000")
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
