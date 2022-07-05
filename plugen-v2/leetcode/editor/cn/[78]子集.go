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
// Related Topics 位运算 数组 回溯 👍 1645 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	return subsets78(nums, 0, make([]int, 0), make([][]int, 0))
}

func subsets78(nums []int, cur int, path []int, result [][]int) [][]int {
	{
		temp := make([]int, 0)
		temp = append(temp, path...)
		result = append(result, temp)
	}

	for i := cur; i < len(nums); i++ {
		path = append(path, nums[i])
		result = subsets78(nums, i+1, path, result)
		path = path[0: len(path) - 1]
	}

	return result
}
//leetcode submit region end(Prohibit modification and deletion)
