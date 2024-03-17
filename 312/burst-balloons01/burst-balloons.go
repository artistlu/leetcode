package burst_balloons01

func maxCoins(nums []int) int {
	var points []int
	nums = append(nums, 1)
	points = append([]int{1}, nums...)

	ln := len(points)
	dp := make([][]int, ln)

	for i := range dp {
		dp[i] = make([]int, ln)
	}

	for i := ln - 2; i >= 0; i-- {
		for j := i + 1; j <= ln-1; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+points[i]*points[j]*points[k])
			}
		}
	}

	return dp[0][ln-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
