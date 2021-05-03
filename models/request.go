package models

/*Request is the model for the request to the API */
type Request struct {
	ShipToSatellites []ShipToSatellite `json:"ships"`
}
