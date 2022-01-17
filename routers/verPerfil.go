package routers

import (
	"encoding/json"
	"net/http"

	"github.com/celisf14/Twittor_celf14/db"
)

/*verPerfil permie extraer los valores del perfil*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentaer buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(perfil)

}
