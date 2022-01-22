package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/celisf14/Twittor_celf14/db"
)

/*LeoTweetsRelacion lee los tweets de los seguidores*/
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a cero", http.StatusBadRequest)
		return
	}

	respuesta, status := db.LeoTweetsSeguidores(IDUsuario, pagina)

	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
