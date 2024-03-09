package longest_increasing_subsequence01

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return arryNum(dp)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func arryNum(nums []int) int {
	var re int
	for _, n := range nums {
		if n > re {
			re = n
		}
	}

	return re
}
