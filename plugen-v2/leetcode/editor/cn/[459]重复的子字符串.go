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
// Related Topics 字符串 字符串匹配 👍 695 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func repeatedSubstringPattern(s string) bool {
	next := getNext459(s)
	n := len(s)

	//fmt.Println(next)

	if next[n-1] <= 0 {
		return false
	}

	var k int
	for k = n-1; k > 0; k-- {
		if next[k-1] == 0 || next[k] <= next[k-1] {
			break
		}
	}
	//fmt.Println(k)

	return n % (n - next[n-1]) == 0
}

func getNext459(str string) []int {
	n := len(str)

	res := make([]int, n)

	var (
		i int
		j int
	)
	for i = 1; i < n; i++ {
		for j > 0 && str[i] != str[j] {
			j = res[j-1]
		}

		if str[i] == str[j] {
			j += 1
		}

		res[i] = j
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
