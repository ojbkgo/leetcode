//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
// 
//
// 示例 3： 
//
// 
//输入：nums = [1]
//输出：[[1]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 6 
// -10 <= nums[i] <= 10 
// nums 中的所有整数 互不相同 
// 
// Related Topics 数组 回溯 👍 1832 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	return _sub2(nums, map[int]bool{}, []int{}, [][]int{})
}

func _sub2(nums []int, used map[int]bool, set []int, res [][]int) [][]int {
	if len(set) == len(nums) {
		res = append(res, set)
		return res
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		used[i] = true
		set = append(set, nums[i])
		res = _sub2(nums, used, set, res)
		used[i] = false
		set = append([]int{}, set[0:len(set)-1]...)
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
