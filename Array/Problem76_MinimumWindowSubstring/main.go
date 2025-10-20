package main

import "fmt"

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	wantMap := make(map[rune]int, len(t))
	wantSize := 0
	haveMap := make(map[rune]int, len(s))
	haveSize := 0
	minLength := len(s) + 1
	var res string
	left, right := 0, 0

	for _, ch := range t {
		wantMap[ch]++
	}
	wantSize = len(wantMap)

	for right < len(s) {
		r_ch := rune(s[right])
		haveMap[r_ch]++
		if wantMap[r_ch] == haveMap[r_ch] {
			haveSize++
		}

		for haveSize == wantSize {
			curr_length := right - left + 1
			if curr_length < minLength {
				minLength = curr_length
				res = s[left : right+1]
			}
			l_ch := rune(s[left])
			haveMap[l_ch]--
			if wantMap[l_ch] > 0 && haveMap[l_ch] < wantMap[l_ch] {
				haveSize--
			}
			left++
		}

		right++
	}

	return res
}

func main() {

	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
