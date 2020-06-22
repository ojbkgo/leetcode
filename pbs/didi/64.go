package didi

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	yc := len(grid)
	xc := len(grid[0])

	dp := make([][]int, yc)
	for c := 0; c < yc; c++ {
		dp[c] = make([]int, xc)
		for i := 0; i < xc; i++ {
			dp[c][i] = -1
		}
	}

	for c := 0; c < yc; c++ {
		for i := 0; i < xc; i++ {
			if c == 0 && i == 0 {
				dp[c][i] = grid[c][i]
			} else if c == 0 {
				dp[c][i] = dp[c][i-1] + grid[c][i]
			} else if i == 0 {
				dp[c][i] = dp[c-1][i] + grid[c][i]
			} else {
				dp[c][i] = int(math.Min(float64(dp[c-1][i]), float64(dp[c][i-1]))) + grid[c][i]
			}
		}
	}

	return dp[yc-1][xc-1]
}

func Test_minPathSum() {
	data := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(data))
}