package routes

import (
	"jwtauth/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// modelo padrão das rotas da API
type Route struct {
	URI                 string
	Method              string
	Function            func(http.ResponseWriter, *http.Request)
	RequiresAuth        bool
	RequiresAuthRefresh bool
}

// Coloca todas as rotas dentro do router
func Configuration(r *mux.Router) *mux.Router {
	//adicionando rota login
	routes := RouteLogin

	routes = append(routes, RouteUser...)

	//retornando todas as rotas
	for _, route := range routes {
		if route.RequiresAuth {
			r.HandleFunc(route.URI, middlewares.Auth(route.Function)).Methods(route.Method)
		} else if route.RequiresAuthRefresh {
			r.HandleFunc(route.URI, middlewares.AuthRefresh(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)
		}

	}

	return r
}
