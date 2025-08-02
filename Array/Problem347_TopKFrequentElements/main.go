package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	var elements []int
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	arrSlice := make([][]int, len(hashMap)+1)
	for key, value := range hashMap {
		arrSlice[value] = append(arrSlice[value], key)
	}
	right := len(arrSlice) - 1
	index := 0
	for right >= 0 {
		elements = append(elements, arrSlice[right]...)
		index += len(arrSlice[right])
		if index >= k {
			break
		}
		right--
	}
	elements = elements[:k]
	return elements
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}
