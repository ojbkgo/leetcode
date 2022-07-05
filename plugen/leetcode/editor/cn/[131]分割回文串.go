
//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
// 
//
// 示例 2： 
//
// 
//输入：s = "a"
//输出：[["a"]]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 16 
// s 仅由小写英文字母组成 
// 
// Related Topics 字符串 动态规划 回溯 👍 1023 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	return __pp(s, []string{}, [][]string{})
}

func __pp(s string, set []string, res [][]string) [][]string {
	if isclbatch(set) && s == "" {
		res = append(res, set)
	}

	for i := 0; i < len(s); i++ {
		set = append(set, s[0: i+1])
		res = __pp(s[i+1:], set, res)
		set = append([]string{}, set[0:len(set) - 1]...)
	}

	return res
}

func isclbatch(s []string) bool {
	for i := 0; i < len(s); i++ {
		if !iscl(s[i]) {
			return false
		}
	}

	return true
}

func iscl(s string) bool {
	l, r := 0, len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}
		l += 1
		r -= 1
	}

	return true
}
//leetcode submit region end(Prohibit modification and deletion)
