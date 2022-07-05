//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,5]
//输出：["1->2->5","1->3"]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：["1"]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100
//
// Related Topics 树 深度优先搜索 字符串 回溯 二叉树 👍 672 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0, 10)

	res = travel(root, "", res)

	return res
}

func travel(root *TreeNode, path string, res []string) []string {
	if root == nil {
		return res
	}

	if root.Left == nil && root.Right == nil {
		path += fmt.Sprint(root.Val)
		res = append(res, path)
		return res
	}

	if root.Left != nil {
		res = travel(root.Left, path+fmt.Sprint(root.Val)+"->", res)
	}

	if root.Right != nil {
		res = travel(root.Right, path+fmt.Sprint(root.Val)+"->", res)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
