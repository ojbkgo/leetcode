//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
// Related Topics 字符串 动态规划 👍 4864 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	i, j := _lcs5(s)

	//fmt.Println(i, j)

	//fmt.Println(s[i-1 : j])

	return s[i : j+1]
}

func _lcs5(s string) (int, int) {
	dp := make([][]bool, len(s))

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	maxi, maxj := 0, 0
	max := -1

	for L := 2; L <= len(s); L++ {
		for i := 0; i < len(s); i++ {
			j := i + L - 1
			if j >= len(s) {
				break
			}
			if s[i] == s[j] {
				if j - i == 1 {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] && max < j - i + 1 {
					max = j - i + 1
					maxi, maxj = i, j
					//fmt.Println(i, j)
				}
			} else {
				dp[i][j] = false
			}
		}
	}

	//fmt.Println(dp)
	//fmt.Println(maxi, maxj)
	return maxi, maxj
}

//leetcode submit region end(Prohibit modification and deletion)
