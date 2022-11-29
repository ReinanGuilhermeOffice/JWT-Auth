package routes

import (
	"jwtauth/src/controllers"
	"net/http"
)

var RouteUser = []Route{
	{
		URI:          "/user",
		Method:       http.MethodPost,
		Function:     func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuth: false,
	},
	{
		URI:          "/user",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		RequiresAuth: true,
	},
}
