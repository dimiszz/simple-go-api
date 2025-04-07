package route

import "net/http"

type Route struct {
	Pattern string
	Handler http.Handler
}
