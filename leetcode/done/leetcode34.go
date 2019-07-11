package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 10))
}
func searchRange(nums []int, target int) []int {
	/*	outLower, outHigher := -1, -1
		var solve func(first, last int)
		solve = func(first, last int) {
			if first > last {
				return
			}

			if nums[first] > target || target > nums[last] {
				return
			} else if nums[first] == target && target == nums[last] {
				if first < outLower {
					outLower = first
				}
				if last > outHigher {
					outHigher = last
				}
				return

			} else if nums[first] < target && target < nums[last] {
				mid := (first + last + 1) / 2
				solve(first, mid-1)
				solve(mid, last)
			} else {
				mid := (first + last + 1) / 2
				if nums[first] == target {
					if outLower == -1 || first < outLower {
						outLower = first
					}
					if outHigher == -1 || first > outHigher {
						outHigher = first
					}
					if nums[mid] == target {
						solve(mid+1, last)
					}
				}
				if nums[last] == target {
					if outHigher == -1 || last > outHigher {
						outHigher = first
					}
					if outHigher == -1 || last < outLower {
						outLower = last
					}
					if nums[mid] == target {
						solve(first, mid-1)
					}
				}
			}
		}
		if len(nums) != 0 {
			solve(0, len(nums)-1)
		}
		return []int{outLower, outHigher}
	*/
	m := make(map[int]int)
	var solve func(first, last int)
	solve = func(first, last int) {
		if first > last {
			return
		}
		mid := (first + last + 1) / 2
		if nums[mid] == target {
			m[mid] = 0
			solve(first, mid-1)
			solve(mid+1, last)
		} else if nums[mid] > target {
			solve(first, mid-1)
		} else {
			solve(mid+1, last)
		}
	}
	if len(nums) != 0 {
		solve(0, len(nums)-1)
	}
	outHigher, outLower := -1, -1
	for k := range m {
		if outHigher == -1 || outHigher < k {
			outHigher = k
		}
		if outLower == -1 || outLower > k {
			outLower = k
		}
	}
	return []int{outLower, outHigher}
}
