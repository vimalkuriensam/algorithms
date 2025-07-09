package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	re := regexp.MustCompile(`[^a-z0-9]`)
	s = re.ReplaceAllString(s, "")
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return true
	}
	isValid := false
	left, right := 0, len(s)-1
	if left == right {
		return true
	}
	for left < right {
		if s[left] != s[right] {
			isValid = false
			break
		}
		isValid = true
		left++
		right--
	}
	return isValid
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
	fmt.Println(isPalindrome("a"))
	fmt.Println(isPalindrome("aa"))
}
