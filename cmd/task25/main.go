package main

import (
	"fmt"
	"math"
)

/*
 Разработать программу нахождения расстояния между двумя точками, которые представлены
 в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	point := NewPoint(1, 2)
	point1 := NewPoint(2, 2)

	fmt.Println(FindDistance(point, point1))
}

func FindDistance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
