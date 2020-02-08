func maxArea(height []int) int {
	max := 0
	now := 0
	for l, r := 0, len(height) - 1; l < r; {
		if height[l] > height[r] {
			now = height[r] * (r - l)
			r--
		} else {
			now = height[l] * (r - l)
			l++
		}
		if now > max {
			max = now
		}
	}
	return max
}