//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入: [1,2,3,null,5,null,4]
//输出: [1,3,4]
// 
//
// 示例 2: 
//
// 
//输入: [1,null,3]
//输出: [1,3]
// 
//
// 示例 3: 
//
// 
//输入: []
//输出: []
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [0,100] 
// -100 <= Node.val <= 100 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 634 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var q [201] *TreeNode
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)

	q[0] = root
	l, r := 0, 1
	last := l
	for l < r {
		for l < r {
			if l == r - 1 {
				res = append(res, q[l].Val)
			}
			l += 1
		}

		for last < l {
			if q[last].Left != nil {
				q[r] = q[last].Left
				r += 1
			}

			if q[last].Right != nil {
				q[r] = q[last].Right
				r += 1
			}

			last += 1
		}
	}


	return res
}
//leetcode submit region end(Prohibit modification and deletion)
