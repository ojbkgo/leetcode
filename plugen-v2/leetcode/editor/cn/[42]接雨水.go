//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//
//
// 示例 1：
//
//
//
//
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//
// 示例 2：
//
//
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//
//
//
// 提示：
//
//
// n == height.length
// 1 <= n <= 2 * 10⁴
// 0 <= height[i] <= 10⁵
//
// Related Topics 栈 数组 双指针 动态规划 单调栈 👍 3476 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func trap(height []int) int {
	n := len(height)
	dpleft := make([]int, n)
	dpright := make([]int, n)

	dpleft[0] = height[0]
	dpright[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		if height[i] < dpleft[i-1] {
			dpleft[i] = dpleft[i-1]
		} else {
			dpleft[i] = height[i]
		}
	}

	for k := n - 2; k >= 0; k-- {
		if height[k] < dpright[k+1] {
			dpright[k] = dpright[k+1]
		} else {
			dpright[k] = height[k]
		}
	}

	var sum int
	for i := 0; i < n; i++ {
		min := dpleft[i]
		if dpright[i] < min {
			min = dpright[i]
		}

		sum += min - height[i]
	}

	return sum
}
func trap2(height []int) int {
	n := len(height)
	var i int = 1
	var total int = 0
	last := 0
	for i < n {
		godown := false
		goup := false
		// 下降
		for i < n && height[i] <= height[i-1] {
			//temp += height[i-1] - height[i]
			godown = true
			i += 1
		}

		fmt.Println(i)

		for i < n && height[i] >= height[i-1] {
			//temp += height[i] - height[i-1]
			i += 1
			goup = true
		}

		fmt.Println(i, last)

		if godown && goup {

			low, hight := height[last], height[i-1]
			if low > hight {
				low, hight = hight, low
			}

			// sum total
			for l := last; l <= i-1; l++ {
				if height[l] <= low {
					total += low - height[l]
				}
			}
		}

		last = i - 1
	}

	return total
}

//leetcode submit region end(Prohibit modification and deletion)
