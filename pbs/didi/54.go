package didi

import "fmt"

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 {
		return res
	}

	res = append(res, matrix[0][0])
	x := 0
	y := 0
	direct := &dir{}

	direct.init(len(matrix[0]), len(matrix), toRight)

	finish := false
	for !finish {
		x, y, finish = direct.move(x, y)
		if x > -1 && y > -1 {
			res = append(res, matrix[y][x])
		}
	}

	return res
}

const (
	toRight = 1
	toDown  = 2
	toLeft  = 3
	toUp    = 4
)

type dir struct {
	left  int
	right int
	up    int
	down  int

	to   int
	size int
}

func (d *dir) init(xs, ys int, to int) {
	d.size = xs*ys - 1

	d.left = 0
	d.right = xs - 1
	d.up = 1
	d.down = ys - 1

	d.to = to
}

func (d *dir) move(x, y int) (int, int, bool) {
	var xx int
	var yy int
	switch d.to {
	case toRight:
		xx = x + 1
		yy = y
		if xx > d.right && d.size > 0 {
			d.to = toDown
			d.right -= 1
			xx = x
			yy = y + 1
		}
	case toDown:
		xx = x
		yy = y + 1
		if yy > d.down && d.size > 0 {
			d.to = toLeft
			d.down -= 1
			xx = x - 1
			yy = y
		}
	case toLeft:
		xx = x - 1
		yy = y
		if xx < d.left && d.size > 0 {
			d.to = toUp
			d.left += 1
			xx = x
			yy = y - 1
		}
	case toUp:
		xx = x
		yy = y - 1
		if yy < d.up && d.size > 0 {
			d.to = toRight
			d.up += 1
			xx = x + 1
			yy = y
		}
	}
	d.size -= 1

	if d.size == -1 {
		return -1, -1, true
	}

	return xx, yy, d.size == 0
}

func Test_spiralOrder() {
	data := [][]int{
		{1},
	}

	fmt.Println(spiralOrder(data))
}
