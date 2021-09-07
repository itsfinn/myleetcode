package addtwonumbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(l *ListNode) {
	fmt.Print("[")
	for l.Next != nil {
		fmt.Print(l.Val, ", ")
		l = l.Next
	}

	fmt.Print(l.Val)
	fmt.Println("]")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l := &ListNode{}
	head := l

	var jinwei int
	if l1.Val+l2.Val <= 9 {
		l.Val = l1.Val + l2.Val
		jinwei = 0
	} else {
		l.Val = l1.Val + l2.Val - 10
		jinwei = 1
	}
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil && l2 != nil {
		c := &ListNode{}
		if l1.Val+l2.Val+jinwei <= 9 {
			c.Val = l1.Val + l2.Val + jinwei
			jinwei = 0
		} else {
			c.Val = l1.Val + l2.Val + jinwei - 10
			jinwei = 1
		}
		l.Next = c
		l = c
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		if jinwei == 0 {
			l.Next = l1
			break
		}
		c := &ListNode{}
		if l1.Val+jinwei <= 9 {
			c.Val = l1.Val + jinwei
			jinwei = 0
		} else {
			c.Val = l1.Val + jinwei - 10
			jinwei = 1
		}
		l.Next = c
		l = c
		l1 = l1.Next

	}
	for l2 != nil {
		if jinwei == 0 {
			l.Next = l2
			break
		}
		c := &ListNode{}
		if l2.Val+jinwei <= 9 {
			c.Val = l2.Val + jinwei
			jinwei = 0
		} else {
			c.Val = l2.Val + jinwei - 10
			jinwei = 1
		}
		l.Next = c
		l = c
		l2 = l2.Next

	}
	if jinwei == 1 {
		c := &ListNode{Val: 1}
		l.Next = c
	}

	return head
}
