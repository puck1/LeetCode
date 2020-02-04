func generateParenthesis(n int) []string {
	ans := make([]string, 0, 2*n)
	if n != 0 {
		helpGenerate("", 0, n, &ans)
	}
	return ans
}

func helpGenerate(s string, count int, n int, ans *[]string) {
	// count: the number of '(' unmatched in string s
	// n: the number of '(' unmatched in total
	if count == n {
		for ; n > 0; n-- {
			s += ")"
		}
		*ans = append(*ans, s)
		return
	}
	helpGenerate(s+"(", count+1, n, ans)
	if count > 0 {
		helpGenerate(s+")", count-1, n-1, ans)
	}
}

// closure Number *
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}		// must return "" but empty array
	}
	ans := []string{}
	for c := 0; c < n; c++ {
		for _, left := range generateParenthesis(c) {
			for _, right := range generateParenthesis(n-c-1) {
				ans = append(ans, "(" + left + ")" + right)
			}
		}
	}
	return ans
}