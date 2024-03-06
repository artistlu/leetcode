package find_first_and_last_position_of_element_in_sorted_array01

func searchRange(nums []int, target int) []int {
	f := left_binary(nums, target)
	if f == -1 {
		return []int{-1, -1}
	}

	return []int{f, right_binary(nums, target)}

}

func left_binary(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			r = mid - 1
		}
	}

	if l < 0 || l > len(nums)-1 {
		return -1
	}

	if nums[l] == target {
		return l
	}

	return -1
}

func right_binary(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] == target {
			l = mid + 1
		}
	}

	if r < 0 || r > len(nums)-1 {
		return -1
	}

	if nums[r] == target {
		return r
	}

	return -1
}
