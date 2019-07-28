package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	length float64
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func printArea(s shape) {
	fmt.Printf("Area for %T: ", s)
	fmt.Println(s.getArea())
}

func main() {

	t := triangle{base: 3, height: 4}
	printArea(t)

	s := square{length: 2}
	printArea(s)

}
