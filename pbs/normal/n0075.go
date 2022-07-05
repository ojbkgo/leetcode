package normal

import (
	"fmt"
	"sort"
)
func sortColors(nums []int) []int  {
	sort.Sort(sort.IntSlice(nums))

	return nums
}

func Test_sortColors() {
	fmt.Println(sortColors([]int{0,1,2,1,2}))
	fmt.Println(sortColors([]int{0}))
	fmt.Println(sortColors([]int{}))
}
