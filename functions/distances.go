package functions

import (
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

//return a list with all the distances that come in the request json
func GetDistances(shipToSatellites []models.ShipToSatellite) (distances []float64) {
	for _, ship := range shipToSatellites {
		distances = append(distances, ship.Distance)
	}
	return distances
}
