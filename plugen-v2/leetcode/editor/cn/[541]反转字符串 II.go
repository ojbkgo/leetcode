//给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。 
//
// 
// 如果剩余字符少于 k 个，则将剩余字符全部反转。 
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "abcdefg", k = 2
//输出："bacdfeg"
// 
//
// 示例 2： 
//
// 
//输入：s = "abcd", k = 2
//输出："bacd"
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 10⁴ 
// s 仅由小写英文组成 
// 1 <= k <= 10⁴ 
// 
// Related Topics 双指针 字符串 👍 296 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	n := len(s)
	b := []byte(s)

	for i := 0; i < n; i+=2*k {
		l := i
		r := 0
		r = i + k - 1
		if r >= n {
			r = n - 1
		}
		//if i + 2 * k > n && i + k < n {
		//} else {
		//
		//}
		for l < r && r < n {
			b[l], b[r] = b[r], b[l]
			l += 1
			r -= 1
		}
	}

	return string(b)
}
//leetcode submit region end(Prohibit modification and deletion)
