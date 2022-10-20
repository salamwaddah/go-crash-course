package main

import (
	"fmt"
	"math"
)

// Shape define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	w, h float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.w * r.h
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c := Circle{x: 0, y: 0, radius: 5}
	r := Rectangle{w: 10, h: 5}

	fmt.Printf("Circle area %f\n", getArea(c))
	fmt.Printf("Rectangle area %f\n", getArea(r))
}
