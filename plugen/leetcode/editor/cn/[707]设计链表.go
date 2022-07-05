package main

//leetcode submit region begin(Prohibit modification and deletion)

type Node2 struct {
	Prev *Node
	Next *Node
	Val int
}

type MyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}


func Constructor() MyLinkedList {
	return MyLinkedList{}
}


func (l *MyLinkedList) Get(index int) int {
	if index >= l.Size {
		return -1
	}

	h, t := l.Head, l.Tail

	_ = t

	for index > 0 {
		h = h.Next
		index -= 1
	}

	return h.Val
}


func (l *MyLinkedList) AddAtHead(val int)  {
	nh := &Node{
		Val: val,
		Next: l.Head,
	}

	l.Head = nh
	if l.Tail == nil {
		l.Tail = l.Head
	}

	l.Size += 1
}


func (l *MyLinkedList) AddAtTail(val int)  {
	nt := &Node{
		Val: val,
	}

	if l.Tail == nil {
		l.Tail = nt
		l.Head = nt
	} else {
		l.Tail.Next = nt
		l.Tail = nt
	}

	l.Size += 1
}


func (l *MyLinkedList) AddAtIndex(index int, val int)  {

	if index >= l.Size {
		l.AddAtTail(val)
		return
	}

	if index <= 0 {
		l.AddAtHead(val)
		return
	}

	h := l.Head
	var last *Node
	for index > 0 {
		last = h
		h  = h.Next
		index -= 1
	}

	last.Next = &Node{
		Val: val,
		Next: h,
	}

	l.Size += 1
}


func (l *MyLinkedList) DeleteAtIndex(index int)  {

	h := l.Head
	var last *Node
	for index > 0 && h != nil {
		index -= 1
		last = h
		h = h.Next
	}

	if h == nil {
		return
	}

	if last == nil {
		l.Head = h.Next
		h.Next = nil
	} else {
		last.Next = h.Next
		h.Next = nil
	}

	l.Size -= 1

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
