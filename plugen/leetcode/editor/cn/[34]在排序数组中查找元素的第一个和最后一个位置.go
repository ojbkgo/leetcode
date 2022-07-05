//给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。 
//
// 如果数组中不存在目标值 target，返回 [-1, -1]。 
//
// 进阶： 
//
// 
// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？ 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4] 
//
// 示例 2： 
//
// 
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1] 
//
// 示例 3： 
//
// 
//输入：nums = [], target = 0
//输出：[-1,-1] 
//
// 
//
// 提示： 
//
// 
// 0 <= nums.length <= 10⁵ 
// -10⁹ <= nums[i] <= 10⁹ 
// nums 是一个非递减数组 
// -10⁹ <= target <= 10⁹ 
// 
// Related Topics 数组 二分查找 👍 1503 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	res404 := []int{-1, -1}
	idx := searchOrdered34(nums, target)
	if idx == -1 || len(nums) == 0 {
		return res404
	}

	l, r := idx-1, idx + 1

	for l >= 0 {
		if nums[l] == target {
			l -= 1
		} else {
			break
		}
	}

	for r < len(nums) {
		if nums[r] == target {
			r += 1
		} else {
			break
		}
	}

	return []int{l+1, r-1}
}

func searchOrdered34(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n - 1

	for l <= r {
		m := l + ((r - l) >> 1)
		fmt.Println(l, r, m)
		if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		} else {
			fmt.Println(m)
			return m
		}
	}

	if r+1 == 0 || r+1 >= n || nums[r+1] != target {
		return -1
	}

	fmt.Println(r+1)

	return r + 1
}
//leetcode submit region end(Prohibit modification and deletion)
