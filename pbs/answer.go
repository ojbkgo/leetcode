package  main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	list := make([]*TreeNode, 0)
	list = append(list, root)

	for len(list) > 0 {
		nextList := make([]*TreeNode, 0)
		resCur := make([]int, 0)
		st := 0
		for st < len(list) {
			resCur = append(resCur, list[st].Val)
			if list[st].Left != nil {
				nextList = append(nextList, list[st].Left)
			}

			if list[st].Right != nil {
				nextList = append(nextList, list[st].Right)
			}
			st += 1
		}
		res = append(res, resCur)
		list = nextList
	}

	return res
}

func main() {
	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(levelOrder(t))
}