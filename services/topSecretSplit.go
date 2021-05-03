package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ximbayer/Fuego-de-Quasar/functions"
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

var NewShipRequest models.Request

// Get satellites by GET into TopSecretSplit
func GetTopSecretSplit(w http.ResponseWriter, r *http.Request) {

	//it goes through the satellites sent by the consumer for the ship (ShipToSatellites) and obtains their data
	coordinates := functions.GetCoordinates(NewShipRequest.ShipToSatellites, TotalSatellitesOperating)
	distances := functions.GetDistances(NewShipRequest.ShipToSatellites)
	messages := functions.GetMessages(NewShipRequest.ShipToSatellites)

	// get the ubication and urgency message of the ship
	//x, y := functions.GetLocation(distances...)
	//message := functions.GetMessage(messages...)
	x, y, message, _, _ := functions.ProcessData(coordinates, distances, messages)

	//create a Position model to load
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

// Post a satellite by POST into TopSecretSplit
func PostTopSecretSplit(w http.ResponseWriter, r *http.Request) {

	//a new model of request to load for the satellite in the Vars
	var aNewShipRequest models.ShipToSatellite

	//flag to detect if it is new or already exists
	var newSat = true

	vars := mux.Vars(r)
	nameSatell, errName := vars["satellite_name"]
	if errName == false {
		http.Error(w, "The name of the satellite is required.", 400)
		return
	}
	aNewShipRequest.NameSatell = nameSatell
	errShip := json.NewDecoder(r.Body).Decode(&aNewShipRequest)
	if errShip != nil {
		fmt.Fprintln(w, "Incorrect data. "+errShip.Error(), 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//if the satellite already exists then update it (distance and message)
	for i, sat := range NewShipRequest.ShipToSatellites {
		if aNewShipRequest.NameSatell == sat.NameSatell {
			NewShipRequest.ShipToSatellites[i].Distance = aNewShipRequest.Distance
			NewShipRequest.ShipToSatellites[i].Message = aNewShipRequest.Message
			newSat = false
			break
		}
	}

	//adding the new ship request to the satellite
	if newSat {
		NewShipRequest.ShipToSatellites = append(NewShipRequest.ShipToSatellites, aNewShipRequest)
	}

	//if we have more than 3 satellites, only we could get the message, but no the location
	if len(NewShipRequest.ShipToSatellites) == 3 {
		fmt.Fprintf(w, "Too many satellites")
		return
	}

	//TotalSatellitesOperating is defined in TopSecret.go. It belongs to the same package
	//it goes through the satellites sent by the consumer for the ship (ShipToSatellites) and obtains their data
	coordinates := functions.GetCoordinates(NewShipRequest.ShipToSatellites, TotalSatellitesOperating)
	distances := functions.GetDistances(NewShipRequest.ShipToSatellites)
	messages := functions.GetMessages(NewShipRequest.ShipToSatellites)

	// get the ubication and urgency message of the ship
	//x, y := functions.GetLocation(distances...)
	//message := functions.GetMessage(messages...)
	x, y, message, errLocation, errMessage := functions.ProcessData(coordinates, distances, messages)

	//create a Position model to load
	pos := models.Position{XCoordinate: x, YCoordinate: y}

	resp := models.Response{Position: pos, Message: message}
	w.Header().Set("Content-Type", "application/json")

	//For validation: X= 9999999999 and Y=9999999999 is an incorrect coordinate. These values are to represent an error in the GetLocation
	if x == 9999999999 || y == 9999999999 || message == "" || errLocation != "" || errMessage != "" {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := "The ship information could not be got"
		json.NewEncoder(w).Encode(errorMessage)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}

	w.WriteHeader(http.StatusCreated)
}
