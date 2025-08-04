package main

import (
	"fmt"
)

func getMinValue(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

func getMaxValue(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func maxArea(height []int) int {
	area := 0
	left, right := 0, len(height)-1
	for left < right {
		currentArea := getMinValue(height[left], height[right]) * (right - left)
		area = getMaxValue(area, currentArea)
		if height[left] < height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		} else {
			left++
			right--
		}
	}
	return area
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
