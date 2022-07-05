package main


type stack struct {
	node []*TreeNode
	i    int
	cap  int
}

func (s *stack) pop() *TreeNode {
	if s.empty() {
		return nil
	}
	no := s.node[s.i]
	s.i -= 1
	return no
}

func (s *stack) push(node *TreeNode) {
	if s.i + 1 == s.cap {
		s.resize()
	}

	s.i += 1
	s.node[s.i] = node
}

func (s *stack) top() *TreeNode {
	if s.empty() {
		return nil
	}

	return s.node[s.i]
}

func (s *stack) resize() {
	nnode := make([]*TreeNode, s.cap * 2)
	for idx, v := range s.node {
		nnode[idx] = v
	}

	s.cap = s.cap * 2
}

func (s *stack) empty() bool {
	return s.i < 0
}

func newstack(size int) *stack {
	return &stack{
		node: make([]*TreeNode, size),
		cap: size,
		i: -1,
	}
}


type queue struct {

}