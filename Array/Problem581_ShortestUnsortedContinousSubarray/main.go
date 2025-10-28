package main

import "fmt"

func getMinMax(nums []int, left, right int) (int, int) {
	var min, max int

	if left < right {
		min = nums[left]
		max = nums[right]
	} else {
		min = nums[right]
		max = nums[left]
	}

	for i := left; i <= right; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}

	return min, max
}

func findUnsortedSubarray(nums []int) int {

	currentLeft, currentRight := 0, len(nums)-1

	for currentLeft < len(nums)-1 {
		if nums[currentLeft] <= nums[currentLeft+1] {
			currentLeft++
		} else {
			break
		}
	}

	for currentRight > 0 {
		if nums[currentRight-1] <= nums[currentRight] {
			currentRight--
		} else {
			break
		}
	}

	if currentRight == 0 && currentLeft == len(nums)-1 {
		return 0
	}

	min, max := getMinMax(nums, currentLeft, currentRight)
	left, right := currentLeft, currentRight

	for left > 0 && nums[left-1] > min {
		left--
	}

	for right < len(nums)-1 && nums[right+1] < max {
		right++
	}

	return right - left + 1
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray([]int{1}))
	fmt.Println(findUnsortedSubarray([]int{2, 1}))
	fmt.Println(findUnsortedSubarray([]int{1, 3, 2, 2, 2}))
	fmt.Println(findUnsortedSubarray([]int{2, 3, 3, 2, 4}))
}
