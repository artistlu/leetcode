package trapping_rain_water01

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	len := len(height)
	lmax, rmax := height[0], height[len-1]
	l, r := 0, len-1

	var ans int
	for l < r {
		lmax = max(lmax, height[l])
		rmax = max(rmax, height[r])

		if lmax < rmax {
			ans += lmax - height[l]
			l++
		} else {
			ans += rmax - height[r]
			r--
		}
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}

	return i
}
