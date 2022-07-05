//给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,2]
//输出：
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// Related Topics 数组 回溯 👍 969 👎 0
package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	return _sub3(nums, make([]bool, len(nums)), []int{}, [][]int{})
}

func _sub3(nums []int, used []bool, set []int, res [][]int) [][]int {
	if len(set) == len(nums) {
		res = append(res, set)
		return res
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
			continue
		}
		if used[i] {
			continue
		}

		set = append(set, nums[i])
		used[i] = true
		res = _sub3(nums, used, set, res)
		used[i] = false
		set = append([]int{}, set[0:len(set)-1]...)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
