package routes

import (
	"api/src/router/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: controllers.SearchUsers,
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.SearchUser,
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthRequired: false,
	},
}