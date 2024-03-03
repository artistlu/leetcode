package move_zeroes01

func moveZeroes(nums []int) {
	var i, j int
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
}
