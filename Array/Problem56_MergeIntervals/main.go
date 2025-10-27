package main

import (
	"fmt"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	updatedIntervals := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		startI, endI := intervals[i][0], intervals[i][1]
		last := len(updatedIntervals) - 1
		if startI <= updatedIntervals[last][1] {
			startI = min(startI, updatedIntervals[last][0])
			endI = max(endI, updatedIntervals[last][1])
			updatedIntervals[last] = []int{startI, endI}
		} else {
			updatedIntervals = append(updatedIntervals, intervals[i])
		}
	}

	return updatedIntervals
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{4, 7}, {1, 4}}))
}
