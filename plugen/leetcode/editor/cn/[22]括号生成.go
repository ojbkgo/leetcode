//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// Related Topics 字符串 动态规划 回溯 👍 2444 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	dp := make([][]string, n+1)
	dp[0] = nil
	dp[1] = []string{"()"}
	l, r := "(", ")"
	_ = l
	_ = r
	for i := 2; i <= n; i++ {
		res := make([]string, 0)
		for k := 0; k < i; k++ {
			if dp[k] == nil {
				for _, it := range dp[i-k-1] {
					res = append(res, "()" + it)
				}
			} else {
				for _, it := range dp[k] {
					if dp[i-k-1] == nil {
						res = append(res, "(" + it + ")" )
					}
					for _, it2 := range dp[i-k-1] {
						res = append(res, "(" + it + ")" + it2)
					}
				}
			}
		}
		dp[i] = res
	}

	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
