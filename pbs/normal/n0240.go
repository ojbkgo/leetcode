package normal

import (
	"encoding/json"
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {

	//x, y := 0, 0


	return true
}

func Test_searchMatrix() {
	v := `[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]`
	data := [][]int{}

	json.Unmarshal([]byte(v), &data)
	fmt.Println(searchMatrix(data, 5))
}
