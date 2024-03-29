//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
//
//
// 示例 1：
//
//
//输入：s = "()"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "()[]{}"
//输出：true
//
//
// 示例 3：
//
//
//输入：s = "(]"
//输出：false
//
//
// 示例 4：
//
//
//输入：s = "([)]"
//输出：false
//
//
// 示例 5：
//
//
//输入：s = "{[]}"
//输出：true
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁴
// s 仅由括号 '()[]{}' 组成
//
// Related Topics 栈 字符串 👍 2800 👎 0

package main

//leetcode submit region begin(Prohibit modification and deletion)
var pair = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	st := make([]byte, len(s) + 1)
	idx := 0

	for _, b := range []byte(s) {
		if b == '(' || b == '[' || b == '{' {
			st[idx] = b
			idx += 1
		} else {
			if idx > 0 && st[idx-1] == pair[b] {
				idx -= 1
			} else {
				return false
			}
		}
	}
	return idx == 0
}


//leetcode submit region end(Prohibit modification and deletion)
