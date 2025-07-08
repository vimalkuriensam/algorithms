package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	sMap := make(map[rune]bool)
	left := 0
	maxLen := 0
	for right := 0; right < len(s); right++ {
		current_char := rune(s[right])
		for sMap[current_char] {
			delete(sMap, rune(s[left]))
			left++
		}
		sMap[current_char] = true
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
