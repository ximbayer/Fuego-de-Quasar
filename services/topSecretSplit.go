package services

import (
	"encoding/json"
	"net/http"

	"github.com/ximbayer/Fuego-de-Quasar/models"
)

// Get satellites by GET into TopSecretSplit
func GetTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	var t models.Satellite
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}

	if t.Name == "" {
		http.Error(w, "The name of the satellite "+t.Name+"is required.", 400)
		return
	}

	/*
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
		if encontrado == true {
			http.Error(w, "Ya existe un usuario registrado con ese email", 400)
			return
		}

		_, status, err := bd.InsertoRegistro(t)
		if err != nil {
			http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
			return
		}

		if status == false {
			http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
			return
		}
	*/

	w.WriteHeader(http.StatusCreated)
}

// Post a satellite by POST into TopSecretSplit
func PostTopSecretSplit(w http.ResponseWriter, r *http.Request) {
	var t models.Satellite
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}

	if t.Name == "" {
		http.Error(w, "The name of the satellite "+t.Name+"is required.", 400)
		return
	}

	/*
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
		if encontrado == true {
			http.Error(w, "Ya existe un usuario registrado con ese email", 400)
			return
		}

		_, status, err := bd.InsertoRegistro(t)
		if err != nil {
			http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
			return
		}

		if status == false {
			http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
			return
		}
	*/

	w.WriteHeader(http.StatusCreated)
}
