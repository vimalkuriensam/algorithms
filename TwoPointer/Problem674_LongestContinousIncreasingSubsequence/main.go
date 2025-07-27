package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	maxStreak := 1
	left, right := 0, 1
	for right < len(nums) {
		if nums[right] > nums[right-1] {
			right++
			continue
		}
		currentStreak := right - left
		if currentStreak > maxStreak {
			maxStreak = currentStreak
		}
		left = right
		right++
	}
	currentStreak := right - left
	if currentStreak > maxStreak {
		maxStreak = currentStreak
	}
	return maxStreak
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))
}
