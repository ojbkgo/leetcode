package interview

import (
	. "github.com/ojbkgo/leetcode-common"
	"github.com/ojbkgo/leetcode/data/interview/i04"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	// find border

	x, find := BinarySearchLastLT(matrix[0], target)
	if find {
		return true
	}

	vl := make([]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		vl[i] = matrix[i][0]
	}
	y, find := BinarySearchLastLT(vl, target)
	if find {
		return true
	}


	for i := 0; i <= y && i < len(matrix); i++ {
		_, find = BinarySearchLastLT(matrix[i][0:x+1], target)
		if find {
			return true
		}
	}

	return false
}

func Test_findNumberIn2DArray() {
	for _, item := range i04.Data {
		Dump(item)
		result := findNumberIn2DArray(item.A1, item.A2)
		Dump(result)
	}
}