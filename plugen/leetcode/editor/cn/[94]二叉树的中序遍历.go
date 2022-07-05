//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
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
//输出：[2,1]
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
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// Related Topics 栈 树 深度优先搜索 二叉树 👍 1191 👎 0

package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)

func inorderTraversal(root *TreeNode) []int {
	return travelstack94(root)
	res := make([]int, 0)

	if root == nil {
		return res
	}

	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

func print(st *stack94) {
	for i := st.i; i >= 0; i-- {
		fmt.Println(st.node[i].Val)
	}

	fmt.Println("==========")
}

func travelstack94(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	st := newstack94(10)

	cur := root

	for cur != nil || !st.empty() {
		if cur != nil {
			st.push(cur)
			cur = cur.Left
		} else {
			cur = st.pop()
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}

	//st.push(root)
	//
	//for !st.empty() {
	//	top := st.top()
	//	print(st)
	//	if top.Left != nil {
	//		st.push(top.Left)
	//		continue
	//	} else {
	//		res = append(res, top.Val)
	//		st.pop()
	//		if top.Right != nil {
	//			st.push(top.Right)
	//		}
	//	}
	//}

	return res
}


type stack94 struct {
	node []*TreeNode
	i    int
	cap  int
}

func (s *stack94) pop() *TreeNode {
	if s.empty() {
		return nil
	}
	no := s.node[s.i]
	s.i -= 1
	return no
}

func (s *stack94) push(node *TreeNode) {
	if s.i + 1 == s.cap {
		s.resize()
	}

	s.i += 1
	s.node[s.i] = node
}

func (s *stack94) top() *TreeNode {
	if s.empty() {
		return nil
	}

	return s.node[s.i]
}

func (s *stack94) resize() {
	nnode := make([]*TreeNode, s.cap * 2)
	for idx, v := range s.node {
		nnode[idx] = v
	}

	s.cap = s.cap * 2
	s.node = nnode
}

func (s *stack94) empty() bool {
	return s.i < 0
}

func newstack94(size int) *stack94 {
	return &stack94{
		node: make([]*TreeNode, size),
		cap: size,
		i: -1,
	}
}


//leetcode submit region end(Prohibit modification and deletion)
