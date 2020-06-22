package didi

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		i int
		j int
	)

	res := make([]int, 0)
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i += 1
		} else {
			res = append(res, nums2[j])
			j += 1
		}
	}

	for i < m {
		res = append(res, nums1[i])
		i += 1
	}

	for j < n {
		res = append(res, nums2[j])
		j += 1
	}

	for k := 0; k < m+n; k++ {
		nums1[k] = res[k]
	}
}

func Test_merge() {
	l1 := []int{1, 2, 3, 0, 0, 0}
	l2 := []int{2, 5, 6}

	merge(l1, 3, l2, 3)
	fmt.Print(l1)
}
