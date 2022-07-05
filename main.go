package main

import (
	"fmt"
	"math"
)

func main() {
	v := getResult([]int{-1, 2, 1, -4}, 1)
	fmt.Println(v)
}

func findMostClose(nums []int, picked map[int]int, pickn, target, maxClose int) int {
	if len(picked) == 3 {
		var val int
		for _, v := range picked {
			val += v
		}
		//fmt.Println(val)
		fmt.Println(picked, maxClose,  math.Abs(float64(target - val)),math.Abs(float64(target - maxClose)) )
		if math.Abs(float64(target - val)) < math.Abs(float64(target - maxClose)) {
			maxClose = val
		}

		return maxClose
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := picked[i]; ok {
			continue
		}
		picked[i] = nums[i]
		val := findMostClose(nums, picked, pickn+1, target, maxClose)
		if math.Abs(float64(target - val)) < math.Abs(float64(target - maxClose)) {
			maxClose = val
		}
		delete(picked, i)
	}

	return maxClose
}

func getResult(nums []int, target int) int {
	return findMostClose(nums, make(map[int]int), 0, target, 0)
}
