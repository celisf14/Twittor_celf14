package db

import (
	"context"
	"time"

	"github.com/celisf14/Twittor_celf14/models"
)

/*BorroRelacion borra la relacion en la base de datos*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
