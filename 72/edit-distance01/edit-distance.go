package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	r, c := len(word1)+1, len(word2)+1
	dp := make([][]int, r)
	for i := 0; i < r; i++ {
		dp[i] = make([]int, c)
	}

	for i := 0; i < r; i++ {
		dp[i][0] = i
	}

	for i := 0; i < c; i++ {
		dp[0][i] = i
	}

	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}

	return dp[r-1][c-1]
}

func min(x, y, z int) int {
	re := y
	if x < y {
		re = x
	}

	if re > z {
		re = z
	}
	return re
}

func main() {
	fmt.Printf("%#v", minDistance("horse", "ros"))
}
