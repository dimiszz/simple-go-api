package main

import (
	"net/http"

	"dimi/server/middleware"
	"dimi/server/routing"
)

func main() {
	router := http.NewServeMux()
	server := &APIServer{
		Addr:   ":10012",
		Router: router,
	}
	routerConfig := routing.DefaultRouterConfig{}
	middlewareConfig := middleware.DefaultMiddlewareConfig{}

	middlewareConfig.AddMiddlewares()

	defaultConfig := &DefaultAPIConfig{
		DefaultRouterConfig:     routerConfig,
		DefaultMiddlewareConfig: middlewareConfig,
	}

	err := server.Run(defaultConfig)
	if err != nil {
		panic(err)

	}
}
