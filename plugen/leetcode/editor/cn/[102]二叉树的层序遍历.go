//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
// Related Topics 树 广度优先搜索 二叉树 👍 1206 👎 0
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

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return nil
	}
	q := make([]*TreeNode, 2001)
	q[0] = root
	l, r := 0, 1
	last := l
	for l < r {
		it := make([]int, 0)
		for l < r { it = append(it, q[l].Val); l += 1 }
		res = append(res, it)
		for last < l {
			if q[last].Left != nil { q[r] = q[last].Left;r += 1 }
			if q[last].Right != nil { q[r] = q[last].Right; r += 1 }
			last += 1
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
