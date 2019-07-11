package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func main() {
	printList(
		removeNthFromEnd(
			listConstructor([]int{1, 2, 3, 4, 5}), 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := make(map[int]*ListNode)
	cnt := 0
	for head.Next != nil {
		tmp[cnt] = head
		cnt++
		head = head.Next
	}
	tmp[cnt] = head
	//cnt++ 一定要小心点啊...

	numCnt := cnt + 1
	targetL := numCnt - n
	if targetL == 0 {
		return tmp[1]
	} else if targetL == cnt {
		tmp[cnt-1].Next = nil
		return tmp[0]
	} else {
		tmp[targetL-1].Next = tmp[targetL+1]
		tmp[targetL].Next = nil
		return tmp[0]
	}
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
