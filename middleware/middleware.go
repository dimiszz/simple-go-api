package middleware

import (
	"net/http"

	"dimi/server/middleware/logging"
)

type Middleware func(http.Handler) http.Handler

type MiddlewareConfig interface {
	RegisterMiddlewares(router http.Handler) http.Handler
}

type DefaultMiddlewareConfig struct {
	middlewares []Middleware
}

func (config *DefaultMiddlewareConfig) AddMiddlewares() {
	config.middlewares = []Middleware{
		logging.LoggingMiddleware,
	}
}

func (config *DefaultMiddlewareConfig) RegisterMiddlewares(router http.Handler) http.Handler {
	for _, middlewareFunc := range config.middlewares {
		router = middlewareFunc(router)
	}

	return router
}
