package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	mapper := make(map[int]int) // store the num and the position where it exist
	for index, num := range nums {
		if prevPosition, found := mapper[num]; found {
			if index-prevPosition <= k {
				return true
			}
		}
		mapper[num] = index
	}
	return false
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
