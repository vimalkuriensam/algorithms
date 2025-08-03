package main

import (
	"fmt"
	"math"
)

func longestValidParenthesesTwoPass(s string) int {
	maxLen := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case byte('('):
			left++
		case byte(')'):
			right++
		}
		if left == right {
			maxLen = int(math.Max(float64(maxLen), float64(right*2)))
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case byte('('):
			right++
		case byte(')'):
			left++
		}
		if left == right {
			maxLen = int(math.Max(float64(maxLen), float64(right*2)))
		} else if right > left {
			left, right = 0, 0
		}
	}
	return maxLen
}

func main() {
	fmt.Println(longestValidParenthesesTwoPass("(()"))
	fmt.Println(longestValidParenthesesTwoPass(")()())"))
	fmt.Println(longestValidParenthesesTwoPass(""))
}
