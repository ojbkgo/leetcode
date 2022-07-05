//编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性： 
//
// 
// 每行的元素从左到右升序排列。 
// 每列的元素从上到下升序排列。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 5
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21
//,23,26,30]], target = 20
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
// 1 <= n, m <= 300 
// -10⁹ <= matrix[i][j] <= 10⁹ 
// 每行的所有元素从左到右升序排列 
// 每列的所有元素从上到下升序排列 
// -10⁹ <= target <= 10⁹ 
// 
// Related Topics 数组 二分查找 分治 矩阵 👍 1018 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

    _, ysize := len(matrix[0]), len(matrix)

	for y := 0; y < ysize; y++ {
		_, f := search240(matrix[y], target)
		if f {
			return true
		}
	}

	return false
}

func search240(nums []int, target int) (int, bool) {
	n := len(nums)
	l, r := 0, n-1

	var m int
	for l <= r {
		m = l + ((r-l) >> 1)

		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m, true
		}
	}

	return r, false
}
//leetcode submit region end(Prohibit modification and deletion)
