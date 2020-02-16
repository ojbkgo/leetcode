package interview

import (
	"fmt"
	. "github.com/ojbkgo/leetcode-common"
)

// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
func movingCount(m int, n int, k int) int {

	count := 0

	cm := make([][]bool, m)
	pass := make([][]bool, m)

	for i := 0; i < m; i++ {
		cm[i] = make([]bool, n)
		pass[i] = make([]bool, n)
	}

	dfsi13(cm, pass, m, n, 0, 0, k, &count)

	return count
}


func dfsi13(cm, pass [][]bool, m, n, x, y, k int, count *int) {

	if y == m || x == n {
		return
	}

	if sum(x, y) > k {
		return
	}

	if !cm[y][x] {
		*count += 1
		cm[y][x] = true
	}

	pass[y][x] = true
	if x+1 < n && !pass[y][x+1] {
		// right
		dfsi13(cm, pass, m, n, x+1, y, k, count)
	}

	if x-1 >= 0 && !pass[y][x-1] {
		// right
		dfsi13(cm, pass, m, n, x-1, y, k, count)
	}

	if y+1 < m && !pass[y+1][x] {
		// up
		dfsi13(cm, pass, m, n, x, y+1, k, count)
	}

	if y-1 >= 0 && !pass[y-1][x] {
		//down
		dfsi13(cm, pass, m, n, x, y-1, k, count)
	}

}

func sum(x, y int) int {
	return SumArray(SplitNumber(x)) + SumArray(SplitNumber(y))
}

func Test_movingCount() {
	fmt.Println(movingCount(11, 8, 16))
	//fmt.Println(movingCount(3, 1, 2))
}
