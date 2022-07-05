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
// Related Topics 回溯 👍 987 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func combine(n int, k int) [][]int {
	return _combine(n, k, 1, make([]int, 0), make([][]int, 0))
}

func _combine(n int, k int, cur int, res []int, result [][]int) [][]int {
	if len(res) == k {
		temp := make([]int, 0)
		temp = append(temp, res...)
		result = append(result, temp)
		return result
	}

	for i := cur; i <= n; i++ {
		res = append(res, i)
		result = _combine(n, k, i+1, res, result)
		res = res[0: len(res) - 1]
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)
