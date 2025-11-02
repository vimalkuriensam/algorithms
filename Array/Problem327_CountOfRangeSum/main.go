package main

import "fmt"

func countRangeSum(nums []int, lower int, upper int) int {
	prefix := make([]int, len(nums)+1)
	prefix[0] = 0
	count := 0

	for j := 0; j < len(nums); j++ {
		prefix[j+1] = prefix[j] + nums[j]
		lowerBound := prefix[j+1] - upper
		upperBound := prefix[j+1] - lower

		for i := 0; i <= j; i++ {
			if prefix[i] <= upperBound && prefix[i] >= lowerBound {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countRangeSum([]int{-2, 5, -1}, -2, 2))
}
