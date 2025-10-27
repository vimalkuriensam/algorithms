package main

import "fmt"

func replaceElements(arr []int) []int {
	current := len(arr) - 1
	max := -1

	for current >= 0 {
		if current == len(arr)-1 {
			max = arr[current]
			arr[current] = -1
		} else {
			tempMax := arr[current]
			arr[current] = max
			if tempMax > max {
				max = tempMax
			}
		}
		current--
	}

	return arr
}

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
	fmt.Println(replaceElements([]int{400}))
}
