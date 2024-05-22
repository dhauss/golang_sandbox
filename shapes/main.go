package main

import (
	"fmt"
	"math"
	"reflect"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{sideLength: 2}
	t := triangle{base: 2, height: 5}
	printArea(s)
	printArea(t)

}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	className := reflect.TypeOf(s)

	fmt.Printf("The area of this %v is %v\n",
		className.Name(),
		s.getArea(),
	)
}
