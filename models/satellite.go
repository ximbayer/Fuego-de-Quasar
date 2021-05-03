package models

/*Satellite is the model for the satellites */
type Satellite struct {
	Name        string  `json:"name"`
	XCoordinate float64 `json:"x"`
	YCoordinate float64 `json:"y"`
}
