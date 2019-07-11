package main

import "fmt"

func main() {
	printList(swapPairs(listConstructor([]int{})))
}
func swapPairs(head *ListNode) *ListNode {
	tmp := make([]*ListNode, 0)
	var tmpNode *ListNode
	tmpNode = head
	for tmpNode != nil {
		tmp = append(tmp, tmpNode)
		tmpNode = tmpNode.Next
	}
	i, j := 0, 0
	if len(tmp)%2 == 0 {
		i = len(tmp) - 1
		j = len(tmp) - 2
		tmpNode = nil
	} else {
		i = len(tmp) - 2
		j = len(tmp) - 3
		tmpNode = tmp[len(tmp)-1]
	}
	for ; i >= 0 && j >= 0; i, j = i-2, j-2 {
		tmp[j].Next = tmpNode
		tmp[i].Next = tmp[j]
		tmpNode = tmp[i]
	}
	return tmpNode
}
func listConstructor(n1 []int) *ListNode {
	l := new(ListNode)
	p := l
	for i := 0; i < len(n1); i++ {
		p.Val = n1[i]
		if i+1 != len(n1) {
			p.Next = new(ListNode)
			p = p.Next
		}
	}
	return l
}
func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
