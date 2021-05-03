package functions

import (
	"strings"

	"github.com/ximbayer/Fuego-de-Quasar/models"
)

func GetCoordinates(satellitesRequest []string, totalSatellitesOperating []models.Satellite) (coordinates []Point) {
	for _, satellite := range satellitesRequest {
		for _, satelliteOfTotal := range totalSatellitesOperating {
			if strings.ToUpper(satellite) == strings.ToUpper(satelliteOfTotal.Name) {
				coordinates = append(coordinates, Point{X: satelliteOfTotal.XCoordinate, Y: satelliteOfTotal.YCoordinate})
			}
		}
	}
	return coordinates
}
