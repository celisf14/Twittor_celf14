package routers

import (
	"errors"

	"strings"

	"github.com/celisf14/Twittor_celf14/db"
	"github.com/celisf14/Twittor_celf14/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email valor de email usuado en todos los EndPoints*/
var Email string

/*IDUsuario el ID devuelto del modelo, sera usado en todos los EndPoints*/
var IDUsuario string

/*ProcesoToken procesa token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {

	miClave := []byte("MastersDelDesarollo_grupoDeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := db.ChequeoYaExisteUsuario(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}

	return claims, false, string(""), err

}
