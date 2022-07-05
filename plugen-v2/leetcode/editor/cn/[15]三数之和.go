//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 4784 👎 0
package main

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return res
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for l, r := i+1, n-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l += 1
				r -= 1
				for l < r && nums[l] == nums[l-1] {
					l += 1
				}
				for l < r && nums[r] == nums[r+1] {
					r -= 1
				}
			} else if sum > 0 {
				r -= 1
			} else if sum < 0 {
				l += 1
			}
		}
	}

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
