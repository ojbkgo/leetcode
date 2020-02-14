package interview

import (
	. "github.com/ojbkgo/leetcode-common"
	"github.com/ojbkgo/leetcode/data/interview/i06"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//head
//1	->	2	->	3	->	4
//		next
//				tmpnext
//newhead

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	newhead := head
	next := newhead.Next
	var tmpnext *ListNode
	if next != nil {
		tmpnext = next.Next
	}
	newhead.Next = nil
	for next != nil {
		next.Next = newhead
		newhead = next
		next = tmpnext
		if next != nil {
			tmpnext = tmpnext.Next
		}
	}
	ret := make([]int, 0)
	TravelList(newhead, func(v int) {
		ret = append(ret, v)
	})

	return ret
}

func Test_reversePrint() {
	for _, item := range i06.Data {
		Dump(item)
		Dump(reversePrint(Array2List(item)))
	}
}