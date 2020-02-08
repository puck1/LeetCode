// approach 1 - BF
func lengthOfLongestSubstring(s string) int {
	max := 1
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if strings.IndexByte(s[i:j], s[i]) == -1 {
				if j - i + 1 > max {
					max = j - i + 1
				}
			} else {
				break
			}
		}
	}  
	return max
}

// approach 2 - slide window with map *
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)	
    max := 0
    n := len(s)
    for i, j := 0, 0 ; j < n ; j++ {
		if k, ok := m[s[j]]; ok && k >= i {
			i = k + 1
		} else if j - i + 1 > max {
			max = j - i + 1
		}
        m[s[j]] = j
	}
	return max
}