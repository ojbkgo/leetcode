//实现 strStr() 函数。
//
// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如
//果不存在，则返回 -1 。
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
//
//
// 提示：
//
//
// 1 <= haystack.length, needle.length <= 10⁴
// haystack 和 needle 仅由小写英文字符组成
//
// Related Topics 双指针 字符串 字符串匹配 👍 1429 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func strStr(haystack string, needle string) int {
	next := getNext28(needle)

	var (
		i, j = 0, 0
	)


	for ; i < len(haystack);   {
		if haystack[i] == needle[j] {
			j += 1
			if j == len(needle) {
				return i - j + 1
			}
			i += 1
			continue
		}
		//fmt.Println(haystack[i])

		//fmt.Println(needle[j])
		if j == 0 {
			i += 1
		}

		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}

	}

	return -1
}

func getNext28(str string) []int {
	n := len(str)
	next := make([]int, n)
	var j = 0
	for i := 1; i < n; i++ {
		for j > 0 && str[j] != str[i] {
			j = next[j-1]
		}

		if str[j] == str[i] {
			j += 1
		}

		next[i] = j
	}

	return next
}

//leetcode submit region end(Prohibit modification and deletion)
