package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la base de datos*/
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://celisf14_user:Zz871203@celisf14.a8m5z.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/*ConectarDB es la funcion que permite conectar la BD*/
func ConectarDB() *mongo.Client {

	log.Println("inicia conexion :) ")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println("conexion :) 1")
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("conexion :) 2")
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa a la base de datos")
	return client
}

/*ChequeoConnection es el ping a la base de datos*/
func ChequeoConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
