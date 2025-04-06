package routing

import "net/http"

type RouterConfig interface {
	RegisterRoutes(router *http.ServeMux)
}

type DefaultRouterConfig struct {
}

type Route struct {
	Pattern string
	Handler http.Handler
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (config *DefaultRouterConfig) RegisterRoutes(router *http.ServeMux) {
	routes := []Route{
		{Pattern: "/hello", Handler: http.HandlerFunc(helloWorldHandler)},
		{Pattern: "/about", Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("About Page"))
		})},
	}

	for _, route := range routes {
		router.Handle(route.Pattern, route.Handler)
	}
}
