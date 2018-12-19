package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var num = rand.Intn(8) + 3
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	// a random distance to Mars (km)
	var distance = rand.Intn(345000001) + 56000000
	fmt.Println()
	fmt.Println(distance)

}
