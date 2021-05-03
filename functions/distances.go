package functions

import (
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

func GetDistances(req models.Request) (distances []float64) {
	for _, satellite := range req.Satellites {
		distances = append(distances, satellite.Distance)
	}
	return distances
}
