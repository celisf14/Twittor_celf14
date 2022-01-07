package middlew

import (
	"net/http"

	"github.com/celisf14/Twittor_celf14/db"
)

/*ChequeoDB es el middlew que me permite conocer el estado de la base de datos*/
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if !db.ChequeoConnection() {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
