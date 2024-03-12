package minimum_insertion_steps_to_make_a_string_palindrome01

func minInsertions(s string) int {
	l := len(s)
	dp := make([][]int, l)

	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[0][l-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
