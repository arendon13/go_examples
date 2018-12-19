package main

import (
	"fmt"
)

type celcius float64

type fahrenheit float64

func (c celcius) fahrenheit() fahrenheit {
	conversion := (c * 9 / 5) + 32
	return fahrenheit(conversion)
}

func (f fahrenheit) celcius() celcius {
	conversion := (f - 32) * 5 / 9
	return celcius(conversion)
}

func main() {
	celciusToFahrenheitTable()
	fmt.Println()
	fmt.Println()
	fahrenheitToCelciusTable()
}

func printBar() {
	fmt.Printf("%23v", "=======================\n")
}

func printHeader(col1, col2 string) {

	var degree byte = 248
	degreeSymbol := string(degree)
	celciusLabel := degreeSymbol + col1
	fahrenheitLabel := degreeSymbol + col2

	printBar()
	fmt.Printf("| %-8v | %-8v |\n", celciusLabel, fahrenheitLabel)
	printBar()
}

func celciusToFahrenheitTable() {

	printHeader("C", "F")

	var temp celcius

	for i := -40.0; i <= 100.0; i += 5.0 {

		temp = celcius(i)

		fmt.Printf("| %-8.1f | %-8.1f |\n", i, temp.fahrenheit())

		printBar()

	}

}

func fahrenheitToCelciusTable() {

	printHeader("F", "C")

	var temp fahrenheit

	for i := -40.0; i <= 100.0; i += 5.0 {

		temp = fahrenheit(i)

		fmt.Printf("| %-8.1f | %-8.1f |\n", i, temp.celcius())

		printBar()

	}

}
