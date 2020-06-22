package didi


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	ll1 := l1
	ll2 := l2

	up := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + up
		if val >= 10 {
			up = 1
			val -= 10
		} else {
			up = 0
		}

		l1.Val = val
		l2.Val = val

		if up > 0 && l1.Next == nil && l2.Next == nil {
			l1.Next = &ListNode{Val: 1}
			return ll1
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		increase(l1, up)
		return ll1
	}

	if l2 != nil {
		increase(l2, up)
		return ll2
	}



	return ll1
}

func increase(l *ListNode, up int) {
	last := l
	for l != nil {
		if l.Val + up > 9 {
			l.Val = 0
			up = 1
		} else {
			up = 0
		}
		last = l
		l = l.Next
	}

	if up > 0 {
		last.Next = &ListNode{Val: 1}
	}
}