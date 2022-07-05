//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。 
//
// 你可以按 任何顺序 返回答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//] 
//
// 示例 2： 
//
// 
//输入：n = 1, k = 1
//输出：[[1]] 
//
// 
//
// 提示： 
//
// 
// 1 <= n <= 20 
// 1 <= k <= n 
// 
// Related Topics 数组 回溯 👍 891 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)

func combine(n int, k int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	res := make([][]int, 0)
	path := make([]int, 0)

	return _combine(n, 1, k, path, res)
}

func _combine(n, cur, k int, path []int, res [][]int) [][]int {
	if len(path) == k {
		res = append(res, path)
		return res
	}
	for i := cur; i <= n; i++ {
		path = append(path, i)
		res = _combine(n, i+1, k, path, res)
		path = append([]int{}, path[0: len(path)-1]...)
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
