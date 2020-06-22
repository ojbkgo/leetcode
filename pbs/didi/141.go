package didi

func hasCycle(head *ListNode) bool {
	c := 8029
	for ; c > 0; c -= 1 {
		head = head.Next
		if head == nil {
			return true
		}
	}

	return false
}
