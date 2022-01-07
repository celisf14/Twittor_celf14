package main

import (
	"log"

	"github.com/celisf14/Twittor_celf14/db"
	"github.com/celisf14/Twittor_celf14/handlers"
)

func main() {

	if !db.ChequeoConnection() {
		log.Fatal("Sin conexion a la base de datos")
		return
	}

	handlers.Manejadores()

}
