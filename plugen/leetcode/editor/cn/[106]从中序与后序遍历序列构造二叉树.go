//给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并
//返回这颗 二叉树 。 
//
// 
//
// 示例 1: 
//
// 
//输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
//输出：[3,9,20,null,null,15,7]
// 
//
// 示例 2: 
//
// 
//输入：inorder = [-1], postorder = [-1]
//输出：[-1]
// 
//
// 
//
// 提示: 
//
// 
// 1 <= inorder.length <= 3000 
// postorder.length == inorder.length 
// -3000 <= inorder[i], postorder[i] <= 3000 
// inorder 和 postorder 都由 不同 的值组成 
// postorder 中每一个值都在 inorder 中 
// inorder 保证是树的中序遍历 
// postorder 保证是树的后序遍历 
// 
// Related Topics 树 数组 哈希表 分治 二叉树 👍 687 👎 0
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[n-1]}
	if len(inorder) == 1 {
		return root
	}

	var i int
	for i = 0; i < n; i++ {
		if inorder[i] == postorder[n-1] {
			break
		}
	}

	var (
		left []int
		postleft []int
		right []int
		postright []int
	)
	if i > 0 {
		left = inorder[0: i]
		postleft = postorder[0: i]
		root.Left = buildTree(left, postleft)
	}
	if i < n - 1 {
		right = inorder[i+1:]
		postright = postorder[i: n-1]
		root.Right = buildTree(right, postright)
	}


	return root
}

//leetcode submit region end(Prohibit modification and deletion)
