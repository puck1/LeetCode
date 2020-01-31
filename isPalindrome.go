// time O(2n), space O(n)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digits := make([]int, 0, 10)		// let cap 10? 
	for x != 0 {
		d := x % 10
		x /= 10
		digits = append(digits, d)
	}
	start, end := 0, len(digits) - 1
	for start < end {
		if digits[start] != digits[end] {
			return false
		} else {
			start++
			end--
		}
	}
	return true
}

// judge directly, time O(n), space O(1) *
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
  base := int(math.Pow10(int(math.Log10(float64(x)))))  
    for ; base >= 10; base /= 100 {
        if x/base != x%10 {
            return false
        }
        x = x % base / 10
    }
    return true
}

// judge if (reversed x) == x ? time O(n), space O(1) *
// will be incorrect if reversed x exceeds max int
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp, rx := x, 0
	for tmp > 0 {
		rx = rx * 10 + tmp % 10
		tmp /= 10
	}
	return rx == x
}

// judge half length of x, time O(n), space O(1) *
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}
	revered := 0
	for x > revered {
		revered = revered * 10 + x % 10
		x /= 10
	}
	return revered == x || x == revered / 10
}
