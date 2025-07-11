package main

import "fmt"

func sortColors(nums []int) {
	left, mid, right := 0, 0, len(nums)-1
	for mid <= right {
		switch nums[mid] {
		case 0:
			nums[left], nums[mid] = nums[mid], nums[left]
			left++
			mid++
		case 1:
			mid++
		case 2:
			nums[right], nums[mid] = nums[mid], nums[right]
			right--
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
