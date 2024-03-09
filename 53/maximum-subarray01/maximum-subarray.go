package maximum_subarray01

func maxSubArray(nums []int) int {
	var (
		re, pre, cur int
	)

	pre = nums[0]
	re = pre
	for i := 1; i < len(nums); i++ {
		cur = max(pre+nums[i], nums[i])
		re = max(re, cur)
		pre = cur
	}

	return re

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
