//这是一个交互问题。
//
// 您有一个升序整数数组，其长度未知。您没有访问数组的权限，但是可以使用 ArrayReader 接口访问它。你可以调用 ArrayReader.get(i)
//:
//
//
//
// 返回数组第iᵗʰ个索引(0-indexed)处的值(即 secret[i])，或者
//
//
// 如果 i 超出了数组的边界，则返回 2³¹ - 1
//
//
//
// 你也会得到一个整数 target。
//
// 如果存在secret[k] == target，请返回索引 k 的值；否则返回 -1
//
// 你必须写一个时间复杂度为 O(log n) 的算法。
//
//
//
// 示例 1：
//
//
//输入: secret = [-1,0,3,5,9,12], target = 9
//输出: 4
//解释: 9 存在在 nums 中，下标为 4
//
//
// 示例 2：
//
//
//输入: secret = [-1,0,3,5,9,12], target = 2
//输出: -1
//解释: 2 不在数组中所以返回 -1
//
//
//
// 提示：
//
//
// 1 <= secret.length <= 10⁴
// -10⁴ <= secret[i], target <= 10⁴
// secret 严格递增
//
// Related Topics 数组 二分查找 交互 👍 50 👎 0
package main

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int { return 0 }

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

var max = (1 << 31) - 1
var step = 1000

func search(reader ArrayReader, target int) int {
	l ,r := 0, 9999
	lv, rv := reader.get(l), reader.get(r)

	if lv == target { return lv}
	if rv == target { return rv}
	if lv > target || rv < target {
		return -1
	}

	for rv == max {
		r = (r >> 1) + 1
		rv = reader.get(r)
	}

}

//leetcode submit region end(Prohibit modification and deletion)
