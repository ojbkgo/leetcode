//实现 strStr() 函数。 
//
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如
//果不存在，则返回 -1 。 
//
// 
//
// 说明： 
//
// 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。 
//
// 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与 C 语言的 strstr() 以及 Java 的 indexOf() 定义相符。 
//
// 
//
// 示例 1： 
//
// 
//输入：haystack = "hello", needle = "ll"
//输出：2
// 
//
// 示例 2： 
//
// 
//输入：haystack = "aaaaa", needle = "bba"
//输出：-1
// 
//
// 示例 3： 
//
// 
//输入：haystack = "", needle = ""
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// 0 <= haystack.length, needle.length <= 5 * 10⁴ 
// haystack 和 needle 仅由小写英文字符组成 
// 
// Related Topics 双指针 字符串 字符串匹配 👍 1136 👎 0
package main

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	if haystack != "" && len(needle) == 0 {
		return 0
	}
	next := getnext(needle)
	//fmt.Println(next)
	n := len(haystack)
	k := 0
	for i := 0; i < n; i++ {
		for k > 0 && haystack[i] != needle[k] {
			k = next[k-1]
		}

		if haystack[i] == needle[k] {
			k += 1
		}
		//fmt.Println(i, k)
		if k == len(needle) {
			return i - k + 1
		}
	}

	return -1
}


func strStr2(haystack string, needle string) int {
	fmt.Println(getnext(needle))
	return 0

	if haystack == needle {
		return 0
	}
	if haystack != "" && len(needle) == 0 {
		return 0
	}
	dp := getdp(needle)
	stat := 0
	for i := 0; i < len(haystack); i++ {
		stat = dp[stat][haystack[i]]
		if stat == len(needle) {
			return i - stat + 1
		}
	}

	return -1
}

var M = 256

func getdp(s string) [][]int {
	n := len(s)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, M)
	}


	dp[0][s[0]] = 1
	var X = 0
	for i := 1; i < n; i++ {
		for j := 0; j < M; j++ {
			if s[i] == byte(j) {
				dp[i][j] = i + 1
			} else {
				dp[i][j] = dp[X][j]
			}
		}
		X = dp[X][s[i]]
	}

	return dp
}


func getnext(s string) []int {
	n := len(s)
	next := make([]int, n)

	next[0] = 0
	k := 0
	for i := 1; i < n; i++ {
		for k > 0 && s[k] != s[i] {
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
