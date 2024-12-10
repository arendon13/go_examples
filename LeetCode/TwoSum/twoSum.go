package main

import "fmt"

func main() {

	sum := 9
	var a1 = []int{15, 7, 11, 2}

	fmt.Println(twoSum(a1, sum))

}

func twoSum(nums []int, target int) [2]int {

	comps := map[int]int{}
	var indexSums [2]int

	for i := 0; i < len(nums); i++ {

		val := nums[i]

		if _, ok := comps[val]; ok {
			indexSums[0] = comps[val]
			indexSums[1] = i
		}

		comps[target-val] = i

	}

	return indexSums

}
