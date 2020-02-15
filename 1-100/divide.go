// BF: time O(dividend/divisor)
// import "math"
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	isNeg := false
	if dividend < 0 {
		dividend = -dividend
		if divisor < 0 {
			divisor = -divisor
		} else {
			isNeg = true
		}
	} else if divisor < 0 {
		divisor = -divisor
		isNeg = true
	}
	quotient := 0
	for dividend >= divisor {
		quotient++
		dividend -= divisor
	}
	if isNeg {
		quotient = -quotient
	}
	return quotient
}

// bit-manipulations *
// By the intent, consider the case of int as int32:
// because math.MinInt32 is -1<<31 and math.MaxInt32 is 1<<31-1,
// we can't handle 1<<31 as it overflows MaxInt32, 
// so we tranfer dividend and divisor into negetive,
// and use a threshhold(half of MinInt32) to avoid overflow.
// import "math"
const thresh = math.MinInt32>>1

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {		// we can't get 1<<31
		return math.MaxInt32
	}
	dividend, divisor, isNeg := getNeg(dividend, divisor)
	quotient := 0
	for dividend <= divisor {
		num, count := divisor, 1
		for num >= thresh && dividend <= num<<1 {
			num <<= 1
			count <<= 1
		}
		// we can subtract count*divisor from dividend
		dividend -= num
		quotient += count
	}
	if isNeg {
		quotient = -quotient
	}
	return quotient
}

func getNeg(dividend, divisor int) (int, int, bool) {
	isNeg := false
	if dividend > 0 {
		dividend = -dividend
		if divisor > 0 {
			divisor = -divisor
		} else {
			isNeg = true
		}
	} else if divisor > 0 {
		divisor = -divisor
		isNeg = true
	}
	return dividend, divisor, isNeg
}