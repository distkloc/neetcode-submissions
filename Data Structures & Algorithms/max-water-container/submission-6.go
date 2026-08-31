func maxArea(heights []int) int {
	res := 0

	left := 0
	right := len(heights) - 1

	for left < right {
		h := min(heights[left], heights[right])
		res = max(res, (right - left) * h)
		
		if heights[left] <= heights[right] {
			left++
		} else {
			right--
		}
	}

	return res
}
