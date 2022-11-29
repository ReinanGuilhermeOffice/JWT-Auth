package middlewares

import (
	"jwtauth/src/auth"
	msgresponse "jwtauth/src/msgResponse"
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidateToken(r); erro != nil {
			msgresponse.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
