package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SecretKey        []byte
	SecretKeyRefresh []byte
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	SecretKey = []byte(os.Getenv("secret_key"))
	SecretKeyRefresh = []byte(os.Getenv("secret_key_refresh"))
}
