package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 2}))
}
func maxArea(height []int) int {
	out := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			if height[i] >= height[j] {
				if out < height[j]*(j-i) {
					out = height[j] * (j - i)
				}
			} else {
				if out < height[i]*(j-i) {
					out = height[i] * (j - i)
				}
			}
		}
	}
	return out
}
