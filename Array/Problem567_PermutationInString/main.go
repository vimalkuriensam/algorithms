package main

func getCanonicalValue(word string) [26]int {
	wCount := [26]int{}
	for _, ch := range word {
		wCount[rune(ch)-'a']++
	}
	return wCount
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Count := getCanonicalValue(s1)
	s2Count := getCanonicalValue(s2[0:len(s1)])
	if s1Count == s2Count {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		s2Count[s2[i]-'a']++
		s2Count[s2[i-len(s1)]-'a']--
		if s1Count == s2Count {
			return true
		}
	}
	return false
}

func main() {}
