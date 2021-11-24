package normal

import (
	"fmt"
	"sort"
)

type L [][]int

func (l L) Len() int {
	return len(l)
}

func (l L) Less(i, j int) bool {
	return l[i][0] < l[j][0]
}

func (l L) Swap(i, j int) {
	l[i][0], l[j][0] = l[j][0], l[i][0]
	l[i][1], l[j][1] = l[j][1], l[i][1]
}

func merge(ll [][]int) [][]int {
	sort.Sort(L(ll))
	//fmt.Println(ll)
	res := make([][]int, 0)

	var j int
	for i := 0; i < len(ll);  {
		max := ll[i][1]
		for j = i+1 ; j < len(ll) ; {
			if max >= ll[j][0] {
				if max < ll[j][1] {
					max = ll[j][1]
				}

				j++
				continue
			} else {
				break
			}
		}

		//if i == j {
		//	res = append(res, []int{ll[i][0], ll[j][1]})
		//}

		res = append(res, []int{ll[i][0], max})
		//res = append(res, []int{ll[i][0], ll[j-1][1]})
		i = j
	}

	return res
}


func Test_merge() {
	fmt.Println(merge([][]int{{1,3},{8,10},{2,6},{15,18}}))
	fmt.Println(merge([][]int{{1,3}, {2,4}, {3,5}}))
	fmt.Println(merge([][]int{{1,2}, {3,4}, {5,6}}))
	fmt.Println(merge([][]int{{1,4}, {2,3}}))
}
