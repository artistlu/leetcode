package koko_eating_bananas01

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, arrayMax(piles)+1

	for l < r {
		mid := l + (r-l)/2
		if canFinish(piles, mid, h) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func arrayMax(nums []int) int {
	var max int
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func canFinish(piles []int, s, h int) bool {
	c := 0
	for _, p := range piles {
		c += p / s

		if p%s != 0 {
			c++
		}
	}

	return c <= h
}
