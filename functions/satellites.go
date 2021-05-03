package functions

import (
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

func GetSatellites(req models.Request) (satellites []string) {
	for _, satellite := range req.Satellites {
		satellites = append(satellites, satellite.NameSatell)
	}
	return satellites
}
