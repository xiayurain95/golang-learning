package main

func main() {

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var solve func(*TreeNode) int
	solve = func(node *TreeNode) int {
		if node != nil {
			left := 1 + solve(node.Left)
			Right := 1 + solve(node.Right)
			if left > Right {
				return left
			} else {
				return Right
			}
		} else {
			return 0
		}
	}
	return solve(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
