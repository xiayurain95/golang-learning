func removeDuplicates(nums []int) int {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); {
		if _, ok := tmp[nums[i]]; ok {
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
		} else {
			tmp[nums[i]]++
			i++
		}

	}

	return len(tmp)
}