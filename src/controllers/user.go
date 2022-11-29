package controllers

import (
	msgresponse "jwtauth/src/msgResponse"
	"jwtauth/src/repositories"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	msgresponse.JSON(w, http.StatusOK, repositories.UserDatabase())
}
