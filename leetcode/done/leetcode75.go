package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortColors([]int{2, 0}))
}
func sortColors(nums []int) []int {
	for cnt, i, j := 0, 0, len(nums)-1; cnt != 3; {
		if nums[i] != cnt {
			for i != j && nums[j] != cnt {
				j--
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		if i == j {
			cnt++
			j = len(nums) - 1
		} else {
			i++
		}
	}
	return nums
}
