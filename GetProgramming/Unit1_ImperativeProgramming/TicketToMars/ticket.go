package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("%-16v %4v %-10v %4v\n", "Spaceline", "Days", "Trip-Type", "Price")
	fmt.Println("======================================")

	for ticketCount := 0; ticketCount < 10; ticketCount++ {

		var spaceline string
		const distanceToMars, hoursPerDay = 62100000, 24
		randSpaceline := rand.Intn(3) + 1

		switch randSpaceline {
		case 1:
			spaceline = "Space Adventures"
		case 2:
			spaceline = "SpaceX"
		case 3:
			spaceline = "Virgin Galactic"
		}

		var shipSpeed = rand.Intn(31-16) + 16
		var days = distanceToMars / (shipSpeed * 3600) / hoursPerDay

		var tripType string
		randTripType := rand.Intn(2) + 1

		switch randTripType {
		case 1:
			tripType = "One-way"
		case 2:
			tripType = "Round-trip"
		}

		var price = shipSpeed + 20
		if tripType == "Round-trip" {
			price *= 2
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline, days, tripType, price)

	}

}
