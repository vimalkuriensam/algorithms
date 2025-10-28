package main

import "fmt"

func divideArray(nums []int) bool {
	mapper := make(map[int]int, len(nums))

	for _, num := range nums {
		mapper[num]++
	}

	count := 0

	for _, value := range mapper {
		if value%2 != 0 {
			return false
		}
		count += (value / 2)
	}

	return count == len(nums)/2
}

func main() {
	fmt.Println(divideArray([]int{3, 2, 3, 2, 2, 2}))
	fmt.Println(divideArray([]int{1, 2, 3, 4}))
}
