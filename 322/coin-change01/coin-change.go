package main

import (
	"fmt"
	"math"
)

var mem map[int]int

// 超时未通过
func coinChange(coins []int, amount int) int {
	mem = make(map[int]int, amount+1)
	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if _, ok := mem[amount]; ok {
		return mem[amount]
	}

	re := math.MaxInt
	for _, c := range coins {
		if amount-c < 0 {
			continue
		}
		sub := dp(coins, amount-c)
		if sub < 0 {
			continue
		}
		re = min(re, 1+sub)
	}

	if re == math.MaxInt {
		re = -1
	}
	mem[amount] = re
	return re
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Printf("%#v", coinChange([]int{2}, 3))
}
