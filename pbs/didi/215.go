package didi

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {

	sort.Sort(sort.IntSlice(nums))

	for i := len(nums) - 1; i >= 0; i -= 1 {
		k--
		if k == 0 {
			return nums[i]
		}
	}

	return 0
}


func Test_findKthLargest() {
	fmt.Println(findKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4))
}