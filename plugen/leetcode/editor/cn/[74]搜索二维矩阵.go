//编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性： 
//
// 
// 每行中的整数从左到右按升序排列。 
// 每行的第一个整数大于前一行的最后一个整数。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false
// 
//
// 
//
// 提示： 
//
// 
// m == matrix.length 
// n == matrix[i].length 
// 1 <= m, n <= 100 
// -10⁴ <= matrix[i][j], target <= 10⁴ 
// 
// Related Topics 数组 二分查找 矩阵 👍 589 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	vsize := len(matrix)
	if vsize == 0 {
		return false
	}
	//hsize := len(matrix[0])

	lo, hi := 0, vsize - 1
	//l, r := 0, hsize - 1

	if vsize == 1 {
		return searchOrdered74(matrix[0], target, true) > -1
	}

	col1 := make([]int, vsize)
	for i := lo; i <= hi; i++ {
		col1[i] = matrix[i][0]
	}
	fmt.Println(col1)
	idx := searchOrdered74(col1, target, false)
	fmt.Println(idx)

	if idx == vsize {
		idx2 := searchOrdered74(matrix[idx-1], target, true)

		return idx2 > -1
	}

	if col1[idx] == target {
		return true
	} else {
		if idx == 0 {
			return false
		}

		idx2 := searchOrdered74(matrix[idx-1], target, true)

		return idx2 > -1
	}
}

func searchOrdered74(nums []int, target int, must bool) int {
	n := len(nums)
	l, r := 0, n - 1
	for l <= r {
		m := l + ((r-l) >> 1)
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}

	if (r+1 == 0 || r+1>=n || nums[r+1] != target) && must {
		return -1
	}

	return r + 1
}
//leetcode submit region end(Prohibit modification and deletion)
