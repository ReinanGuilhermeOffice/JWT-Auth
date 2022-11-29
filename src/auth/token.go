package auth

import (
	"errors"
	"fmt"
	"jwtauth/src/config"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserToken struct {
	ID            int64
	access_token  string
	refresh_token string
}

func CreateTokenLogin(userID uint64) (map[string]string, int64, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Second * 10).Unix()
	permission["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	t, erro := token.SignedString([]byte(config.SecretKey))
	if erro != nil {
		return nil, 0, erro
	}

	permissionRefresh := jwt.MapClaims{}
	permissionRefresh["exp"] = time.Now().Add(time.Second * 19).Unix()
	permissionRefresh["userID"] = userID
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, permissionRefresh)
	rt, erro := refreshToken.SignedString([]byte(config.SecretKeyRefresh))
	if erro != nil {
		return nil, 0, erro
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, int64(userID), nil
}

func CreateTokenRefresh() {

}

func CreateToken(userID uint64) (string, error) {
	permission := jwt.MapClaims{}
	permission["authorized"] = true
	permission["exp"] = time.Now().Add(time.Second * 10).Unix()
	permission["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permission)
	return token.SignedString([]byte(config.SecretKey))
}

// verifica se o token passado na requisição é valido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	//extraindo dados do token
	token, erro := jwt.Parse(tokenString, returnKeyVerification)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido!")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	//verifica se o token que está sendo recebido está no formato esperado
	//Bearer + TOKEN
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}