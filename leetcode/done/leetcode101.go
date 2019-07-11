package main

func main() {

}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(nodeLeft, nodeRight *TreeNode) bool {
	if (nodeLeft != nil && nodeRight == nil) || (nodeLeft == nil && nodeRight != nil) {
		return false
	} else if nodeLeft == nil && nodeRight == nil {
		return true
	} else if nodeLeft.Val == nodeRight.Val {
		return isMirror(nodeLeft.Left, nodeRight.Right) && isMirror(nodeLeft.Right, nodeRight.Left)
	} else {
		return false
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
