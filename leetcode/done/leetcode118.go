package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate(5))
}
func generate(numRows int) [][]int {
	var out [][]int
	if numRows == 0 {
		return out
	}
	out = append(out, []int{1})
	if numRows == 1 {
		return out
	}
	out = append(out, []int{1, 1})
	for i := 2; i < numRows; i++ {
		func() {
			tmp := []int{1}
			for i := 0; i < len(out[len(out)-1])-1; i++ {
				tmp = append(tmp, out[len(out)-1][i]+out[len(out)-1][i+1])
			}
			tmp = append(tmp, 1)
			out = append(out, tmp)
		}()
	}
	return out
}
