package main

import (
	"leetcode/tree"
)

func main() {
	tree.ConstractTree([]string{})
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//要点在于小于0要即时清零
func maxPathSum(root *TreeNode) int {
	var max = 0
	m := map[*TreeNode]int{}
	//outTmp := map[*TreeNode]int{}
	var solve func(node *TreeNode) int
	solve = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Val > 0 {
			m[node] = node.Val
		} else {
			m[node] = 0
		}
		if node.Left == nil && node.Right == nil {
			return m[node]
		}
		leftVal := solve(node.Left)
		rightVal := solve(node.Right)
		if leftVal > 0 && leftVal > rightVal {
			m[node] = node.Val + leftVal
			//outTmp[node] = leftVal
		}
		if rightVal > 0 && leftVal <= rightVal {
			m[node] = node.Val + rightVal
			//outTmp[node] = rightVal
		}
		if m[node] < 0 {
			m[node] = 0
		}
		return m[node]
	}
	var procAns func(*TreeNode)
	procAns = func(node *TreeNode) {
		if node == nil {
			return
		}
		tmp := node.Val + m[node.Left] + m[node.Right]
		if tmp > max {
			max = tmp
		}
		procAns(node.Left)
		procAns(node.Right)
	}
	solve(root)
	max = m[root.Left] + m[root.Right] + root.Val
	procAns(root)
	return max
}
