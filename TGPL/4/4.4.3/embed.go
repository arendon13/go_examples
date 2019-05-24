package main

import (
	"fmt"
	"strings"
)

type point struct {
	X, Y int
}

type circle struct {
	point
	Radius int
	X      int
}

type wheel struct {
	circle
	Spokes int
}

func (w wheel) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Spokes: %d\n", w.Spokes))
	builder.WriteString(fmt.Sprintf("Radius: %d\n", w.circle.Radius))
	builder.WriteString(fmt.Sprintf("Circle X: %d\n", w.circle.X))
	builder.WriteString(fmt.Sprintf("X: %d\n", w.circle.point.X))
	builder.WriteString(fmt.Sprintf("Y: %d\n", w.circle.point.Y))

	return builder.String()
}

func main() {

	var w wheel

	// w.circle.point.X = 5
	// fmt.Println(w.X)
	// fmt.Println(w)

	// w.point.X = 10
	// fmt.Println(w.X)
	// fmt.Println(w)

	w.X = 15
	fmt.Println(w.X)
	fmt.Println(w)

}
