package main

import "fmt"

func sumOfThree(num int64) []int64 {
	sumSlice := []int64{}
	middleNum := num / 3
	if middleNum*3 == num {
		values := []int64{middleNum - 1, middleNum, middleNum + 1}
		sumSlice = append(sumSlice, values...)
	}
	return sumSlice
}

func main() {
	fmt.Println(sumOfThree(33))
	fmt.Println(sumOfThree(4))
}
