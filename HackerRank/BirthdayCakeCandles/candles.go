package main

import (
	"fmt"
)

// From https://www.hackerrank.com/challenges/birthday-cake-candles

func main() {
	var arr = []int32{3, 2, 1, 3}

	fmt.Printf("%d candle(s) will be blown out", birthdayCakeCandles(arr))
}

func birthdayCakeCandles(ar []int32) int32 {
	candles := make(map[int32]int32)
	tallest := int32(0)

	for _, e := range ar {
		tempHeight := e
		if tempHeight > tallest {
			tallest = tempHeight
		}
		candles[e]++
	}

	return candles[tallest]
}
