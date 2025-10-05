package main

func firstUniqChar(s string) int {
	counter := [26]int{}
	for _, char := range s {
		counter[rune(char)-97]++
	}
	for index, char := range s {
		freqValue := counter[rune(char)-97]
		if freqValue == 1 {
			return index
		}
	}
	return -1
}

func main() {}
