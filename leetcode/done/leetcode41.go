package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
}
func firstMissingPositive(nums []int) int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i && (0 <= nums[i]-1 && nums[i]-1 < len(nums)) {
			tmp = nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			for 1 <= tmp && tmp <= len(nums) && nums[tmp-1] != tmp {
				tmp, nums[tmp-1] = nums[tmp-1], tmp
			}
		}
	}
	//for 1 <= tmp && tmp <= len(nums) && tmp != nums[tmp-1] {
	// tmp, nums[tmp-1] = nums[tmp-1], tmp
	// }
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums) + 1
}
