package router

import (
	"jwtauth/src/router/routes"

	"github.com/gorilla/mux"
)

// Vai retornar um router com as rotas configuradas.
func GenerateRouter() *mux.Router {
	r := mux.NewRouter()

	return routes.Configuration(r)
}
