//给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现
//次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。 
//
// 
//
// 示例 1： 
//
// 
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
// 
//
// 示例 2: 
//
// 
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9] 
//
// 
//
// 提示： 
//
// 
// 1 <= nums1.length, nums2.length <= 1000 
// 0 <= nums1[i], nums2[i] <= 1000 
// 
//
// 
//
// 进阶： 
//
// 
// 如果给定的数组已经排好序呢？你将如何优化你的算法？ 
// 如果 nums1 的大小比 nums2 小，哪种方法更优？ 
// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？ 
// 
// Related Topics 数组 哈希表 双指针 二分查找 排序 👍 672 👎 0
package main

import (
	"fmt"
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	sort.Sort(sort.IntSlice(nums1))
	sort.Sort(sort.IntSlice(nums2))
	fmt.Println(nums1, nums2)

	i1, i2 := 0, 0
	n1, n2 := len(nums1), len(nums2)
	res := make([]int, 0)
	for i1 < n1 && i2 < n2 {
		if nums1[i1] < nums2[i2] {
			i1 += 1
		} else if nums1[i1] > nums2[i2] {
			i2 += 1
		} else {
			res = append(res, nums2[i2])
			i1 += 1
			i2 += 1
		}
	}
	return res
}

func intersectWithMap(nums1, nums2 []int) []int {
	return nil
}
//leetcode submit region end(Prohibit modification and deletion)
