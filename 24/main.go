package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func Distance(p1, p2 Point) float64 {
	x := p2.x - p1.x
	x *= x

	y := p2.y - p1.y
	y *= y
	return math.Sqrt(x + y)
}

func main() {
	fmt.Printf("Расстояние: %f\n", Distance(
		Point{x: 1, y: 2},
		Point{x: 4, y: 6},
	))
}
