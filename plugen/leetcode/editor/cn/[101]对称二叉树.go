//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 1787 👎 0
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
func isSymmetric(root *TreeNode) bool {
	return cmp(root.Left, root.Right)
}

func cmp(l, r *TreeNode) bool {
	if l == nil && r != nil {
		return false
	} else if l != nil && r == nil {
		return false
	} else if l == nil && r == nil {
		return true
	}

	if !cmp(l.Left, r.Right) {
		return false
	}

	if !cmp(l.Right, r.Left) {
		return false
	}

	if l.Val != r.Val {
		return false
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
