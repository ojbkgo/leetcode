//给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位
//。
//
// 返回 滑动窗口中的最大值 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
//输出：[3,3,5,5,6,7]
//解释：
//滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// 示例 2：
//
//
//输入：nums = [1], k = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
// Related Topics 队列 数组 滑动窗口 单调队列 堆（优先队列） 👍 1637 👎 0
package main

//leetcode submit region begin(Prohibit modification and deletion)

func maxSlidingWindow(nums []int, k int) []int {
	st := newSingularStack239(k)

	for i := 0; i < k; i++ {
		st.push(nums[i])
	}
	res := make([]int, 0)
	res = append(res, st.front())
	for i := k; i < len(nums); i++ {
		//fmt.Println("to shift:", st.data, nums[i-k])
		st.shift(nums[i-k])
		//fmt.Println("to push:", st.data, nums[i])
		st.push(nums[i])
		//fmt.Println("pushed:", st.data)
		res = append(res, st.front())
	}

	return res
}

type singularStack239 struct {
	data []int
}

func (s *singularStack239) push(val int) {
	var k int
	for k = len(s.data) - 1; k >= 0; k-- {
		if val <= s.data[k] {
			break
		}
	}

	if k == -1 {
		s.data = make([]int, 0)
	} else {
		s.data = s.data[0:k+1]
	}


	s.data = append(s.data, val)
}

func (s *singularStack239) shift(val int) {
	if len(s.data) > 0 && s.data[0] == val {
		s.data = s.data[1:]
	}
}

func (s *singularStack239) front() int {
	return s.data[0]
}

func newSingularStack239(size int) *singularStack239 {
	s := &singularStack239{
		data: make([]int, 0),
	}

	//for i := 0; i < len(s.data); i++ {
	//	s.data[i] = 99999
	//}

	return s
}

//leetcode submit region end(Prohibit modification and deletion)
