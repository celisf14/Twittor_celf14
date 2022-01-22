package db

import (
	"context"
	"time"

	"github.com/celisf14/Twittor_celf14/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores lee los twees de los seguidores*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvotweetsSeguidores, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	limit := 20
	skip := (pagina - 1) * limit

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": limit})

	cursor, _ := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvotweetsSeguidores
	err := cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true

}
