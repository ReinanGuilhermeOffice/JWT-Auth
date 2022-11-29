package routes

import (
	"jwtauth/src/controllers"
	"net/http"
)

var RouteLogin = Route{
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	RequiresAuth: false,
}
