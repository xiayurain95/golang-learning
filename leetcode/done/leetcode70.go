package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(3))
}
func climbStairs(n int) int {
	//out := 0
	m := map[int]int{1: 1, 2: 2}
	var solve func(int) int
	solve = func(n int) int {
		if num, ok := m[n]; ok {
			return num
		} else {
			m[n] = solve(n-1) + solve(n-2)
			return m[n]
		}
	}
	return solve(n)
}
