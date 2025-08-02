package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	indices := make([]int, 0, 2)
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			indices = append(indices, []int{left + 1, right + 1}...)
			break
		}
	}
	return indices
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
