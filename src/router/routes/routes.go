package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct for routes
type Route struct {
	Uri       string
	Method    string
	Func      func(http.ResponseWriter, *http.Request)
	AuthGuard bool
}

func SetUp(r *mux.Router) *mux.Router {
	routes := append(usersRoutes, authRoutes...)

	for _, route := range routes {
		httpHandleFunc := middlewares.Logger(route.Func)
		if route.AuthGuard {
			r.HandleFunc(route.Uri, middlewares.Authenticate(httpHandleFunc)).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, httpHandleFunc).Methods(route.Method)
		}
	}

	return r
}
