package main

/*
import (
	"fmt"
)

func main() {
	fmt.Print(findMedianSortedArrays([]int{2}, []int{1, 3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	isEven := (l1 + l2 + 1) % 2
	num := (l1 + l2) / 2
	cnt := 0
	i, j := 0, 0
	num = num - isEven
	out1, out2 := 0, 0
	for i < l1 && j < l2 && cnt < num {
		if nums1[i] < nums2[j] {
			cnt++
			i++
		} else {
			cnt++
			j++
		}
	}
	if j == l2 && cnt <= num {
		for cnt < num {
			i++
			cnt++
		}
		if isEven == 1 {
			out1, out2 = nums1[i], nums1[i+1]
		} else {
			out1 = nums1[i]
		}
	} else if i == l1 && cnt <= num {
		for cnt < num {
			j++
			cnt++
		}
		if isEven == 1 {
			out1, out2 = nums2[j], nums2[j+1]
		} else {
			out1 = nums2[j]
		}
	} else if i < l1 && j < l2 {
		if isEven == 1 {
			if i+1 < l1 && nums1[i+1] < nums2[j] {
				out1, out2 = nums1[i], nums1[i+1]
			} else if j+1 < l2 && nums2[j+1] < nums1[i] {
				out1, out2 = nums2[j], nums2[j+1]
			} else {
				out1, out2 = nums1[i], nums2[j]
			}
		} else {
			if nums1[i] < nums2[j] {
				out1 = nums1[i]
			} else {
				out1 = nums2[j]
			}
		}

	}
	if isEven == 1 {
		return (float64(out1) + float64(out2)) / 2
	} else {
		return float64(out1)
	}
}
*/
