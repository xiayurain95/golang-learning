package main

func main() {
	sortedArrayToBST([]int{10. - 3, 0, 5, 9})
}
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	node := new(TreeNode)
	node.Val = nums[len(nums)/2]
	node.Left = sortedArrayToBST(nums[:len(nums)/2])
	node.Right = sortedArrayToBST(nums[len(nums)/2+1 : len(nums)])
	return node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
