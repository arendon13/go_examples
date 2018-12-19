// How long does it take to get to Mars?
package main

import (
	"fmt"
)

func main() {
	const (
		lightspeed  = 299792
		hoursPerDay = 24
	)
	var distance, speed = 56000000, 100800

	fmt.Println(distance/lightspeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightspeed, "seconds")

	distance = 96300000
	fmt.Println()
	fmt.Println(distance/speed/hoursPerDay, "days")
}
