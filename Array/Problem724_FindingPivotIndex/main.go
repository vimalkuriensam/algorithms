package main

import "fmt"

func pivotIndex(nums []int) int {
	var totalSum, currentSum int
	for _, num := range nums {
		totalSum += num
	}

	for index, num := range nums {
		currentSum += num
		if currentSum == totalSum {
			return index
		}
		totalSum -= num
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}
