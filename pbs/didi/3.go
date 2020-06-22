package didi

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}
	sn := len(s)
	var max = -1

	for i := 0; i < sn; i++ {
		uniq := make(map[byte]bool)
		cc := 0
		for j := i; j < sn; j++ {
			if i == j {
				cc = 1
				uniq[s[i]] = true
			} else if i == j-1 {
				//fmt.Println("b===", i, j, s[i:j], string(s[j]))
				if s[i] != s[j] {
					cc = 2
					uniq[s[j]] = true
				} else {
					i = j - 1
					break
				}
			} else {
				//fmt.Println("a===", i, j, s[i:j], string(s[j]))
				if _, ok := uniq[s[j]]; !ok {
					cc += 1
					uniq[s[j]] = true
				} else {
					i = strings.Index(s[i:j], string(s[j]))  + i
					//fmt.Println(i)
					cc = j - i + 1
					break
				}

			}
			if cc > max {
				max = cc
			}
		}

	}

	return max
}

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//
//	if len(s) == 1 {
//		return 1
//	}
//	sn := len(s)
//	dp := make([][]int, sn)
//	for i := 0; i < sn; i++ {
//		dp[i] = make([]int, sn)
//	}
//
//	max := -1
//	for i := 0; i < sn; i++ {
//		for j := i ; j < sn; j++ {
//			if i == j {
//				dp[i][j] = 1
//			} else if i == j -1 {
//				if s[i] != s[j] {
//					dp[i][j] = 2
//				} else {
//					i = j
//					dp[i][j] = 1
//				}
//			} else {
//				//fmt.Println("a===", i, j, s[i: j], string(s[j]))
//				if !strings.Contains(string(s[i: j]), string(s[j])) {
//					dp[i][j] = dp[i][j-1] + 1
//				} else {
//					i = strings.Index(s[i: j], string(s[j])) + 1 + i
//					//fmt.Println(i)
//					dp[i][j] = j - i + 1
//				}
//				//fmt.Println("b===", i, j, s[i: j], string(s[j]))
//			}
//			if dp[i][j] > max {
//				max = dp[i][j]
//			}
//		}
//	}
//	return max
//}

func Test_lengthOfLongestSubstring() {
	fmt.Println(lengthOfLongestSubstring("aab"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
