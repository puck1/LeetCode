// import "math"
func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans * 10 + x % 10
		x /= 10
	}
	if ans < math.MinInt32 || ans > math.MaxInt32 {
	// if ans < -(1 << 31) || ans > 1 << 31 - 1 {
		ans = 0
	}
	return ans
}

// * 
// slow but correct in 32-bit system
// import "math"
func reverse(x int) int {
	ans := 0
	for x != 0 {
		r := x % 10
		if ans > math.MaxInt32 / 10 || (ans == math.MaxInt32 / 10 && r > 7) {
			// math.MaxInt32 : 2147483647
			return 0
		}
		if ans < math.MinInt32 / 10 || (ans == math.MinInt32 / 10 && r > 8) {
			// math.MinInt32 : -2147483648
			return 0
		}
		ans = ans * 10 + r
		x /= 10
	}
	return ans
}