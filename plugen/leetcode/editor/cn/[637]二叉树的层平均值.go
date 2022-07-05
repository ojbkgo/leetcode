//给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10⁻⁵ 以内的答案可以被接受。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：[3.00000,14.50000,11.00000]
//解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
//因此返回 [3, 14.5, 11] 。
// 
//
// 示例 2: 
//
// 
//
// 
//输入：root = [3,9,20,15,7]
//输出：[3.00000,14.50000,11.00000]
// 
//
// 
//
// 提示： 
//
// 
//
// 
// 树中节点数量在 [1, 10⁴] 范围内 
// -2³¹ <= Node.val <= 2³¹ - 1 
// 
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 315 👎 0
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
var q [10001]*TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0)

	if root == nil {
		return res
	}

	l, r := 0, 1
	q[0] = root
	last := 0
	for l < r {
		sum := 0.0
		for l < r {
			sum += float64(q[l].Val)
			l += 1
		}
		res = append(res, sum / float64(r - last))

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
