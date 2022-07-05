package normal

import (
	"encoding/json"
	"fmt"
)

type pos struct {
	x  int
	y  int
	dx int
	dy int
}

func rotate(matrix [][]int) [][]int {
	if len(matrix) < 2 {
		return matrix
	}

	n := len(matrix)

	r := n - 2
	if r == 0 {
		r = 1
	}

	for i := 0; i < r; i++ {
		sp := startpos(i, n)
		for j := 0; j < n-i*2-1; j++ {
			fmt.Println("move", i, j)
			matrix[sp[0].y][sp[0].x], matrix[sp[1].y][sp[1].x], matrix[sp[2].y][sp[2].x], matrix[sp[3].y][sp[3].x] =
				matrix[sp[3].y][sp[3].x], matrix[sp[0].y][sp[0].x], matrix[sp[1].y][sp[1].x], matrix[sp[2].y][sp[2].x]
			for k := 0; k < 4; k++ {
				sp[k].x += sp[k].dx
				sp[k].y += sp[k].dy
			}
		}

	}

	return matrix
}

//9,4,7
//8,5,2
//3,6,1

func startpos(n int, size int) []pos {
	sp := []pos{
		{0, 0, 1, 0},                // 左上角
		{size - 1, 0, 0, 1},         //右上角
		{size - 1, size - 1, -1, 0}, //右下角
		{0, size - 1, 0, -1},        // 左下角
	}

	sp[0].x, sp[0].y = sp[0].x+n, sp[0].y+n
	sp[1].x, sp[1].y = sp[1].x-n, sp[1].y+n
	sp[2].x, sp[2].y = sp[2].x-n, sp[2].y-n
	sp[3].x, sp[3].y = sp[3].x+n, sp[3].y-n

	return sp
}

func Test_rotate() {
	data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	v := `[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]`
	//v = `[[1, 2],[3, 4]]`
	data = [][]int{}
	json.Unmarshal([]byte(v), &data)

	fmt.Println(rotate(data))
}
