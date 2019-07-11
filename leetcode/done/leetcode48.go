package main

import (
	"fmt"
)

func main() {
	fmt.Println(rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
func rotate(matrix [][]int) [][]int {
	//tmp := make([]int, len(matrix))
	for i := 0; i < len(matrix)/2; i++ {
		matrix[i], matrix[len(matrix)-1-i] = matrix[len(matrix)-1-i], matrix[i]

	}
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}
