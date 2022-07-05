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
// Related Topics 数组 回溯 👍 974 👎 0
package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	res := make([]int, 0)
	used := make(map[int]bool)
	result := make([][]int, 0)
	return _combine40(candidates, target, 0, 0, used, res, result)
}


func _combine40(nums []int, target, sum, cur int, used map[int]bool, res []int, result [][]int) [][]int {
	if sum > target {
		return result
	}

	if sum == target {
		temp := make([]int, 0)
		temp = append(temp, res...)
		result = append(result, temp)
		return result
	}

	single := make(map[int]bool)
	for i := cur; i < len(nums); i++ {
		if _, ok := used[i]; ok {
			continue
		}
		if _, ok := single[nums[i]]; ok {
			continue
		}
		single[nums[i]] = true

		used[i] = true
		res = append(res, nums[i])
		result = _combine40(nums, target, sum + nums[i], i + 1, used, res, result)
		delete(used, i)
		res = res[0: len(res) - 1]
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)
