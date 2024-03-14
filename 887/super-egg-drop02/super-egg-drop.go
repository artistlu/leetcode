package super_egg_drop02

func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	var m int
	for dp[k][m] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i][m-1] + dp[i-1][m-1] + 1

		}
	}

	return m
}
