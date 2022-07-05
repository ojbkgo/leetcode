
//有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词
//不同，你能对此优化吗? 
//
// 示例： 
//
// 
//输入：words = ["I","am","a","student","from","a","university","in","a","city"], 
//word1 = "a", word2 = "student"
//输出：1 
//
// 提示： 
//
// 
// words.length <= 100000 
// 
// Related Topics 数组 字符串 👍 49 👎 0
package main

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func findClosest(words []string, word1 string, word2 string) int {
	idx := map[string]int{}
	dist := 100001
	for i, v := range words {
		idx[v] = i
		_, ok1 := idx[word1]
		_, ok2 := idx[word2]
		if ok1 && ok2 {
			t := math.Abs(float64(idx[word1] - idx[word2]))
			if int(t) < dist {
				dist = int(t)
			}
		}
	}

	return dist
}
//leetcode submit region end(Prohibit modification and deletion)
