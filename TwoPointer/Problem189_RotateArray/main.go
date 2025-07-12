package main

import "fmt"

func reverseNums(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func rotate(nums []int, k int) {
	left, right := 0, len(nums)-1
	if k > len(nums) {
		k = k % len(nums)
	}
	reverseNums(nums, left, right)
	left, right = 0, k-1
	reverseNums(nums, left, right)
	left, right = k, len(nums)-1
	reverseNums(nums, left, right)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)
}
