package main

import "fmt"

func getCanonicalValue(word string) [26]int {
	wPoint := [26]int{}
	for _, ch := range word {
		wPoint[rune(ch)-'a']++
	}
	return wPoint
}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	sCount := getCanonicalValue(s[0:len(p)])
	pCount := getCanonicalValue(p)
	res := make([]int, 0, len(s))
	if sCount == pCount {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		sCount[s[i]-'a']++
		sCount[s[i-len(p)]-'a']--
		if sCount == pCount {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
