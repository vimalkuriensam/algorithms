package main

import "fmt"

func isValid(s string) bool {
	parenthesisStack := make([]byte, 0, len(s))
	stackHash := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case byte('('), byte('['), byte('{'):
			parenthesisStack = append(parenthesisStack, s[i])
		case byte(')'), byte(']'), byte('}'):
			stackLength := len(parenthesisStack)
			if stackLength == 0 {
				return false
			}
			popStack := parenthesisStack[stackLength-1]
			parenthesisStack = parenthesisStack[:stackLength-1]
			if popStack != stackHash[s[i]] {
				return false
			}
		}
	}
	return len(parenthesisStack) == 0
}

func main() {
	fmt.Println(isValid("((()))"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("([)]"))
}
