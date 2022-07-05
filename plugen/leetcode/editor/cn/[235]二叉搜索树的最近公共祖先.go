//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。 
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。” 
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5] 
//
// 
//
// 
//
// 示例 1: 
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6 
//解释: 节点 2 和节点 8 的最近公共祖先是 6。
// 
//
// 示例 2: 
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
//输出: 2
//解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。 
//
// 
//
// 说明: 
//
// 
// 所有节点的值都是唯一的。 
// p、q 为不同节点且均存在于给定的二叉搜索树中。 
// 
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 👍 787 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	p1, p2 := getpath(root, p), getpath(root, q)
	s1, s2 := len(p1), len(p2)
	i, last := 0, -1
	for i < s1 && i < s2 && p1[i].Val == p2[i].Val {
		last = i
		i += 1
	}


	return p1[last]
}

func getpath(root, no *TreeNode) []*TreeNode {
	path := make([]*TreeNode, 0)

	for root != nil {
		path = append(path, root)
		if root.Val == no.Val {
			return path
		} else if root.Val > no.Val {
			root = root.Left
		} else if root.Val < no.Val {
			root = root.Right
		}
	}

	return path
}

//leetcode submit region end(Prohibit modification and deletion)
