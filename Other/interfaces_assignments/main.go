package main

import "math"
import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	tri := triangle{
		height: 12.5,
		base:   25,
	}
	sq := square{
		sideLength: 9.8,
	}

	printArea(tri)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Printf("%.2f", s.getArea())
	fmt.Println()
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}
