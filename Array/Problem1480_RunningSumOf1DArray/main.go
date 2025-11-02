package main

import "fmt"

func runningSum(nums []int) []int {
	prefix := make([]int, len(nums))

	for index, num := range nums {
		if index == 0 {
			prefix[0] = num
		} else {
			prefix[index] = prefix[index-1] + num
		}
	}

	return prefix
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}
