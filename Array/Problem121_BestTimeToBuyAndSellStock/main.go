package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxGain := 0

	for _, currentPrice := range prices {
		if currentPrice < minPrice {
			minPrice = currentPrice
		}

		profit := currentPrice - minPrice

		if profit > maxGain {
			maxGain = profit
		}
	}

	return maxGain
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
