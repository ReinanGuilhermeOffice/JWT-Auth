package repositories

import "jwtauth/src/models"

var user models.User

func UserDatabase() models.User {
	user.ID = 1
	user.Email = "amigodeleo@gmail.com"
	user.Name = "Amigo de Leo"
	user.Password = "1234"
	return user
}
