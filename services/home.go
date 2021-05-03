package services

import (
	"fmt"
	"net/http"
)

//// Home method: load home page "/" with a welcome message to API
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Rest API - Fuego de Quasar. \n\nDeveloped by Luciano Carreras.\n")
}
