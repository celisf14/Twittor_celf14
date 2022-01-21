package routers

import (
	"net/http"

	"github.com/celisf14/Twittor_celf14/db"
)

/*EliminarTweet permite borrar un tweet determinado*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := db.BorroTweet(ID, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar un tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
