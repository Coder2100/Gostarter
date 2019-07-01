package main

import (
	"fmt"
	"math"
)

// define interface

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectable struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectable) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectable{width: 10, height: 5}
	// print result
	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

}
