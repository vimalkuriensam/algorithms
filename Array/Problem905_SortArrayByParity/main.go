package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[right]%2 == 0 && nums[left]%2 != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
			continue
		}

		if nums[left]%2 == 0 {
			left++
		}

		if nums[right]%2 != 0 {
			right--
		}
	}

	return nums
}

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}
