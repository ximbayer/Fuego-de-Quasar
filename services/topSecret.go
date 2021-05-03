package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ximbayer/Fuego-de-Quasar/functions"
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

//totalSatellitesOperating is an array with the total satellites operating in Quasar
var TotalSatellitesOperating []models.Satellite

/*Registro es la funcion para crear en la BD el registro de usuario */
func TopSecret(w http.ResponseWriter, r *http.Request) {
	TotalSatellitesOperating = append(TotalSatellitesOperating,
		models.Satellite{Name: "Kenobi", XCoordinate: -500, YCoordinate: -200},
		models.Satellite{Name: "Skywalker", XCoordinate: 100, YCoordinate: -100},
		models.Satellite{Name: "Sato", XCoordinate: 500, YCoordinate: 100})
	var request models.Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		//fmt.Fprintln(w, "Incorrect data. "+err.Error(), request)
		fmt.Fprintln(w, "Incorrect data. "+err.Error(), 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, satellite := range request.Satellites {
		if satellite.NameSatell == "" {
			http.Error(w, "The name of all satellites is required.", 400)
			return
		}
	}
	//recorre los satelites enviados por el consumidor
	satellitesRequest := functions.GetSatellites(request)
	coordinates := functions.GetCoordinates(satellitesRequest, TotalSatellitesOperating)
	distances := functions.GetDistances(request)

	messages := functions.GetMessages(request)

	// obtiene la ubicacion de la nave que envio los llamados de auxilio.
	//x, y := functions.GetLocation(distances...)
	x, y, message, _, _ := functions.ProcessData(coordinates, distances, messages)
	pos := models.Position{XCoordinate: x, YCoordinate: y}

	resp := models.Response{Position: pos, Message: message}
	w.Header().Set("Content-Type", "application/json")

	//For validation: X= 9999999999 and Y=9999999999 is an incorrect coordinate. These values are to represent an error in the GetLocation
	if x == 9999999999 || y == 9999999999 || message == "" {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := "The ship information could not be got"
		json.NewEncoder(w).Encode(errorMessage)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}

	w.WriteHeader(http.StatusCreated)
}
