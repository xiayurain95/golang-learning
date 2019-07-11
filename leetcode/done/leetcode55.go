package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
func canJump(nums []int) bool {
	type stackEle struct {
		MaxStep, Sum int
	}
	if len(nums) == 1 {
		return true
	}
	var stack []stackEle
	visited := make([]bool, len(nums))
	stack = append(stack, stackEle{nums[0], 0})
	visited[0] = true
	for len(stack) != 0 {
		element := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for i := 1; i < element.MaxStep+1; i++ {
			visit := nums[element.Sum+i]
			if i+element.Sum == len(nums)-1 {
				return true
			} else {
				if visited[i+element.Sum] == false {
					visited[i+element.Sum] = true
					stack = append(stack, stackEle{visit, i + element.Sum})
				}
			}
		}
	}
	return false
}
