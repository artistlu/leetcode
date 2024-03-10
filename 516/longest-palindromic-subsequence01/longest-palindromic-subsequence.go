package longest_palindromic_subsequence01

func longestPalindromeSubseq(s string) int {
	sl := len(s)
	dp := make([][]int, sl)
	for i := 0; i < sl; i++ {
		dp[i] = make([]int, sl)
	}

	for i := 0; i < sl; i++ {
		dp[i][i] = 1
	}

	for i := sl - 2; i >= 0; i-- {
		for j := i + 1; j < sl; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}

		}
	}

	return dp[0][sl-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
