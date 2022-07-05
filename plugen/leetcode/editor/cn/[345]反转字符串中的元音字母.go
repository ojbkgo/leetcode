//给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。 
//
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "hello"
//输出："holle"
// 
//
// 示例 2： 
//
// 
//输入：s = "leetcode"
//输出："leotcede" 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 3 * 10⁵ 
// s 由 可打印的 ASCII 字符组成 
// 
// Related Topics 双指针 字符串 👍 240 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func reverseVowels(s string) string {
	l ,r := 0, len(s) - 1
	sb := []byte(s)
	for l < r {
		for l < r && !isy(s[l]) {l+=1}
		for l < r && !isy(s[r]) {r-=1}
		sb[l], sb[r] = sb[r], sb[l]
		l += 1
		r -= 1
	}

	return string(sb)
}

var yy = map[byte]bool {
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func isy(c byte) bool {
	_, ok := yy[c]

	return ok
}
//leetcode submit region end(Prohibit modification and deletion)
