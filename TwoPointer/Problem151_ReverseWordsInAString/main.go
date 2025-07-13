package main

import (
	"fmt"
)

func removeSpaces(s string) []byte {
	b := []byte(s)
	read, write := 0, 0
	for read < len(b) {
		if (write > 0 && b[read] == 32 && b[read-1] != 32) || b[read] != 32 {
			b[write] = b[read]
			write++
		}
		read++
	}
	if write > 0 && b[write-1] == 32 {
		return b[:write-1]
	}
	return b[:write]
}

func reverseString(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func reverseWords(s string) string {
	trimmedBytes := removeSpaces(s)
	reverseString(trimmedBytes, 0, len(trimmedBytes)-1)
	left, right := 0, 0
	for right < len(trimmedBytes) {
		if right == len(trimmedBytes)-1 {
			reverseString(trimmedBytes, left, right)
			left = right + 1
		} else if trimmedBytes[right] == 32 {
			reverseString(trimmedBytes, left, right-1)
			left = right + 1
		}
		right++
	}
	return string(trimmedBytes)
}

func main() {
	sentence := "  hello world  "
	fmt.Println(reverseWords(sentence))
}
