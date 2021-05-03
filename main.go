package main

import (
	"log"

	"github.com/ximbayer/Fuego-de-Quasar/database"
	"github.com/ximbayer/Fuego-de-Quasar/handlers"
)

func main() {
	//ConnectionCheck return 1 if the connector has the ping OK to DB,
	//or 0 if the ping return an error
	if database.ConnectionCheck() == 0 {
		log.Fatal("No connection to the DB")
		return
	}
	handlers.Handlers()
}
