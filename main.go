package main

import (
	"fmt"
	"jwtauth/src/config"
	"jwtauth/src/router"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {

	//carregando as configurações de ambiente
	config.Load()

	// pegando as rotas
	r := router.GenerateRouter()

	//liberando cors
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "DELETE", "POST", "PUT"},
	})

	handler := c.Handler(r)

	//iniciando instancia do servidor
	fmt.Println("Start Server!!")
	log.Fatal(http.ListenAndServe(":5000", handler))
}
