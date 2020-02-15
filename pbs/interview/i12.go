package interview

import (
	. "github.com/ojbkgo/leetcode-common"
	"github.com/ojbkgo/leetcode/data/interview/i12"
)

// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/submissions/
func exist(board [][]byte, word string) bool {

	xsize := len(board)
	ysize := len(board[0])
	var x, y int
	for x < xsize  {
		y = 0
		for y < ysize {
			if board[x][y] == word[0] {
				is := make([][]bool, xsize)
				for i := 0; i < len(is); i++ {
					is[i] = make([]bool, ysize)
				}
				find := dfs(board, is, x, y, xsize, ysize, word, 0)
				if find {
					return true
				}
			}
			y += 1
		}
		x += 1
	}

	return false
}

func dfs(board [][]byte, is [][]bool, x, y, xsize, ysize int, word string, i int) bool {

	if x == xsize || y == ysize {
		return false
	}

	if board[x][y] != word[i] {
		return false
	}

	if i == len(word) - 1 {
		return true
	}

	is[x][y] = true
	if x+1 < xsize && !is[x+1][y] {
		find := dfs(board, is, x+1, y, xsize, ysize, word, i+1)
		if find {
			return true
		}

	}

	if x-1 >= 0 && !is[x-1][y] {
		find := dfs(board, is, x-1, y, xsize, ysize, word, i+1)
		if find {
			return true
		}

	}

	if y+1 < ysize && !is[x][y+1] {
		find := dfs(board, is, x, y+1, xsize, ysize, word, i+1)
		if find {
			return true
		}
	}

	if y-1 >= 0 && !is[x][y-1] {
		find := dfs(board, is, x, y-1, xsize, ysize, word, i+1)
		if find {
			return true
		}

	}

	is[x][y] = false
	return false
}

func Test_exist() {
	for _, item := range i12.Data {
		Dump(item)
		Dump(exist(item.Map, item.Word))
	}
}