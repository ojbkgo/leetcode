//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 749 👎 0
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
func preorderTraversal(root *TreeNode) []int {
	return bystack144(root)
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}
	return res
}


func bystack144(root *TreeNode) []int {
	st := newstack(10)
	res := make([]int, 0)
	if root == nil {
		return res
	}

	st.push(root)
	for !st.empty() {
		top := st.pop()
		res = append(res, top.Val)
		if top.Right != nil {
			st.push(top.Right)
		}
		if top.Left != nil {
			st.push(top.Left)
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
