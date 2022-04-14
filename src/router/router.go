package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Return configured routers
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.SetUp(r)
}
