package main

import "fmt"

func main() {

	j := "aA"
	s := "aAAbbbb"

	fmt.Println(numJewelsInStones(j, s))

}

func numJewelsInStones(J, S string) int {

	jCount := 0
	jewels := map[rune]bool{}

	for _, c := range J {
		jewels[c] = true
	}

	for _, c := range S {
		if _, ok := jewels[c]; ok {
			jCount++
		}
	}

	return jCount

}
