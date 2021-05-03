package functions

import (
	"math"
)

func square(v float64) float64 {
	return v * v
}

func normalize(p Point) float64 {
	return math.Sqrt(square(p.X) + square(p.Y))
}

func dot(p1, p2 Point) float64 {
	return p1.X*p2.X + p1.Y*p2.Y
}

func subtract(p1, p2 Point) Point {
	return Point{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
	}
}

func add(p1, p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func divide(p Point, v float64) Point {
	return Point{
		X: p.X / v,
		Y: p.Y / v,
	}
}

func multiply(p Point, v float64) Point {
	return Point{
		X: p.X * v,
		Y: p.Y * v,
	}
}

func RoundUp(input Point, places int) Point {
	pow := math.Pow(10, float64(places))
	input.X = pow * input.X
	input.Y = pow * input.Y
	input.X = (math.Ceil(input.X)) / pow
	input.Y = (math.Ceil(input.Y)) / pow
	return input
}
