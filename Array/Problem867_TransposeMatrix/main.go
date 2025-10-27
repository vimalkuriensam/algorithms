package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	newMatrix := make([][]int, len(matrix[0]))
	for i := 0; i < len(newMatrix); i++ {
		newMatrix[i] = make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}
	return newMatrix
}

func main() {
	fmt.Println(transpose([][]int{{2, 4, -1}, {-10, 5, 11}, {18, -7, 6}}))
}
