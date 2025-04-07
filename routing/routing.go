package routing

import (
	"fmt"
	"net/http"
	"regexp"

	"dimi/server/controller"
	"dimi/server/controller/user"
)

type RouterConfig interface {
	RegisterControllers(router *http.ServeMux)
}

type DefaultRouterConfig struct {
	Controllers []controller.Controller
}

func (config *DefaultRouterConfig) AddControllers() {
	config.Controllers =
		[]controller.Controller{
			user.NewUserController("/user/"),
		}
}

func (config *DefaultRouterConfig) RegisterControllers(router *http.ServeMux) {

	rgx, err := regexp.Compile(`/.*/`)
	if err != nil {
		panic(fmt.Sprintf("string %s in the wrong format: %s", rgx.String(), err.Error()))
	}

	for _, controller := range config.Controllers {
		prefix := controller.GetPrefix()
		controller.AddRoutes()
		controllerRouter := controller.RegisterRoutes()

		matched := rgx.Match([]byte(prefix))

		if !matched {
			panic(fmt.Sprintf("string %s in the wrong format", prefix))
		}

		router.Handle(prefix, http.StripPrefix(prefix[:len(prefix)-1], controllerRouter))
	}
}
