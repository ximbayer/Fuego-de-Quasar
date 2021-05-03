package functions

import (
	"fmt"
	"strings"
)

var NoSolutionLocation = fmt.Errorf("No solution for localization.")
var CoordinatesError = fmt.Errorf("The number of coordinates to analyze is incorrect. It must be one, two, or three maximum coordinates.")
var NoSolutionMessages = fmt.Errorf("The message cannot be decrypted.")

type Point struct {
	X float64
	Y float64
	R float64
}

var SatellCoordinates []Point

func ProcessData(coordinates []Point, distances []float64, messages [][]string) (xcoord float64, ycoord float64, message string, errLocation string, errMessage string) {
	SatellCoordinates = coordinates
	xcoord, ycoord, errLocation = GetLocation(distances...)
	message, errMessage = GetMessage(messages...)
	return
}

//X= 9999999999 and Y=9999999999 is an incorrect coordinate. These values are to represent an error in the Location
func GetLocation(distances ...float64) (xcoord float64, ycoord float64, err string) {
	countCoordinates := len(SatellCoordinates)
	p1 := Point{}
	p2 := Point{}
	p3 := Point{}

	switch countCoordinates {
	case 1:
		p1 = SatellCoordinates[0]
	case 2:
		p1 = SatellCoordinates[0]
		p2 = SatellCoordinates[1]
	case 3:
		p1 = SatellCoordinates[0]
		p2 = SatellCoordinates[1]
		p3 = SatellCoordinates[2]
	default:
		return 9999999999, 9999999999, CoordinatesError.Error()
	}

	ex := divide(subtract(p2, p1), normalize(subtract(p2, p1)))
	i := dot(ex, subtract(p3, p1))
	a := subtract(subtract(p3, p1), multiply(ex, i))
	ey := divide(a, normalize(a))
	d := normalize(subtract(p2, p1))
	j := dot(ey, subtract(p3, p1))

	x := (square(p1.R) - square(p2.R) + square(d)) / (2 * d)
	y := (square(p1.R)-square(p3.R)+square(i)+square(j))/(2*j) - (i/j)*x

	if x == 0 && y == 0 {
		return 9999999999, 9999999999, NoSolutionLocation.Error()
	}
	location := add(p1, add(multiply(ex, x), multiply(ey, y)))
	location = RoundUp(location, 3)

	return location.X, location.Y, ""
}

func GetMessage(messages ...[]string) (msg string, err string) {
	var lengthMin int
	var message []string
	if len(messages) < 1 {
		return "", NoSolutionMessages.Error()
	}

	lengthMin = len(messages[0])

	//Get the length of the minimum message
	for y := 0; y < len(messages); y++ {
		if lengthMin > len(messages[y]) {
			lengthMin = len(messages[y])
		}
	}

	for i := 0; i < lengthMin; i++ {
		for _, m := range messages {
			//Get the mismatch of the current message
			messageMismatch := len(m) - lengthMin
			//Verify mismatch of the current message. If it is positive, we will move the number of this mismatch in the current message to the right
			if messageMismatch > 0 {
				//Verify the word in the current message. If it is not "", we will append it to the message slice
				if m[i+messageMismatch] != "" {
					if !strings.Contains(strings.Join(message, " "), m[i+messageMismatch]) {
						message = append(message, m[i+messageMismatch])
						break
					}
				}
			} else {
				//If the mismatch of the current message is 0 (zero), we don´t need to move in the current message because it is not necessary
				if m[i] != "" {
					if !strings.Contains(strings.Join(message, " "), m[i]) {
						message = append(message, m[i])
						break
					}
				}
			}
		}
	}
	if len(message) < lengthMin {
		return "", NoSolutionMessages.Error()
	}
	msg = strings.Join(message, " ")
	return msg, ""
}
