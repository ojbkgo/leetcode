//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
// 
//
// 示例 2： 
//
// 
//输入：head = [1,2]
//输出：[2,1]
// 
//
// 示例 3： 
//
// 
//输入：head = []
//输出：[]
// 
//
// 
//
// 提示： 
//
// 
// 链表中节点的数目范围是 [0, 5000] 
// -5000 <= Node.val <= 5000 
// 
//
// 
//
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？ 
// 
// 
// Related Topics 递归 链表 👍 2330 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	return reverseListRecursive(head)
	pre, cur := (*ListNode)(nil), head

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}

	return pre
}

func reverseListRecursive(head *ListNode) *ListNode {

	return reverse(nil, head)
}

func reverse(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	temp := cur.Next
	cur.Next = pre
	pre = cur
	return reverse(pre, temp)
}

//leetcode submit region end(Prohibit modification and deletion)
