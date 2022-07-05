//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。 
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。 
//
// 
//
// 
//
// 示例 1： 
//
// 
//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 
//
// 示例 2： 
//
// 
//输入：digits = ""
//输出：[]
// 
//
// 示例 3： 
//
// 
//输入：digits = "2"
//输出：["a","b","c"]
// 
//
// 
//
// 提示： 
//
// 
// 0 <= digits.length <= 4 
// digits[i] 是范围 ['2', '9'] 的一个数字。 
// 
// Related Topics 哈希表 字符串 回溯 👍 1763 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	str, cc := getstr(digits)
	return _letter(str, cc, len(digits),0,0, "", []string{})
}

func _letter(str string, cc []int, max, level, cur int, path string, res []string) []string {
	//n := len(str)
	if len(path) == max {
		res = append(res, path)
		return res
	}

	for i := cur; i < len(str); i++ {
		if i < cur + cc[level] {
			res = _letter(str, cc, max, level + 1, cur + cc[level], path + string(str[i]), res)
		}
	}

	return res
}


var m = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func getstr(digits string) (string, []int) {
	res := ""
	num := []int{}

	for _, b := range digits {
		res += m[byte(b)]
		num = append(num, len(m[byte(b)]))
	}

	return res, num
}
//leetcode submit region end(Prohibit modification and deletion)
