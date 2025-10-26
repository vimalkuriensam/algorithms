package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}

	for _, interval := range intervals {
		newStart, newEnd := newInterval[0], newInterval[1]
		start, end := interval[0], interval[1]

		if end < newStart {
			result = append(result, interval)
		} else if start > newEnd {
			result = append(result, newInterval)
			newInterval = interval
		} else {
			newInterval[0] = min(newStart, interval[0])
			newInterval[1] = max(newEnd, interval[1])
		}
	}

	result = append(result, newInterval)

	return result
}

func main() {

	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3}))
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
}
