
//有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
//
// 
// 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 
//和 "192.168@1.1" 是 无效 IP 地址。 
// 
//
// 给定一个只包含数字的字符串 s ，用以表示一个 IP 地址，返回所有可能的有效 IP 地址，这些地址可以通过在 s 中插入 '.' 来形成。你 不能 重新
//排序或删除 s 中的任何数字。你可以按 任何 顺序返回答案。 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "25525511135"
//输出：["255.255.11.135","255.255.111.35"]
// 
//
// 示例 2： 
//
// 
//输入：s = "0000"
//输出：["0.0.0.0"]
// 
//
// 示例 3： 
//
// 
//输入：s = "101023"
//输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
// 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 20 
// s 仅由数字组成 
// 
// Related Topics 字符串 回溯 👍 820 👎 0
package main

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	if s == "0000" { return []string{"0.0.0.0"}}
	return _restore(s, len(s), []string{}, []string{})
}

func _restore(s string, size int, set []string, res []string) []string {
	//if len(set) > 0 && set[0] == "0" {
	//	return res
	//}

	for idx, it := range set {
		if idx < len(set) - 1 && len(it) > 3 {
			return res
		}
	}

	if len(set) == 4 {
		ip := strings.Join(set, ".")
		if isip(set) && len(ip) - 3 == size {
			res = append(res, ip)
		}
		return res
	}

	//fmt.Println(s)
	for i := 0; i < 3; i++ {
		if i + 1 > len(s) {
			break
		}
		set = append(set, s[0:i+1])
		res = _restore(s[i+1:], size, set, res)
		set = append([]string{}, set[0:len(set) - 1]...)
	}

	return res
}

func isip(str []string) bool {
	if len(str) != 4 {
		return false
	}
	for _, it := range str {
		 v, _ := strconv.Atoi(it)
		if v <0 || v > 255 {
			return false
		}

		if len(it) > 1 && strings.HasPrefix(it, "0") {
			return false
		}
	}

	return true
}
//leetcode submit region end(Prohibit modification and deletion)
