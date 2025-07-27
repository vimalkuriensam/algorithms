package main

import "fmt"

func maxPower(s string) int {
	if len(s) == 1 {
		return 1
	}
	maxStreak := 0
	left, right := 0, 1
	for right < len(s) {
		if s[right] == s[right-1] {
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
	fmt.Println(maxPower("leetcode"))
	fmt.Println(maxPower("abbcccddddeeeeedcba"))
}
