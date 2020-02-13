package interview

import (
	"github.com/ojbkgo/leetcode/common"
	"github.com/ojbkgo/leetcode/data/i03"
)

func findRepeatNumber(nums []int) int {

	r := make(map[int]bool)

	for _, it := range nums {
		if _, ok := r[it]; ok {
			return it
		} else {
			r[it] = true
		}
	}

	return -1
}

func Test_findRepeatNumber() {
	for _, item := range i03.I03data {
		common.Dump(item)
		res := findRepeatNumber(item)
		common.Dump(res)
	}
}
