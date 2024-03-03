package container_with_most_water

import "math"

func maxArea(height []int) int {
	var max int
	i, j := 0, len(height)-1
	for i < j {
		max = int(math.Max(float64(max), math.Min(float64(height[i]), float64(height[j]))*float64(j-i)))
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
