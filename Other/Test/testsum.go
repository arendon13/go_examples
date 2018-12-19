package main

import "fmt"

func main() {
	// nums := []int{5, 9, 2, 8}
	// nums := []int{9, 4, 7, 3}
	nums := []int{3, 8, 6, 5}

	for i := range nums {
		fmt.Println("Current Number:", nums[i])
		for j := range nums {
			fmt.Println("Iterable Number:", nums[i], "/", nums[j])
			// fmt.Println("Iterable Number:", nums[i]%nums[j])

			// if i != j && (nums[i]%nums[j]) == 0 {
			// 	fmt.Println("Result:", nums[i]/nums[j])
			// }
		}
	}
}
