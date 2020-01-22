// approach 1 - expand around center *
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	min, max := 0, 0
	for i := 0; i < n-1; i++ {
		if i == 0 {
			max = 1
		} else {
			left, right := i-1, i+1
			for left >= 0 && right < n && s[left] == s[right] {
				if right-left+1 > max-min {
					max = right + 1
					min = left
				}
				left--
				right++
			}
		}
		left, right := i, i+1
		for left >= 0 && right < n && s[left] == s[right] {
			if right-left+1 > max-min {
				max = right + 1
				min = left
			}
			left--
			right++
		}
	}
	return s[min:max]
}

// approach 2 - dynamic programming *
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	dp := make([][]bool, n)
	min, max := 0, 0
	for i := 0; i < n-1; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			if max-min < 1 {
				max = i + 1
				min = i
			}
		} else {
			dp[i][i+1] = false
		}
	}
	dp[n-1] = make([]bool, n)
	dp[n-1][n-1] = true

	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				if j-i > max-min {
					max = j
					min = i
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return s[min : max+1]
}