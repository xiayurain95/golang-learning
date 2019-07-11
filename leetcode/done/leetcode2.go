package main

/*
import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n1 := listConstructor([]int{5, 9, 5})
	n2 := listConstructor([]int{5})
	printList(n1)
	printList(n2)
	printList(addTwoNumbers(n1, n2))
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	p := l3
	leftOver := 0
	for l2 != nil && l1 != nil {
		p.Next = new(ListNode)
		p = p.Next

		p.Val = (l1.Val + l2.Val + leftOver) % 10
		leftOver = (l1.Val + l2.Val + leftOver) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		var temp *ListNode
		if l1 != nil {
			temp = l1
		} else {
			temp = l2
		}
		for temp != nil {
			p.Next = new(ListNode)
			p = p.Next
			p.Val = (temp.Val + leftOver) % 10
			leftOver = (temp.Val + leftOver) / 10
			temp = temp.Next
		}
	}
	if leftOver != 0 {
		p.Next = new(ListNode)
		p.Next.Val = 1
	}
	return l3.Next
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
*/
