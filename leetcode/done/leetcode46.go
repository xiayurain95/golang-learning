package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1}))
}
func permute(nums []int) [][]int {
	var out [][]int
	var solve func([]int, int)
	solve = func(num []int, n int) {
		if n == len(num)-1 {
			out = append(out, num)
		} else {
			for i := 0; i < len(num)-n; i++ {
				num[len(num)-1-n], num[i] = num[i], num[len(num)-1-n]
				tmpNum := make([]int, len(num))
				copy(tmpNum, num)
				num[len(num)-1-n], num[i] = num[i], num[len(num)-1-n]
				solve(tmpNum, n+1)
			}
		}
	}
	solve(nums, 0)
	return out
}
