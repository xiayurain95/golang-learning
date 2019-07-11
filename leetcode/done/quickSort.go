package main

import (
	"fmt"
)

func main() {
	fmt.Println(fastSort([]int{45, 2, 6, 83, 2, 5, 82, 5, 73, 1}))
}

//要点:1.在i暂停时,一定是i-1==j时,交换时一定要注意i<j
//2.递归时一定不要把数轴值放进去
//3.条件判断,条件判断,条件判断,~~快排超数组真容易...
func fastSort(nums []int) []int {
	var solve func([]int)
	solve = func(num []int) {
		if len(num) <= 1 {
		} else {
			pivot := num[0]
			pivotNum := 0
			for i, j := 1, len(num)-1; i < len(num) && j >= 0; {
				for i < len(num) && num[i] <= pivot {
					i++
				}
				for j >= 0 && num[j] > pivot {
					j--
				}
				if i < len(num) && j >= 0 && i <= j {
					num[i], num[j] = num[j], num[i]
				}
				if i == j+1 {
					pivotNum = j
					break
				}
			}
			num[pivotNum], num[0] = num[0], num[pivotNum]
			solve(num[:pivotNum])
			solve(num[pivotNum+1:])
		}
	}
	solve(nums)
	return nums
}
