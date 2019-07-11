package tree

import (
	"fmt"
	"strconv"
)

func ConstractTree(nums []string) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, 0)
	for i := 0; i < len(nums); i++ {
		tmp := new(TreeNode)
		if nums[i] == "nil" {
			tmp = nil
		} else {
			if v, err := strconv.Atoi(nums[i]); err != nil {
				fmt.Println("error: %e/n", err)
				return nil
			} else {
				tmp.Val = v
			}
		}
		nodes = append(nodes, tmp)
	}
	for i := 0; i < len(nodes); i++ {
		if nodes[i] != nil {
			if i*2+2 < len(nodes) {
				nodes[i].Left = nodes[i*2+1]
				nodes[i].Right = nodes[i*2+2]
			}
		}
	}
	return nodes[0]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
