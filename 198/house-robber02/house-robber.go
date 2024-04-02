package house_robber02

func rob(nums []int) int {
	f, s, c := 0, 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		c = max(f, s+nums[i])
		s = f
		f = c
	}

	return c
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
