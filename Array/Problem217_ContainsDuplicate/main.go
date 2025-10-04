package main

func containsDuplicate(nums []int) bool {
	mapper := make(map[int]bool)
	for _, num := range nums {
		if _, found := mapper[num]; found {
			return true
		}
		mapper[num] = true
	}
	return false
}

func main() {}
