// elegant
var m = map[rune]int{
	'I':1,
	'V':5,
	'X':10,
	'L':50,
	'C':100,
	'D':500,
	'M':1000,
}
func romanToInt(s string) int {
	num := 0
	last := math.MaxInt64
	for _, r := range s {
		if m[r] > last {
			num += m[r] - 2 * last
		} else {
			num += m[r]
		}
		last = m[r]
	}
	return num
}