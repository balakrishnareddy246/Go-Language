package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x,y,radius float64
}

type Rectangle struct {
	width, height float64
}

func(circle Circle) area() float64 {
	return math.Pi * circle.radius.radius * circle.radius
}

func(rect Rectangle) area() float64 {
	return rect.width * rect.height

}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x:0, y:0,raduis:5}
	fmt.Prinltn("Circle area:- %f\n", getArea(circle))
	fmt.Println("Rectangle area: %f\n", getArea(rectangle))
}