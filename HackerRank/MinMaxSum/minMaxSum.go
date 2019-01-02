package main

import "fmt"

// From https://www.hackerrank.com/challenges/mini-max-sum

func main() {
	var arr1 = []int32{1, 2, 3, 4, 5}
	var arr2 = []int32{7, 69, 2, 221, 8974}
	var arr3 = []int32{256741038, 623958417, 467905213, 714532089, 938071625}

	miniMaxSum(arr1)
	miniMaxSum(arr2)
	miniMaxSum(arr3)
}

func miniMaxSum(arr []int32) {
	arrLength := len(arr)
	mini := int64(0)
	max := int64(0)

	for i := 0; i < arrLength; i++ {
		tempSum := int64(0)
		tempSum += int64(arr[(arrLength+i)%arrLength])
		tempSum += int64(arr[(arrLength+i+1)%arrLength])
		tempSum += int64(arr[(arrLength+i+2)%arrLength])
		tempSum += int64(arr[(arrLength+i+3)%arrLength])

		if i == 0 {
			mini = tempSum
			max = tempSum
			continue
		}

		if tempSum < mini {
			mini = tempSum
		}

		if tempSum > max {
			max = tempSum
		}

	}

	fmt.Printf("%d %d\n", mini, max)
}
