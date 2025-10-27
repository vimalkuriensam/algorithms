package main

import "fmt"

func moveZeroes(nums []int) {
	insertPos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[insertPos] = nums[i]
			insertPos++
		}
	}

	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
