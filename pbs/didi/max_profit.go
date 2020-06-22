package didi

import (
	"fmt"
)

func maxProfit(prices []int) int {

	if len(prices) < 2 {
		return 0
	}
	profit := prices[1] - prices[0]
	tmp := profit
	for i := 2; i < len(prices); i++ {
		sub := prices[i]-prices[i-1]
		if tmp < 0 {
			tmp = sub
		} else {
			tmp += sub
		}
		if tmp > profit {
			profit = tmp
		}
	}
	if profit > 0 {
		return profit
	} else {
		return 0
	}
}

func Test_maxProfit() {
	// 7 1 5 3 6 4
	// -6 4 -2 3 -2

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
