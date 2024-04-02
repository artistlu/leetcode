package house_robber_ii01

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(helper(nums, 1, n-1), helper(nums, 0, n-2))
}

func helper(nums []int, s, e int) int {
	c, f, se := 0, 0, 0
	for i := e; i >= s; i-- {
		c = max(f, se+nums[i])
		se = f
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
