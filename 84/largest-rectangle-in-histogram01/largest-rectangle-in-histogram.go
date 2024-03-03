package largest_rectangle_in_histogram01

func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	heights = append(heights, 0)
	ln := len(heights)
	var res int
	for i := 0; i < ln; i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			l := stack[len(stack)-1]
			res = max(res, (i-l-1)*h)
		}

		stack = append(stack, i)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
