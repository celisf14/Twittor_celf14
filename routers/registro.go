package routers

import (
	"encoding/json"
	"net/http"

	"github.com/celisf14/Twittor_celf14/db"
	"github.com/celisf14/Twittor_celf14/models"
)

/*Registro es la funcion para crear en la DB el registro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Error, el Email de usuario es requeido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Error, debe especificar una contraseña de almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := db.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		http.Error(w, "Error, ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := db.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Error, al intentar realizar el registro del usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Error, No se a logrado insertar el rgistro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
