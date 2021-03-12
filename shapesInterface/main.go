package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	s := square{6}
	t := triangle{2, 4}

	fmt.Println("Square area: ", printArea(s))
	fmt.Println("Triangle area: ", printArea(t))
}

func printArea(s shape) float64 {
	return s.getArea()
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
