package main

import "fmt"

func canonicalValue(word string) [26]int {
	wPoint := [26]int{}
	for _, ch := range word {
		wPoint[rune(ch)-'a']++
	}
	return wPoint
}

func groupAnagrams(strs []string) [][]string {
	mapper := make(map[[26]int][]string)
	result := make([][]string, 0, len(strs))
	for _, str := range strs {
		cValue := canonicalValue(str)
		mapper[cValue] = append(mapper[cValue], str)
	}
	for _, value := range mapper {
		result = append(result, value)
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
