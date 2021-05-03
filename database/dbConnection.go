package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Connector to the DB
var Connector = DBConnect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://ximbayer:eav-1234@clustermeli.89pnn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//DBConnect is the function that allows me to connect the BD
func DBConnect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		// return a nil client
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		//We have a client but it cannot send or receive the ping
		return client
	}

	log.Println("Successful DB Connection")
	return client
}

//ConnectionCheck is the ping to DB by the Connetor
func ConnectionCheck() int {
	err := Connector.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
