func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	} else if n == 1 {
		return strs[0]
	}
	i := 0
	for ; i < len(strs[0]); i++{
		r := strs[0][i]
		for j := 1; j < n; j++ {
			if i >= len(strs[j]) || strs[j][i] != r {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:i]
}