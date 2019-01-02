package main

import "fmt"

// From https://www.hackerrank.com/challenges/plus-minus

func main() {
	var arr = []int32{-4, 3, -9, 0, 4, 1}

	plusMinus(arr)
}

func plusMinus(arr []int32) {
	groups := populateMapFromArray(arr)

	var pos, neg, zero = 0.0, 0.0, 0.0
	size := float64(len(arr))

	pos = groups["positive"] / size
	neg = groups["negative"] / size
	zero = groups["zero"] / size

	fmt.Printf("%.5f\n", pos)
	fmt.Printf("%.5f\n", neg)
	fmt.Printf("%.5f\n", zero)
}

func populateMapFromArray(arr []int32) map[string]float64 {
	groups := make(map[string]float64)

	for _, e := range arr {
		if e < 0 {
			groups["negative"]++
		} else if e > 0 {
			groups["positive"]++
		} else {
			groups["zero"]++
		}
	}

	return groups
}
