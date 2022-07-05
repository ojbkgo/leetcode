//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。 
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。 
//
// 
// 
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0]
//输出：[[],[0]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 10 
// -10 <= nums[i] <= 10 
// 
// 
// 
// Related Topics 位运算 数组 回溯 👍 833 👎 0
package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return subsets90(nums, 0, make([]int, 0), make([][]int, 0))
}

func subsets90(nums []int, cur int, res []int, result [][]int ) [][]int {
	{
		temp := make([]int, 0)
		temp = append(temp, res...)
		result = append(result, temp)
	}

	uniq := make(map[int]bool)
	for i := cur; i < len(nums); i++ {
		if _, ok := uniq[nums[i]]; ok {
			continue
		}
		uniq[nums[i]] = true

		res = append(res, nums[i])
		result = subsets90(nums, i+1, res, result)
		res = res[0:len(res) - 1]
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)
