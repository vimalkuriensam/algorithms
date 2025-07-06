package main

import "fmt"

// Brute force approach makes it not the optimal solution.
// gives a time complexity of O(n2)
// as input increases to 10^4, it becomes slow
// func twoSum(nums []int, target int) []int {
// 	var indices []int
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return indices
// }

// Here time complexity and space complexity will be O(n)
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index1, num := range nums {
		complement := target - num
		if index2, found := numMap[complement]; found {
			return []int{index1, index2}
		}
		numMap[num] = index1
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
