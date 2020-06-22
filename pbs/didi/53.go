package didi

import "fmt"

func maxSubArray(nums []int) int {
	cur := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if cur < 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}

		if sum < cur {
			sum = cur
		}
	}

	return sum
}

func Test_maxSubArray() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}