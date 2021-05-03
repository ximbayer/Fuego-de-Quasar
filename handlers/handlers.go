package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/ximbayer/Fuego-de-Quasar/services"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Handlers() {
	//with StrictSlash: accept that the last slash of the url will be valid
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", services.Home)
	router.HandleFunc("/topsecret", services.TopSecret).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite_name}", services.PostTopSecretSplit).Methods("POST")
	router.HandleFunc("/topsecret_split", services.GetTopSecretSplit).Methods("GET")

	//search and define the PORT
	PORT := DefinePort()

	//give all security permissions to the handler
	handler := cors.AllowAll().Handler(router)

	//listen and serve from the PORT through the handler
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func DefinePort() string {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	return PORT
}
