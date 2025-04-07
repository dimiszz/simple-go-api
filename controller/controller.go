package controller

import (
	"net/http"
)

type Controller interface {
	RegisterRoutes() *http.ServeMux
	GetPrefix() string
	AddRoutes()
}
