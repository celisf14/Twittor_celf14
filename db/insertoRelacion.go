package db

import (
	"context"
	"time"

	"github.com/celisf14/Twittor_celf14/models"
)

/*InsertoRelacion graba la  relacion en la base de datos*/
func InsertoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}
