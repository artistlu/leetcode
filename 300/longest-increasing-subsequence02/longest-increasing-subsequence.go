package longest_increasing_subsequence02

func lengthOfLIS(nums []int) int {
	top := make([]int, len(nums))

	piles := 0
	for _, n := range nums {
		l, r := 0, piles
		for l < r {
			mid := l + (r-l)/2
			if top[mid] > n {
				r = mid
			} else if top[mid] < n {
				l = mid + 1
			} else if top[mid] == n {
				r = mid
			}
		}

		if l == piles {
			piles++
		}

		top[l] = n
	}

	return piles
}
