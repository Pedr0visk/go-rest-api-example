package routes

import (
	"api/src/controllers"
	"net/http"
)

var authRoutes = []Route{
	{
		Uri:       "/auth/login",
		Method:    http.MethodPost,
		Func:      controllers.Login,
		AuthGuard: false,
	},
}
