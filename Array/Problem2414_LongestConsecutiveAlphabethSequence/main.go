package main

import "fmt"

func longestContinuousSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	maxStreak := 1
	left, right := 0, 1
	for right < len(s) {
		if s[right] == s[right-1]+1 {
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
	fmt.Println(longestContinuousSubstring("abacaba"))
	fmt.Println(longestContinuousSubstring("abcde"))
}
