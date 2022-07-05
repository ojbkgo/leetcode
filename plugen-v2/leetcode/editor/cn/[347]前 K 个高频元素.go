//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//
//
// 示例 2:
//
//
//输入: nums = [1], k = 1
//输出: [1]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
//
//
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
// Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 1192 👎 0
package main

import (
	heap2 "container/heap"
	"fmt"
)

//leetcode submit region begin(Prohibit modification and deletion)
import heap2 "container/heap"
func newheap(size int) *minheap {
	heap := &minheap{
		data: make([]heapval, size + 1),
		size: size,
		heapsize: 0,
	}
	return heap
}

type heapval struct {
	val    int
	weight int
}

type minheap struct {
	data     []heapval
	size     int
	heapsize int
}

func (h *minheap) push(val heapval) {
	if h.heapsize == h.size {
		h.pop()
	}
	h.heapsize += 1
	h.data[h.heapsize] = val
	h.justifyUp()
}

func (h *minheap) justifyUp() {
	i := h.heapsize

	for i > 1 {
		p := h.parent(i)
		if h.data[p].weight > h.data[i].weight {
			h.data[p], h.data[i] = h.data[i], h.data[p]
			i = p
			continue
		}
		break
	}
}

func (h *minheap) justifyDown() {
	p := 1
	l := h.left(p)
	r := h.right(p)

	for p < h.heapsize {
		if h.data[l].weight > h.data[r].weight {

		} else {

		}
	}
}

func (h *minheap) pop() heapval {
	val := h.data[1]

	h.data[1] = h.data[h.heapsize]
	h.heapsize -= 1
	h.justifyDown()
	return val
}

func (m *minheap) peek() heapval {
	return m.data[1]
}

func (m *minheap) left(i int) int {
	return i << 1
}

func (m *minheap) right(i int) int {
	return (i << 1) + 1
}

func (m *minheap) parent(i int) int {
	return i >> 1
}

type Heap []heapval

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	hv := *h
	return hv[i].weight < hv[j].weight
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	v := x.(heapval)

	*h = append(*h, v)
}

func (h *Heap) Pop() interface{} {
	hv := *h
	n := len(hv)
	ret := hv[n-1]
	*h = hv[0: n-1]
	return ret
}

func topKFrequent(nums []int, k int) []int {

	h := &Heap{}
	heap2.Init(h)

	c := make(map[int]int)
	for _, v := range nums {
		if _, ok := c[v]; !ok {
			c[v] = 0
		}

		c[v] += 1
	}

	for key, val := range c {
		heap2.Push(h, heapval{val: key, weight: val})
		if h.Len() > k {
			heap2.Pop(h)
		}
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		val := heap2.Pop(h).(heapval)
		res = append(res, val.val)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
