package normal

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)

	max := -1
	maxv := -1
	for _, it := range nums {
		if _, ok := m[it]; !ok {
			m[it] = 1
		} else {
			m[it] += 1
		}

		if m[it] > max {
			max = m[it]
			maxv = it
		}

		if max > len(nums) / 2 {
			return maxv
		}
	}

	return maxv
}

func Test_majorityElement() {
	fmt.Println(majorityElement([]int{3,2,3}))
	fmt.Println(majorityElement([]int{2,2,1,1,1,2,2}))
}