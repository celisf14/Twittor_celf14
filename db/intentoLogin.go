package db

import (
	"github.com/celisf14/Twittor_celf14/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el chequeo del login en la base de datos*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usu, encontro, _ := ChequeoYaExisteUsuario(email)

	if !encontro {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
