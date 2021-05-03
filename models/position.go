package models

/*Position is the model for the location of the ship */
type Position struct {
	XCoordinate float64 `json:"x"`
	YCoordinate float64 `json:"y"`
}
