package routes

import (
	"jwtauth/src/controllers"
	"net/http"
)

var RouteLogin = []Route{
	{
		URI:                 "/login",
		Method:              http.MethodPost,
		Function:            controllers.Login,
		RequiresAuth:        false,
		RequiresAuthRefresh: false,
	},
	{
		URI:                 "/refreshtoken",
		Method:              http.MethodPost,
		Function:            controllers.RefreshToken,
		RequiresAuth:        false,
		RequiresAuthRefresh: true,
	},
}
