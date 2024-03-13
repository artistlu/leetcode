package main

import "fmt"

func maxA(n int) int {
	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1

		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], dp[j-2]+dp[j-2]*(i-j))
		}
	}

	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Printf("i:%#v a: %v\n", i, maxA(i))
	}
}
