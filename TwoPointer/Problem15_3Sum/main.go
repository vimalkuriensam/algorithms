package main

import (
	"fmt"
	"sort"
)

func threesums(nums []int) [][]int {
	triplets := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				triplets = append(triplets, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				if left < right && nums[left] == nums[left-1] {
					left++
				}
				if left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}

	}
	return triplets
}

func main() {
	fmt.Println(threesums([]int{-1, 0, 1, 2, -1, -4}))
}
