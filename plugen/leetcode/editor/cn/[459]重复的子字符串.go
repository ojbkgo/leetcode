//给定一个非空的字符串 s ，检查是否可以通过由它的一个子串重复多次构成。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abab"
//输出: true
//解释: 可由子串 "ab" 重复两次构成。
// 
//
// 示例 2: 
//
// 
//输入: s = "aba"
//输出: false
// 
//
// 示例 3: 
//
// 
//输入: s = "abcabcabcabc"
//输出: true
//解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)
// 
//
// 
//
// 提示： 
//
// 
//
// 
// 1 <= s.length <= 10⁴ 
// s 由小写英文字母组成 
// 
// Related Topics 字符串 字符串匹配 👍 630 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	next := getnext459(s)
	n := len(s)
	last := next[n - 1]
	return last > 0 && (n % (n-last)) == 0
}

func getnext459(s string) []int {
	n := len(s)
	next := make([]int, n)

	k := 0
	for i := 1; i < n; i++ {
		for k > 0 && s[i] != s[k] {
			k = next[k-1]
		}

		if s[i] == s[k] {
			k += 1
		}

		next[i] = k
	}

	return next
}
//leetcode submit region end(Prohibit modification and deletion)
