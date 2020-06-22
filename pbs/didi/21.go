package didi

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var nh *ListNode

	if l1.Val < l2.Val {
		nh = l1
		l1 = l1.Next
	} else {
		nh = l2
		l2 = l2.Next
	}
	var nhtmp *ListNode = nh
	var tmp *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp = l1
			l1 = l1.Next
		} else {
			tmp = l2
			l2 = l2.Next
		}
		tmp.Next = nil
		nhtmp.Next = tmp
		nhtmp = tmp
	}

	for l1 != nil {
		nhtmp.Next = l1
		nhtmp = l1
		l1 = l1.Next
	}

	for l2 != nil {
		nhtmp.Next = l2
		nhtmp = l2
		l2 = l2.Next
	}
	return nh
}

func Test_mergeTwoLists() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	l2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 6,
			},
		},
	}
	print(l1)
	print(l2)

	l3 := mergeTwoLists(l1, l2)
	print(l3)
}

func print(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		if l.Next != nil {
			fmt.Print("->")
		}
		l = l.Next
	}

	fmt.Println("")
}
