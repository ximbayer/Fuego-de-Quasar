package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ximbayer/Fuego-de-Quasar/functions"
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

//TotalSatellitesOperating is an array with the total satellites operating in Quasar
var TotalSatellitesOperating []models.Satellite = functions.GetAllSatellitesOperatingData()

//ShipRequest is a model to load all ship requests to satellites
var ShipRequest models.Request

//TopSecret is the service to obtain the ubication and urgency message of the ship
func TopSecret(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&ShipRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Incorrect data. "+err.Error(), 400)
		return
	}
	for _, satellite := range ShipRequest.ShipToSatellites {
		if satellite.NameSatell == "" {
			http.Error(w, "The name of all satellites is required.", 400)
			return
		}
	}

	//if we have more than 3 satellites, only we could get the message, but no the location
	if len(ShipRequest.ShipToSatellites) > 3 {
		fmt.Fprintf(w, "Too many satellites")
		return
	}

	//it goes through the satellites sent by the consumer for the ship (ShipToSatellites) and obtains their data
	coordinates := functions.GetCoordinates(ShipRequest.ShipToSatellites, TotalSatellitesOperating)
	distances := functions.GetDistances(ShipRequest.ShipToSatellites)
	messages := functions.GetMessages(ShipRequest.ShipToSatellites)

	//if coordinates slice is nil, no satellite of those analyzed is operational
	if coordinates == nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := "The location information of the ship could not be got since no satellite of those analyzed is operational"
		json.NewEncoder(w).Encode(errorMessage)
	}

	// get the ubication and urgency message of the ship
	//x, y := functions.GetLocation(distances...)
	//message := functions.GetMessage(messages...)
	x, y, message, errLocation, errMessage := functions.ProcessData(coordinates, distances, messages)

	//create a Position model to load
	pos := models.Position{XCoordinate: x, YCoordinate: y}

	resp := models.Response{Position: pos, Message: message}
	w.Header().Set("Content-Type", "application/json")

	//For validation: X= 9999999999 and Y=9999999999 is an incorrect coordinate. These values are to represent an error in the GetLocation
	if x == 9999999999 || y == 9999999999 || errLocation != "" { //|| message == ""  || errMessage != "" {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := "The location information of the ship could not be got"
		json.NewEncoder(w).Encode(errorMessage)
	} else if message == "" || errMessage != "" {
		w.WriteHeader(http.StatusNotFound)
		errorMessage := "The ship's distress message could not be got"
		json.NewEncoder(w).Encode(errorMessage)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
