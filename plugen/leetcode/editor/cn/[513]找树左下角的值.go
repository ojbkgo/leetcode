//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。 
//
// 假设二叉树中至少有一个节点。 
//
// 
//
// 示例 1: 
//
// 
//
// 
//输入: root = [2,1,3]
//输出: 1
// 
//
// 示例 2: 
//
// 
//
// 
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
// 
//
// 
//
// 提示: 
//
// 
// 二叉树的节点个数的范围是 [1,10⁴] 
// -2³¹ <= Node.val <= 2³¹ - 1 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 244 👎 0
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
func findBottomLeftValue(root *TreeNode) int {
	maxdep, maxval = 0, root.Val
	maxleft(root, 1)
	return maxval
}

var maxdep, maxval int

func maxleft(root *TreeNode, dep int)  {
	if root.Left == nil && root.Right == nil {
		// 叶子节点
		if dep > maxdep {
			maxdep = dep
			maxval = root.Val
		}
	}

	if root.Left != nil {
		maxleft(root.Left, dep + 1)
	}
	if root.Right != nil {
		maxleft(root.Right, dep + 1)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
