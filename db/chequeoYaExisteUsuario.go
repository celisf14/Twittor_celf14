package db

import (
	"context"
	"time"

	"github.com/celisf14/Twittor_celf14/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario Recibe un parametro y valida si ya existe en la DB*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("celisf14")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID

}
