package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/celisf14/Twittor_celf14/db"
	"github.com/celisf14/Twittor_celf14/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro "+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{

		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, intente nuevamente "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se a logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
