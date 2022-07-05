package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//type Node struct {
//	Val      int
//	Children []*Node
//}

type Node struct {
	Val       int
	Neighbors []*Node
}
