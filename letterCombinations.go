// non-recursive trackback
var m = map[byte]string {
'2':"abc",
'3':"def",
'4':"ghi",
'5':"jkl",
'6':"mno",
'7':"pqrs",
'8':"tuv",
'9':"wxyz",
}
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	s := ""
	ans := []string{}
	index := make([]int, len(digits))
	for i := 0; i >= 0; {
		s += m[digits[i]][index[i]:index[i]+1]
		index[i]++
		if i == len(digits)-1 {
			ans = append(ans, s)
			s = s[:i]
			for i >= 0 && index[i] == len(m[digits[i]]) {	// backtrack
				index[i] = 0
				i--
				if i >= 0 {
					s = s[:i]
				}
			}
		} else {
			i++
		}
	}
	return ans
}

// dfs *
var m = map[byte]string {
	'2':"abc",
	'3':"def",
	'4':"ghi",
	'5':"jkl",
	'6':"mno",
	'7':"pqrs",
	'8':"tuv",
	'9':"wxyz",
	}
func letterCombinations(digits string) []string {
	combinations := []string{}
	if len(digits) != 0 {
		helpCombine("", digits, &combinations)		
		//切片参数还要使用指针，是因为使用了append函数：https://blog.csdn.net/sgsgy5/article/details/81590184
	}
	return combinations
}
func helpCombine(combination string, digitsLeft string, combinations *[]string) {
	if len(digitsLeft) == 0 {
		*combinations = append(*combinations, combination)
	} else {
		for i := 0; i < len(m[digitsLeft[0]]); i++ {
			helpCombine(combination + m[digitsLeft[0]][i:i+1], digitsLeft[1:], combinations)
		}
	}
}