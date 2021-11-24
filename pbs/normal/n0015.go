package normal

import "fmt"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)

	if len(nums) < 3 {
		return res
	}

	return nil
}


func Test_threeSum2() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0}))
	fmt.Println(threeSum([]int{}))
}
