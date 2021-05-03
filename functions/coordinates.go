package functions

import (
	"strings"

	"github.com/ximbayer/Fuego-de-Quasar/models"
)

//return a list with all the coordinate points (X and Y) that the operational satellites have
func GetCoordinates(shipToSatellites []models.ShipToSatellite, totalSatellitesOperating []models.Satellite) (coordinates []Point) {
	for _, satellite := range shipToSatellites {
		for _, satelliteOfTotal := range totalSatellitesOperating {
			if strings.ToUpper(satellite.NameSatell) == strings.ToUpper(satelliteOfTotal.Name) {
				coordinates = append(coordinates, Point{X: satelliteOfTotal.XCoordinate, Y: satelliteOfTotal.YCoordinate})
			}
		}
	}
	return coordinates
}
