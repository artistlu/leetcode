package super_egg_drop01

import (
	"fmt"
	"math"
)

var m map[string]int

// 超出时间限制
func superEggDrop(k int, n int) int {
	m = make(map[string]int)
	return dp(k, n)
}

func dp(k, n int) int {
	if k == 1 {
		return n
	}

	if n == 0 {
		return 0
	}

	key := fmt.Sprintf("%d:%d", k, n)
	if _, ok := m[key]; ok {
		return m[key]
	}

	res := math.MaxInt
	for i := 1; i <= n; i++ {
		res = min(res, max(dp(k-1, i-1), dp(k, n-i))+1)
	}

	m[key] = res

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
