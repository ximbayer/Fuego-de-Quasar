package functions

import (
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

//get all the satellites that are operating
func GetAllSatellitesOperatingData() (TotalSatellitesOperating []models.Satellite) {
	TotalSatellitesOperating = append(TotalSatellitesOperating,
		models.Satellite{Name: "Kenobi", XCoordinate: -500, YCoordinate: -200},
		models.Satellite{Name: "Skywalker", XCoordinate: 100, YCoordinate: -100},
		models.Satellite{Name: "Sato", XCoordinate: 500, YCoordinate: 100})

	return
}
