package main

import "fmt"

func reverse(sentence []rune, start, end int) {
	for start < end {
		sentence[start], sentence[end] = sentence[end], sentence[start]
		start++
		end--
	}
}

func reverseWords(s string) string {
	sentence := []rune(s)
	left, current := 0, 0

	length := len(sentence)
	for current <= length {
		if current == length || sentence[current] == 32 { // 32 is the rune value of a space
			reverse(sentence, left, current-1)
			left = current + 1
			current = left
		}
		current++
	}

	return string(sentence)
}

func main() {
	fmt.Println(reverseWords("Hello Vimal. How are you"))
}
