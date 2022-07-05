
//给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用 一次 。 
//
// 注意：解集不能包含重复的组合。 
//
// 
//
// 示例 1: 
//
// 
//输入: candidates = [10,1,2,7,6,1,5], target = 8,
//输出:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//] 
//
// 示例 2: 
//
// 
//输入: candidates = [2,5,2,1,2], target = 5,
//输出:
//[
//[1,2,2],
//[5]
//] 
//
// 
//
// 提示: 
//
// 
// 1 <= candidates.length <= 100 
// 1 <= candidates[i] <= 50 
// 1 <= target <= 30 
// 
// Related Topics 数组 回溯 👍 867 👎 0
package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))

	return _cb(candidates, target, 0,0, []int{}, [][]int{})
}

func _cb(raw []int, target, sum, cur int, set []int, res [][]int) [][]int {
	if sum == target {
		res = append(res, set)
		return res
	}

	if sum > target {
		return res
	}
	last := 0
	for i := cur; i < len(raw); i++ {
		if raw[i] == last {
			continue
		}
		last = raw[i]

		set = append(set, raw[i])
		res = _cb(raw, target, sum + raw[i], i + 1, set, res)
		set = append([]int{}, set[0: len(set) - 1]...)
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
