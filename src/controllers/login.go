package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"jwtauth/src/auth"
	"jwtauth/src/models"
	msgresponse "jwtauth/src/msgResponse"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		msgresponse.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//convertendo de JSON para STRUCT
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var userDatabase models.User
	userDatabase.ID = 1
	userDatabase.Email = "amigodeleo@gmail.com"
	userDatabase.Name = "Amigo de Leo"
	userDatabase.Password = "1234"

	if user.Email != userDatabase.Email || user.Password != userDatabase.Password {
		msgresponse.Erro(w, http.StatusInternalServerError, errors.New("Erro teste"))
		return
	}

	token, userID, erro := auth.CreateTokenLogin(userDatabase.ID)
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	userToken := userTokenLogin{userID, token}

	msgresponse.JSON(w, http.StatusCreated, userToken)
}

type userTokenLogin struct {
	ID    int64
	Token map[string]string
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		msgresponse.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	//convertendo de JSON para STRUCT
	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		msgresponse.Erro(w, http.StatusBadRequest, erro)
		return
	}

	refreshToken, userID, erro := auth.CreateTokenRefresh(user.ID)
	if erro != nil {
		msgresponse.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	userToken := userTokenLogin{userID, refreshToken}

	msgresponse.JSON(w, http.StatusCreated, userToken)
}
