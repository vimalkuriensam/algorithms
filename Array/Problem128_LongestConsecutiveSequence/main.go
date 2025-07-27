package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	maxStreak := 0
	numSet := make(map[int]bool, len(nums))
	for _, num := range nums {
		numSet[num] = true
	}
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			streak := 1
			for numSet[currentNum+1] {
				currentNum++
				streak++
			}
			if streak > maxStreak {
				maxStreak = streak
			}
		}
	}
	return maxStreak
}

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{1, 0, 1, 2}))
}
