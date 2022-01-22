package routers

import (
	"net/http"

	"github.com/celisf14/Twittor_celf14/db"
	"github.com/celisf14/Twittor_celf14/models"
)

/*BajaRelacion realiza el borrado  de la relacion entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parametro id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := db.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Error al borrar la relacion", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se logro borrar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
