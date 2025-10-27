package main

import "fmt"

func mergeTwoSortedArrays(nums1, nums2 []int) []int {
	updatedNums := make([]int, len(nums1)+len(nums2))

	p, q, r := len(nums1)-1, len(nums2)-1, len(updatedNums)-1

	for r >= 0 {
		if p < 0 {
			updatedNums[r] = nums2[q]
			q--
		} else if q < 0 {
			updatedNums[r] = nums1[p]
			p--
		} else if nums1[p] > nums2[q] {
			updatedNums[r] = nums1[p]
			p--
		} else {
			updatedNums[r] = nums2[q]
			q--
		}
		r--
	}

	return updatedNums
}

func sortedSquares(nums []int) []int {
	posIndex := -1 // all the elements in the array could be negetive
	for index, num := range nums {
		if num >= 0 {
			posIndex = index
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	if posIndex == -1 {
		left, right := 0, len(nums)-1
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
		return nums
	} else if posIndex > 0 {
		nums2 := make([]int, posIndex)
		for i := posIndex - 1; i >= 0; i-- {
			nums2[posIndex-i-1] = nums[i]
		}
		nums = mergeTwoSortedArrays(nums[posIndex:], nums2)
	}

	return nums
}

func main() {

	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{0}))
	fmt.Println(sortedSquares([]int{-8, -6, -4, -1}))
}
