package main

import (
	"fmt"
	"math"
)

// Point поля x и y не экспортируются
type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{x, y}
}

func Distance(pointA Point, pointB Point) float64 {
	distX := pointA.x - pointB.x
	distY := pointA.y - pointB.y
	return math.Sqrt(float64(distX*distX + distY*distY))
}

func main() {
	pointA := NewPoint(10, 10)
	pointB := NewPoint(40, 50)

	fmt.Printf("distance between %v and %v is - %f", pointA, pointB, Distance(pointA, pointB))
}
