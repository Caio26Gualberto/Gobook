package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		Uri:                     "/users",
		Method:                  http.MethodPost,
		Function:                controllers.CreateUser,
		IsRequireAuthentication: false,
	},
	{
		Uri:                     "/users",
		Method:                  http.MethodGet,
		Function:                controllers.GetAllUsers,
		IsRequireAuthentication: false,
	},
	{
		Uri:                     "/users/{userId}",
		Method:                  http.MethodGet,
		Function:                controllers.GetUserById,
		IsRequireAuthentication: false,
	},
	{
		Uri:                     "/users/{userId}",
		Method:                  http.MethodPut,
		Function:                controllers.UpdateUser,
		IsRequireAuthentication: false,
	},
	{
		Uri:                     "/users/{userId}",
		Method:                  http.MethodDelete,
		Function:                controllers.DeleteUser,
		IsRequireAuthentication: false,
	},
}
