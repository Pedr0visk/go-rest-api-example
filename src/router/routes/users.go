package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		Uri:       "/users",
		Method:    http.MethodGet,
		Func:      controllers.GetUsers,
		AuthGuard: true,
	},
	{
		Uri:       "/users",
		Method:    http.MethodPost,
		Func:      controllers.CreateUser,
		AuthGuard: true,
	},
	{
		Uri:       "/users/{userId}",
		Method:    http.MethodGet,
		Func:      controllers.GetUser,
		AuthGuard: true,
	},
	{
		Uri:       "/users/{userId}",
		Method:    http.MethodPut,
		Func:      controllers.UpdateUser,
		AuthGuard: true,
	},
	{
		Uri:       "/users/{userId}",
		Method:    http.MethodDelete,
		Func:      controllers.DeleteUser,
		AuthGuard: true,
	},
}
