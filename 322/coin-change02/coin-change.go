package coin_change02

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i-c < 0 {
				continue
			}

			dp[i] = min(dp[i], 1+dp[i-c])
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
