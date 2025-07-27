package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	maxStreak := 0
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 1 {
			right++
			continue
		}
		currentStreak := right - left
		if currentStreak > maxStreak {
			maxStreak = currentStreak
		}
		right++
		left = right
	}
	currentStreak := right - left
	if currentStreak > maxStreak {
		maxStreak = currentStreak
	}
	return maxStreak
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
