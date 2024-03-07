package split_array_largest_sum01

func splitArray(nums []int, k int) int {
	lo, hi := arrayMax(nums), arraySum(nums)

	for lo <= hi {
		mid := lo + (hi-lo)/2
		n := split(nums, mid)
		if n == k {
			hi = mid - 1
		} else if n > k {
			lo = mid + 1
		} else if n < k {
			hi = mid - 1
		}
	}

	return lo
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

func arraySum(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}

func split(nums []int, max int) int {
	count := 1
	sum := 0
	for _, n := range nums {
		if sum+n > max {
			count++
			sum = n
		} else {
			sum += n
		}
	}

	return count
}
