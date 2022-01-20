package routers

import (
	"encoding/json"
	"net/http"

	"github.com/celisf14/Twittor_celf14/db"
	"github.com/celisf14/Twittor_celf14/models"
)

/**ModificarPerfil modifica el perfil del usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	status, e := db.ModificoRegistro(t, IDUsuario)

	if e != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro, Reintente nuevamente"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se a logrado modificar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
