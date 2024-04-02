package house_robber01
	
var m map[int]int

func rob(nums []int) int {
	m = make(map[int]int)
	return dp(nums, 0)
}

func dp(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}

	if _, ok := m[start]; ok {
		return m[start]
	}

	re := max(dp(nums, start+1), nums[start]+dp(nums, start+2))

	m[start] = re
	return re
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
