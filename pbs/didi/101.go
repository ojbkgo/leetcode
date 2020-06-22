package didi

import "fmt"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := make([]*TreeNode, 0)
	right := make([]*TreeNode, 0)

	travelLR(root.Left, &left)
	travelRL(root.Right, &right)

	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != nil && right[i] != nil && left[i].Val != right[i].Val {
			return false
		} else if (left[i] != nil && right[i] == nil) || (left[i] == nil && right[i] != nil) {
			return false
		}
	}

	return true
}

func travelLR(r *TreeNode, res *[]*TreeNode) {
	*res = append(*res, r)
	if r == nil {
		return
	}

	travelLR(r.Left, res)
	travelLR(r.Right, res)
}

func travelRL(r *TreeNode, res *[]*TreeNode) {
	*res = append(*res, r)
	if r == nil {
		return
	}

	travelRL(r.Right, res)
	travelRL(r.Left, res)
}


func Test_isSymmetric() {
	r := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Print(isSymmetric(r))
}