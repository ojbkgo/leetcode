package didi

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := make([]*ListNode, 0)

	tmp := head
	for tmp != nil {
		l = append(l, tmp)
		tmp = tmp.Next
	}

	i := len(l) - n
	if i < 0 || i > len(l) {
		return head
	}

	if len(l) == 1 {
		return nil
	}

	if i == 0 {
		return l[1]
	}

	if i == len(l) - 1 {
		l[i - 1].Next = nil
		return l[0]
	}

	l[i - 1].Next = l[i + 1]
	return l[0]
}

func Test_removeNthFromEnd() {
	l := &ListNode{
		Val: 1, // 0
		Next: &ListNode{
			Val: 2, // 1
			//Next: &ListNode{
			//	Val: 3, // 2
			//	Next: &ListNode{
			//		Val: 4, // 3
			//		Next: &ListNode{
			//			Val: 5, // 4
			//		},
			//	},
			//},
		},
	}
	print(removeNthFromEnd(l, 2))
}