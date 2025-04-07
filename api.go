package main

import (
	"log"
	"net/http"

	"dimi/server/middleware"
	"dimi/server/routing"
)

type APIServer struct {
	Addr   string
	Router *http.ServeMux
}

type APIConfig interface {
	routing.RouterConfig
	middleware.MiddlewareConfig
}

type DefaultAPIConfig struct {
	routing.DefaultRouterConfig
	middleware.DefaultMiddlewareConfig
}

func (s *APIServer) Run(config APIConfig) error {
	config.RegisterControllers(s.Router)

	server := &http.Server{
		Addr:    s.Addr,
		Handler: config.RegisterMiddlewares(s.Router),
	}

	log.Printf("server has started %s", s.Addr)

	return server.ListenAndServe()
}
