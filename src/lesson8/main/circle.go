package main

import (
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) perim() float64  {
	return 2 * math.Pi * c.Radius
}