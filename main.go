package main

import (
	"fmt"
	"jwtauth/src/config"
	"jwtauth/src/router"
	"log"
	"net/http"
)

func main() {

	//carregando as configurações de ambiente
	config.Load()
	fmt.Println(config.SecretKey, config.SecretKeyRefresh)

	// pegando as rotas
	r := router.GenerateRouter()

	//iniciando instancia do servidor
	fmt.Println("Start Server!!")
	log.Fatal(http.ListenAndServe(":5000", r))
}
