
//n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
//
// Related Topics 数组 回溯 👍 1210 👎 0
package main

import (
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	pos := make([][]string, n)
	for i := 0; i < n; i++ {
		pos[i] = make([]string, n)
	}

	return _qq(n, 0,  pos, [][]string{})
}

func _qq(n int, row int, pos [][]string, res [][]string) [][]string {
	if row >= n {
		it := make([]string, 0)
		for _, v := range pos {
			it = append(it, strings.Join(v, ""))
		}

		res = append(res, it)
		return res
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			pos[row][j] = "."
		}
		if canput(n, i, row, pos) {
			pos[row][i] = "Q"
		} else {
			continue
		}

		res = _qq(n, row+1, pos, res)
		for j := 0; j < n; j++ {
			pos[row][j] = ""
		}
	}


	return res
}

var dir = [][]int{
	{1, 0},   // 右边
	{-1, 0},  // 左边
	{0, -1},  // 上边
	{0, 1},   // 下边
	{-1, -1}, // 左上
	{-1, 1},  // 左下
	{1, -1},  // 右上
	{1, 1},   // 右下

}

func canput(n, x, y int, pos [][]string) bool {

	if pos[y][x] == "Q" {
		return false
	}

	if !bounder(n, x, y) {
		return false
	}

	for _, it := range dir {
		for i := 1; i <= n*2; i++ {
			if !bounder(n, x+it[0]*i, y+it[1]*i) {
				break
			}

			if pos[y+it[1]*i][x+it[0]*i] == "Q" {
				return false
			}
		}
	}

	return true
}

func bounder(n, x, y int) bool {
	if x >= 0 && x < n && y >= 0 && y < n {
		return true
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
