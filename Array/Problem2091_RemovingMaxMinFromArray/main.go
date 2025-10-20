package main

import (
	"math"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimumDeletions(nums []int) int {
	min_value, min_value_index := math.MaxInt64, -1
	max_value, max_value_index := math.MinInt64, -1

	for index, num := range nums {
		if num < min_value {
			min_value = num
			min_value_index = index
		}
		if num > max_value {
			max_value = num
			max_value_index = index
		}
	}

	leftSideDeletionValue := getMax(min_value_index, max_value_index) + 1
	rightSideDeletionValue := len(nums) - getMin(min_value_index, max_value_index)
	eitherSideDeletion := len(nums) + getMin(min_value_index, max_value_index) - getMax(min_value_index, max_value_index) + 1
	return getMin(leftSideDeletionValue, getMin(rightSideDeletionValue, eitherSideDeletion))
}

func main() {
	minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6})
}
