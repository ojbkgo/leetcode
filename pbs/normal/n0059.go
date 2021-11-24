package normal

import "fmt"
func generateMatrix(n int) [][]int {
	x := 0
	y := 0

	dx := 1
	dy := 0

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}


	for v := 1; v <= n*n; v++ {
		res[y][x] = v
		//fmt.Println(x, y)
		if (x == n - 1 && dx == 1) || (x == 0 && dx == -1) || (y == 0 && dy == -1) || (y == n-1 && dy == 1) || res[y+dy][x+dx] != 0 {
			dx, dy = changeDir(dx, dy)
		}

		x += dx
		y += dy
	}

	return res
}

func changeDir(dx, dy int) (int, int) {
	if dx == 1 && dy == 0 {
		return 0, 1
	} else if dx == 0 && dy == 1 {
		return -1, 0
	} else if dx == -1 && dy == 0 {
		return 0, -1
	} else if dx == 0 && dy == -1 {
		return 1, 0
	}

	return 0, 0
}

func Test_generateMatrix() {
	fmt.Println(generateMatrix(3))
}
