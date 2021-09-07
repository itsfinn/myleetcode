package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTowNumbers(t *testing.T) {
	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	l3 := []int{7, 0, 8}
	assert.Equal(t, true, SliceEqualListNode(l3, addTwoNumbers(MakeListNodeFromSliceInt(l1), MakeListNodeFromSliceInt(l2))))
	l1 = []int{0}
	l2 = []int{0}
	l3 = []int{0}
	assert.Equal(t, true, SliceEqualListNode(l3, addTwoNumbers(MakeListNodeFromSliceInt(l1), MakeListNodeFromSliceInt(l2))))
	l1 = []int{9, 9, 9, 9, 9, 9, 9}
	l2 = []int{9, 9, 9, 9}
	l3 = []int{8, 9, 9, 9, 0, 0, 0, 1}
	assert.Equal(t, true, SliceEqualListNode(l3, addTwoNumbers(MakeListNodeFromSliceInt(l1), MakeListNodeFromSliceInt(l2))))
	l1 = []int{9, 9, 1}
	l2 = []int{1}
	l3 = []int{0, 0, 2}
	assert.Equal(t, true, SliceEqualListNode(l3, addTwoNumbers(MakeListNodeFromSliceInt(l1), MakeListNodeFromSliceInt(l2))))
	l1 = []int{9}
	l2 = []int{1, 9, 9, 9, 9, 9, 8, 9, 9, 9}
	l3 = []int{0, 0, 0, 0, 0, 0, 9, 9, 9, 9}
	assert.Equal(t, true, SliceEqualListNode(l3, addTwoNumbers(MakeListNodeFromSliceInt(l1), MakeListNodeFromSliceInt(l2))))
}

func TestSliceEqualListNode(t *testing.T) {
	s := []int{2, 4, 3}
	assert.Equal(t, true, SliceEqualListNode(s, MakeListNodeFromSliceInt(s)))
}

func MakeListNodeFromSliceInt(a []int) *ListNode {
	l := &ListNode{}
	l.Val = a[0]
	h := l
	for i := 1; i < len(a); i++ {
		c := &ListNode{
			Val: a[i],
		}
		l.Next = c
		l = c
	}

	return h
}

func SliceEqualListNode(a []int, l *ListNode) bool {
	for i := 0; i < len(a); i++ {

		if a[i] != l.Val {
			return false
		}
		if i != len(a)-1 {
			if l.Next == nil {
				return false
			}
			l = l.Next
		} else {
			if l.Next != nil {
				return false
			}
		}

	}
	return true
}
