package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
func subsets(nums []int) [][]int {
	var out [][]int
	var solve func([]int, int)
	solve = func(output []int, pos int) {
		out = append(out, output)
		if pos < len(nums) {
			for i := pos; i < len(nums); i++ {
				outTmp := make([]int, len(output))
				copy(outTmp, output)
				outTmp = append(outTmp, nums[i])
				solve(outTmp, i+1)
			}
		}
	}
	solve([]int{}, 0)
	return out
}
