// recursion with some efficiency improvements
func isMatch(s string, p string) bool {
	ns ,np := len(s), len(p)
	// edge condition
	if ns == 0 {
		if np == 0 {
			return true
		} else if np >= 2 && np % 2 == 0 {
			for i := 1; i < np; i+=2 {
				if p[i] != '*' {
					return false
				}
			}
			return true
		} else {	// like "a**", wrong pattern
			return false
		}
	} else if np == 0 {
		return false
	}

	//ns, np != 0, 0
	i, j := 0, 0
	for i < ns && j < np {
		if j + 1 < np && p[j + 1] == '*' {
			if p[j] == '.' {
				// handle pattern ".*"
				if j + 2 >= np {
					return true
				}
				// if s[i] is not equal to p[j+2](and "." and "a*" is not after p[j]),
				// we can skip s[i] to reduce recursion scale
				for (i < ns) && (p[j + 2] != '.' && s[i] != p[j + 2]) && !(j + 3 < np && p[j + 3] == '*') {
					i++
				}
				for ; i < ns; i++ {
					// 0 or more any character(s)
					if isMatch(s[i:], p[j + 2:]) {
						return true
					}
				}
				j += 2
			} else {
				// handle pattern "a*"
				for ; i < ns && s[i] == p[j]; i++ {
					// 0 or more s[i]
					if isMatch(s[i:], p[j + 2:]) {
						return true
					}
				}
				j += 2
			}
		} else if p[j] != '.' && s[i] != p[j] {
			return false
		} else {
			// s[i] == p[j] or p[j] == '.'
			i++
			j++
		}
	}
	if i != ns {
		return false
	} else if j != np {
		return isMatch("", p[j:])
	}
	return true
}

// Complete recursion, low efficiency *
func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) != 0 && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

// recursive dp, top-down *
// dp(i, j): does s[i:] and p[j:] match?
var memo [][]int
func isMatch(s string, p string) bool {
	memo = make([][]int, len(s) + 1)
	for i := range memo {
		memo[i] = make([]int, len(p))
	}
	return dp(s, p, 0 , 0)
}
func dp(s string, p string, i int, j int) bool {
	var ans int
	if j == len(p) {	// can use (len(s)+1) * (len(p)+1) matrix instead
		return i == len(s)
	}
	if memo[i][j] == 1 {
		return true
	} else if memo[i][j] == 2 {
		return false
	}
	firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')
	if j + 1 < len(p) && p[j+1] == '*' {
		if dp(s, p, i, j+2) || (firstMatch && dp(s, p, i+1, j)) {
			ans = 1
		} else {
			ans = 2
		}
	} else {
		if firstMatch && dp(s, p, i+1, j+1) {
			ans = 1
		} else {
			ans = 2
		}
	}
	switch ans {
	case 1:
		memo[i][j] = 1
		return true
	default:	// case 2
		memo[i][j] = 2
		return false
	}
}

// non-recursive dp, bottom-up *
var dp [][]bool
func isMatch(s string, p string) bool {
	dp = make([][]bool, len(s) + 1)
	for i := range dp {
		dp[i] = make([]bool, len(p) + 1)
	}
	dp[len(s)][len(p)] = true		// empty string mathches another one
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}