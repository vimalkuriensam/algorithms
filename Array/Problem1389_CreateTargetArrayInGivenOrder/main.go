package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	isInserted := make(map[int]struct{})
	current := 0

	for current < len(nums) {
		targetIndex := index[current]
		if _, found := isInserted[targetIndex]; found {
			shiftIndex := len(nums) - 1
			for shiftIndex > targetIndex {
				target[shiftIndex] = target[shiftIndex-1]
				if _, found := isInserted[shiftIndex-1]; found {
					isInserted[shiftIndex] = struct{}{}
				}
				shiftIndex--
			}
			target[targetIndex] = nums[current]
		} else {
			target[targetIndex] = nums[current]
			isInserted[targetIndex] = struct{}{}
		}
		current++
	}

	return target
}

func main() {
	fmt.Println(createTargetArray([]int{4, 2, 4, 3, 2}, []int{0, 0, 1, 3, 1}))
}
