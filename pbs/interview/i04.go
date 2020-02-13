package interview

import (
	"github.com/ojbkgo/leetcode-common"
	"github.com/ojbkgo/leetcode/data/i04"
)

func findNumberIn2DArray(matrix [][]int, target int) bool {



	return true
}

func Test_findNumberIn2DArray() {
	for _, item := range i04.Data {
		common.Dump(item)
		result := findNumberIn2DArray(item.A1, item.A2)
		common.Dump(result)
	}
}