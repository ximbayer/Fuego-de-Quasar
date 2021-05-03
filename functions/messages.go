package functions

import (
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

//get the messages received on each satellite
func GetMessages(shipToSatellites []models.ShipToSatellite) (messages [][]string) {
	for _, satellite := range shipToSatellites {
		messages = append(messages, satellite.Message)
	}
	return messages
}
