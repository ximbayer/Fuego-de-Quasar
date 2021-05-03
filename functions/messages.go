package functions

import (
	"github.com/ximbayer/Fuego-de-Quasar/models"
)

//get the messages received on each satellite
func GetMessages(req models.Request) (messages [][]string) {
	for _, satellite := range req.Satellites {
		messages = append(messages, satellite.Message)
	}
	return messages
}
