package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	p, q, r := m-1, n-1, m+n-1
	for r >= 0 {
		if p >= 0 && q >= 0 && nums1[p] > nums2[q] {
			nums1[r] = nums1[p]
			p--
		} else if q >= 0 {
			nums1[r] = nums2[q]
			q--
		}
		r--
	}
}

func main() {
	nums1 := []int{2, 0}
	nums2 := []int{1}
	merge(nums1, 1, nums2, 1)
	fmt.Println(nums1)
}
