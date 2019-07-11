package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 2))
}
func uniquePaths(m int, n int) int {
	posMap := map[position]int{{1, 1}: 1, {1, 2}: 1, {2, 1}: 1, {2, 2}: 2}
	var solve func(currentM, currentN int) int
	solve = func(currentM, currentN int) int {
		if num, ok := posMap[position{currentM, currentN}]; ok {
			return num
		} else if currentM-1 > 0 && currentN-1 > 0 {
			num := solve(currentM-1, currentN) + solve(currentM, currentN-1)
			posMap[position{currentM, currentN}] = num
			return num
		} else if currentM-1 <= 0 && currentN-1 > 0 {
			num := solve(currentM, currentN-1)
			posMap[position{currentM, currentN}] = num
			return num
		} else if currentM-1 > 0 && currentN-1 <= 0 {
			num := solve(currentM-1, currentN)
			posMap[position{currentM, currentN}] = num
			return num
		} else {
			return 0
		}
	}
	return solve(m, n)
}

type position struct {
	m, n int
}
