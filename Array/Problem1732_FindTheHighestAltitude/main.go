package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	max := 0
	sum := 0

	for _, num := range gain {
		sum += num
		if sum > max {
			max = sum
		}
	}

	return max
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
