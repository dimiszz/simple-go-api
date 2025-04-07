package user

import (
	"net/http"

	"dimi/server/controller/route"
)

type UserController struct {
	routes []route.Route
	path   string
}

func (controller *UserController) AddRoutes() {
	controller.routes = []route.Route{
		{Pattern: "/hello", Handler: http.HandlerFunc(helloWorldHandler)},
	}
}

func (controller *UserController) GetPrefix() string {
	return controller.path
}

func (controller UserController) RegisterRoutes() *http.ServeMux {
	userRouter := http.NewServeMux()
	for _, route := range controller.routes {
		userRouter.HandleFunc(route.Pattern, route.Handler.ServeHTTP)
	}

	return userRouter
}

func NewUserController(prefix string) *UserController {
	return &UserController{path: prefix}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
