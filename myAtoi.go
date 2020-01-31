func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	ans := 0
	i := 0
	n := len(str)
	found := false
	isNeg := false
	for i < n && str[i] == ' ' {i++}
	if i >= n {
		return 0
	}
	if str[i] == '+' {
		found = true
	} else if str[i] == '-' {
		found = true
		isNeg = true
	} else if str[i] >= '0' && str[i] <= '9' {
		ans += int(str[i] - 48)
		found = true
	}

	if found {
		for i++; i < n ; i++ {
			if str[i] < '0' || str[i] > '9' {
				break
			} else {
				d := int(str[i] - 48)
				if !isNeg {
					max := math.MaxInt32
					if ans > max/10 || (ans == max/10 && d >= 7) {
						ans = max
						break
					} else {
						ans = ans*10 + d
					}
				} else {
					min := math.MinInt32
					if ans < min/10 || (ans == min/10 && d >= 8) {
						ans = min
						break
					} else {
						ans = ans*10 - d
					}
				}
			}
		}
	}
	return ans
}