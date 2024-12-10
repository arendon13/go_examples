package main

import "fmt"

func main() {

	sum := 12
	var a1 = []int{1, 2, 3, 9}
	var a2 = []int{1, 2, 4, 4}

	fmt.Println(hasPairWithSumSorted(a1, sum))
	fmt.Println(hasPairWithSumSorted(a2, sum))

	fmt.Println()
	a1 = []int{1, 9, 2, 3}
	a2 = []int{4, 2, 1, 4}

	fmt.Println(hasPairWithSumUnsorted(a1, sum))
	fmt.Println(hasPairWithSumUnsorted(a2, sum))

}

func hasPairWithSumSorted(data []int, sum int) bool {

	low := 0
	high := len(data) - 1

	for low < high {

		s := data[low] + data[high]

		if s == sum {
			return true
		} else if s < sum {
			low++
		} else {
			high--
		}

	}

	return false

}

func hasPairWithSumUnsorted(data []int, sum int) bool {

	comps := map[int]bool{}

	for _, val := range data {

		if _, ok := comps[val]; ok {
			return true
		}

		comps[sum-val] = true

	}

	return false

}
