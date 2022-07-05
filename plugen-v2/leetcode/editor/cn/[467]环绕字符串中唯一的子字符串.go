//把字符串 s 看作是 “abcdefghijklmnopqrstuvwxyz” 的无限环绕字符串，所以 s 看起来是这样的： 
//
// 
// "...zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcd...." . 
// 
//
// 现在给定另一个字符串 p 。返回 s 中 唯一 的 p 的 非空子串 的数量 。 
//
// 
//
// 示例 1: 
//
// 
//输入: p = "a"
//输出: 1
//解释: 字符串 s 中只有一个"a"子字符。
// 
//
// 示例 2: 
//
//
//输入: p = "cac"
//输出: 2
//解释: 字符串 s 中的字符串“cac”只有两个子串“a”、“c”。.
// 
//
// 示例 3: 
//
// 
//输入: p = "zab"
//输出: 6
//解释: 在字符串 s 中有六个子串“z”、“a”、“b”、“za”、“ab”、“zab”。
// 
//
// 
//
// 提示: 
//
// 
// 1 <= p.length <= 10⁵ 
// p 由小写英文字母构成 
// 
// Related Topics 字符串 动态规划 👍 230 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func findSubstringInWraproundString(p string) int {

	dp := make([]int, 26)
	count := make([]int, len(p))

	last := 1
	count[0] = last
	for i := 1; i < len(p); i++ {
		if p[i] - p[i-1] == 1 {
			last += 1
		} else if p[i] == 'a' && p[i-1] == 'z' {
			last += 1
		} else {
			last = 1
		}

		count[i] = last
	}

	for i := 0; i < len(p); i++ {
		idx := p[i] - 'a'
		if count[i] > dp[idx] {
			dp[idx] = count[i]
		}
	}

	sum := 0
	for _, v := range dp {
		sum += v
	}

	return sum
}
//leetcode submit region end(Prohibit modification and deletion)
