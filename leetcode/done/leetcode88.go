package main

import (
	"fmt"
)

func main() {
	fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))
}
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j, k := m+n-1, m-1, n-1
	for ; j >= 0 && k >= 0; i-- {
		if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
	}
	for ; j >= 0; i, j = i-1, j-1 {
		nums1[i] = nums1[j]
	}
	for ; k >= 0; i, k = i-1, k-1 {
		nums1[i] = nums2[k]
	}
	return nums1
}
