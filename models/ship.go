package models

/*Ship is the model for the ships */
type ShipToSatellite struct {
	NameSatell string   `json:"name"`
	Distance   float64  `json:"distance"`
	Message    []string `json:"message"`
}
