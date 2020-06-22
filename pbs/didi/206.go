package didi

import "fmt"

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head
	head = head.Next
	tmp.Next = nil
	for head != nil {
		tmp2 := head.Next
		head.Next = tmp
		tmp = head
		head = tmp2
	}
	return tmp
}


func Test_reverseList() {
	l := &ListNode{
		Val:  1,
		Next: &ListNode{
			Val: 2,
			//Next: &ListNode{
			//	Val: 3,
			//	Next: &ListNode{
			//		Val: 4,
			//	},
			//},
		},
	}

	printList(reverseList(l))

}

func printList(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}