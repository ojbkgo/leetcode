//给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。 
//
//
// 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。 
//
//
// 
// 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。 
// 
//
// 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。 
//
// 
//
// 示例 1： 
//
// 
//输入：tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
//输出：["JFK","MUC","LHR","SFO","SJC"]
// 
//
// 示例 2： 
//
// 
//输入：tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL",
//"SFO"]]
//输出：["JFK","ATL","JFK","SFO","ATL","SFO"]
//解释：另一种有效的行程是 ["JFK","SFO","ATL","JFK","ATL","SFO"] ，但是它字典排序更大更靠后。
// 
//
// 
//
// 提示： 
//
// 
// 1 <= tickets.length <= 300 
// tickets[i].length == 2 
// fromi.length == 3 
// toi.length == 3 
// fromi 和 toi 由大写英文字母组成 
// fromi != toi 
// 
// Related Topics 深度优先搜索 图 欧拉回路 👍 523 👎 0
package main

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
type arr [][]string

func (a arr) Len() int {
	return len(a)
}

func (a arr) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}

	return a[i][0] < a[j][0]
}

func (a arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func findItinerary(tickets [][]string) []string {
	sort.Sort(arr(tickets))

	return _tk(tickets, make([]bool, len(tickets)), [][]string{})
}

func _tk(tickets [][]string, used []bool, set [][]string) []string {
	if len(set) == len(tickets) {
		return _merge_tk(set)
	}

	if len(set) > 0 && set[0][0] != "JFK" {
		return nil
	}

	for i := 0; i < len(tickets); i++ {
		if used[i] {
			continue
		}

		if len(set) > 0 && set[len(set) - 1][1] != tickets[i][0] {
			continue
		}

		used[i] = true
		set = append(set, tickets[i])
		res := _tk(tickets, used, set)
		if res != nil {
			return res
		}
		used[i] = false
		set = append([][]string{}, set[0: len(set) - 1]...)
	}

	return nil
}

func _merge_tk(tks [][]string) []string {
	if len(tks) == 0 {
		return nil
	}
	res := make([]string, 0)

	res = append(res, tks[0][0], tks[0][1])
	for i := 1; i < len(tks); i++ {
		res = append(res, tks[i][1])
	}

	return res
}
//leetcode submit region end(Prohibit modification and deletion)
