package nsum

func nSum(nums []int, n, start, target int) [][]int {
	if n < 2 || len(nums)-start < n {
		return [][]int{}
	}

	if n == 2 {
		lo, hi := start, len(nums)-1
		var tSum = [][]int{}
		for lo < hi {
			sum := nums[lo] + nums[hi]
			if sum > target {
				hi--
				for lo < hi && nums[hi] == nums[hi+1] {
					hi--
				}
			} else if sum < target {
				lo++
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
			} else {
				tSum = append(tSum, []int{nums[lo], nums[hi]})
				lo++
				hi--
				for lo < hi && nums[hi] == nums[hi+1] {
					hi--
				}
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
			}
		}

		return tSum
	}

	res := [][]int{}
	for i := start; i <= len(nums)-n; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		trs := nSum(nums, n-1, i+1, target-nums[i])
		for _, r := range trs {
			res = append(res, append([]int{nums[i]}, r...))
		}
	}

	return res
}
