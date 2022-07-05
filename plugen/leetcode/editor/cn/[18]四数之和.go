//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）： 
//
// 
// 0 <= a, b, c, d < n 
// a、b、c 和 d 互不相同 
// nums[a] + nums[b] + nums[c] + nums[d] == target 
// 
//
// 你可以按 任意顺序 返回答案 。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// 
//
// 示例 2： 
//
// 
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= nums.length <= 200 
// -10⁹ <= nums[i] <= 10⁹ 
// -10⁹ <= target <= 10⁹ 
// 
// Related Topics 数组 双指针 排序 👍 1147 👎 0
package main

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	sort.Sort(sort.IntSlice(nums))

	res := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return res
	}

	if nums[n - 1] * 4 < target {
		return res
	}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		for j := i + 1; j < n-1 ; j++ {
			//if nums[i] + nums[j] > target {
			//	fmt.Println(i, j)
			//	return res
			//}

			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}

			l := j + 1
			r := n - 1

			//if l + 1 <= r && nums[l+1] > 0 && nums[i] + nums[j] + nums[l] > target {
			//	continue
			//}
			//
			//if l + 1 <= r && nums[l+1] < 0 && nums[i] + nums[j] + nums[l] < target {
			//	continue
			//}

			for l < r {
				v := nums[i] + nums[j] + nums[l] + nums[r]
				//fmt.Println(i, j, l, r, v)
				if v > target {
					r -= 1
				} else if v < target {
					l += 1
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})

					for l < r && nums[l] == nums[l + 1] {l+=1}
					for l < r && nums[r] == nums[r - 1] {r-=1}
					l += 1
					r -= 1
				}
			}
		}
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
