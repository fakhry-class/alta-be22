package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Luas() float64 {
	return r.width * r.height
}

func (c Circle) Luas() float64 {
	return math.Pi * c.radius * c.radius
}

func KelilingPersegi() {

}

func KelilingLingkaran() {

}

func main() {
	rect := Rect{5.0, 4.0}
	cir := Circle{5.0}
	fmt.Printf("Area of rectangle rect = %0.3f\n", rect.Luas())
	fmt.Printf("Area of circle cir = %0.3f\n", cir.Luas())
}
