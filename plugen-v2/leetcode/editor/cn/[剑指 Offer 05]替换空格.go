//请实现一个函数，把字符串 s 中的每个空格替换成"%20"。 
//
// 
//
// 示例 1： 
//
// 输入：s = "We are happy."
//输出："We%20are%20happy." 
//
// 
//
// 限制： 
//
// 0 <= s 的长度 <= 10000 
// Related Topics 字符串 👍 289 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func replaceSpace(s string) string {
	count := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count += 1
		}
	}

	b := make([]byte, len(s) + count * 2)

	i, j := len(s) - 1, len(b) - 1
	for i >= 0 {
		if s[i] != ' ' {
			b[j] = s[i]
			j -= 1
			i -= 1
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			j -= 3
			i -= 1
		}
	}

	return string(b)
}

//leetcode submit region end(Prohibit modification and deletion)
