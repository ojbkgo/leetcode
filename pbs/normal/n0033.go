package normal

import "fmt"

func singleNumber(nums []int) int {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	for _, it := range nums {

		if _, ok := m1[it]; !ok {
			m1[it] = true
		} else {
			m2[it] = true
		}
	}

	for k, _ := range m1 {
		if _, ok := m2[k]; !ok {
			return k
		}
	}

	return 0
}

func Test_singleNumber() {
	fmt.Println(singleNumber([]int{1, 2, 3, 4, 1, 2, 3, 6, 6}))
}