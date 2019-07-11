package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{1, 3}, 2))
}
func search(nums []int, target int) int {
	var findAns func(first, last int) int
	findAns = func(first, last int) int {
		if first > last {
			return -1
		}
		if nums[first] <= nums[last] {
			if !(nums[first] <= target && nums[last] >= target) {
				return -1
			}
		}
		mid := (last + first + 1) / 2
		if nums[mid] == target {
			return mid
		} else {
			leftOut := findAns(first, mid-1)
			if leftOut != -1 {
				return leftOut
			} else {
				return findAns(mid+1, last)
			}
		}
	}
	if len(nums) == 0 {
		return -1
	} else {
		return findAns(0, len(nums)-1)
	}
}
