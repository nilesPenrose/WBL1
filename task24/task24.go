package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	X, Y float64 // если под инкапсуляцией подразумевается сокрытие реализации, то X и Y должны быть строчными
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func (p *Point) Distance(rhs *Point) float64 {
	return math.Sqrt(e2(rhs.X-p.X) + e2(rhs.Y-p.Y))
}

// Работает более эффективно, чем math.Pow
func e2(v float64) float64 {
	return v * v
}

func main() {
	tests := []struct {
		p1 *Point
		p2 *Point
	}{
		{
			p1: NewPoint(0, 0),
			p2: NewPoint(1, 0),
		},
		{
			p1: NewPoint(1, 0),
			p2: NewPoint(0, 1),
		},
		{
			p1: NewPoint(0, 1),
			p2: NewPoint(1, 0),
		},
		{
			p1: NewPoint(1.23456, 6.54321),
			p2: NewPoint(0, 0),
		},
	}

	for _, t := range tests {
		fmt.Printf("distance %v, %v = %f\n", t.p1, t.p2, t.p1.Distance(t.p2))
	}
}
