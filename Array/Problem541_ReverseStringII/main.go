package main

import "fmt"

func reverse(a, b int, word []rune) {
	left, right := a, b

	for left < right {
		word[left], word[right] = word[right], word[left]
		left++
		right--
	}
}

func reverseStr(s string, k int) string {
	words := []rune(s)

	for i := 0; i < len(s); i += 2 * k {
		left := i
		right := i + k - 1
		if right >= len(s) {
			right = len(s) - 1
		}
		reverse(left, right, words)
	}

	return string(words)
}

func main() {
	fmt.Println(reverseStr("abcdefghijklm", 2))
}
