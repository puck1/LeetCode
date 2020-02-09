// 1. 
// import "strings" 
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 2. BF
func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	if l2 == 0 {
		return 0
	}
	for t := 0; t <= l1 - l2; t++{
		if haystack[t] == needle[0] {
			i, j:= t, 0
			for ; j < l2 && haystack[i] == needle[j]; i, j = i+1, j+1 {}
			if j == l2 {
				return t
			}
		}
	}
	return -1
}

// 3. KMP *
func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	if l2 == 0 {
		return 0
	}
	i, j := 0, 0
	next := getNext(needle)
	for i < l1 && j < l2 {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == l2 {
		return i - j
	} else {
		return -1
	}
}

func getNext(needle string) (next []int) {		// note: needle can't be empty string
	n := len(needle)
	next = make([]int, len(needle), len(needle))
	next[0] = -1
	k := -1
	j := 1
	for j < n {
		if k == -1 || needle[j-1] == needle[k] {
			k++
			if needle[j] != needle[k] {
				next[j] = k
			} else {
				next[j] = next[k]
			}
			j++
		} else {
			k = next[k]
		}
	}
	return
}

// 4. BM algorithm *
// ...

// 5. Sunday algorithm *
func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	if l2 == 0 {
		return 0
	}

	lastIndexMap := make(map[byte]int)
	for i := range needle {
		lastIndexMap[needle[i]] = i
	}

	for i := 0; i <= l1-l2; {
		if haystack[i:i+l2] == needle {
			return i
		}
		i += l2
		if i < l1 {
			if k, ok := lastIndexMap[haystack[i]]; ok {
				i -= k
			} else {
				i++
			}
		}
	}
	return -1
}