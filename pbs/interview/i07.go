package interview

import (
	. "github.com/ojbkgo/leetcode-common"
	"github.com/ojbkgo/leetcode/data/interview/i07"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{}
	buildTreeRec(preorder, inorder, root)

	return root
}

func buildTreeRec(preorder []int, inorder []int, root *TreeNode) {
	if len(preorder) == 0 {
		return
	}

	root.Val = preorder[0]
	idx := ArrayIndexOf(inorder, root.Val)
	left := idx

	var (
		preleft []int
		preright []int
		inleft []int
		inright []int
	)

	if len(preorder) > 1 {
		preleft = preorder[1: 1+left]
	}
	if 1 + left < len(preorder) {
		preright = preorder[1+left:]
	}

	inleft = inorder[0:idx]
	if idx + 1 < len(inorder) {
		inright = inorder[idx+1:]
	}

	if len(preleft) > 0 {
		root.Left = &TreeNode{}
		buildTreeRec(preleft, inleft, root.Left)
	}

	if len(preright) > 0 {
		root.Right = &TreeNode{}
		buildTreeRec(preright, inright, root.Right)
	}
}




func Test_buildTree() {
	for _, item := range i07.Data {
		Dump(item)
		DumpTree(buildTree(item.PreList, item.MidList))
	}
}

