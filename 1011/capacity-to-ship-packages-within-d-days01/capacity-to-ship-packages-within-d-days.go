package capacity_to_ship_packages_within_d_days01

func shipWithinDays(weights []int, days int) int {
	l, r := arrayMax(weights), arraySum(weights)

	for l <= r {
		mid := l + (r-l)/2
		n := split(weights, mid)
		if n == days {
			r = mid - 1
		} else if n < days {
			r = mid - 1
		} else if n > days {
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
