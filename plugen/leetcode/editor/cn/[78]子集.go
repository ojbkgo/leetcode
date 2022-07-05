//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。 
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
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
// nums 中的所有元素 互不相同 
// 
// Related Topics 位运算 数组 回溯 👍 1516 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	set := make([]int, 0)
	//res = append(res, set)
	return _s(nums, 0, set, res)
}

func _s(nums []int, cur int, set []int, res [][]int) [][]int {
	if cur > len(nums) {
		return res
	}

	res = append(res, set)

	for i := cur; i < len(nums); i++ {
		set = append(set, nums[i])
		res = _s(nums, i + 1, set, res)
		set = append([]int{}, set[0:len(set) - 1]...)
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
