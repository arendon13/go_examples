package main

import "fmt"

// From https://www.hackerrank.com/challenges/diagonal-difference

func main() {
	var arr = [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	diff := diagonalDifference(arr)
	fmt.Println(diff)
}

// calculates diagonal sums for a matix and
// returns the absolute value of the difference
func diagonalDifference(arr [][]int32) int32 {
	matrixSize := len(arr)

	lrIncrements, rlIncrements := calculateIndexIncrements(matrixSize)

	lrSum, rlSum := calculateDiagonalSums(matrixSize, lrIncrements, rlIncrements, arr)

	diff := lrSum - rlSum
	diff = absVal(&diff)

	return diff
}

// returns two arrays that contain index inrements for
// desired diagonal values based on matrix size
func calculateIndexIncrements(size int) ([]int, []int) {
	lrIndexes := []int{}
	rlIndexes := []int{}

	rlInc := size
	lrInc := 0
	for i := 0; i < size; i++ {
		if i == 0 {
			lrIndexes = append(lrIndexes, 0)
		} else {
			lrIndexes = append(lrIndexes, lrInc)
		}
		rlIndexes = append(rlIndexes, rlInc-1)

		rlInc += size - 1
		lrInc += size + 1
	}

	return lrIndexes, rlIndexes
}

// return the sum for the two desired diagonal values
func calculateDiagonalSums(size int, lrIncrements, rlIncrements []int, arr [][]int32) (int32, int32) {
	lrSum := int32(0)
	rlSum := int32(0)
	rowReverse := size - 1
	lrCol := 0
	rlCol := 0
	for row := 0; row < size; row++ {
		lrCol = lrIncrements[row] - size*row
		rlCol = rlIncrements[rowReverse] - size*rowReverse
		lrSum += arr[row][lrCol]
		rlSum += arr[rowReverse][rlCol]

		rowReverse--
	}

	return lrSum, rlSum
}

// returns the absolute value of an int32
func absVal(num *int32) int32 {
	if *num < int32(0) {
		*num = -*num
	}

	return *num
}
