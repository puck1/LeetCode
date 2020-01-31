// import "bytes"
func intToRoman(num int) string {
	var buffer bytes.Buffer
	for ; num >= 1000; num -= 1000 {
		buffer.WriteByte('M')
	}
	if num >= 900 {
		buffer.WriteString("CM")
		num -= 900
	} else if num >= 500 {
		buffer.WriteByte('D')
		num -= 500
	} else if num >= 400 {
		buffer.WriteString("CD")
		num -= 400
	}
	for ; num >= 100; num -= 100 {
		buffer.WriteByte('C')
	}
	if num >= 90 {
		buffer.WriteString("XC")
		num -= 90
	} else if num >= 50 {
		buffer.WriteByte('L')
		num -= 50
	} else if num >= 40 {
		buffer.WriteString("XL")
		num -= 40
	}
	for ; num >= 10; num -= 10 {
		buffer.WriteByte('X')
	}
	if num >= 9 {
		buffer.WriteString("IX")
		num -= 9
	} else if num >= 5 {
		buffer.WriteByte('V')
		num -= 5
	} else if num >= 4 {
		buffer.WriteString("IV")
		num -= 4
	}
	for ; num > 0; num-- {
		buffer.WriteByte('I')
	}
	return buffer.String()
}


// elegant, low efficiency *
func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	lets := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    rom := make([]string, 0, 20)

	for i := 0; i < len(vals); i++ {
		n := num / vals[i]
		for j := 0; j < n; j++ {
            rom = append(rom, lets[i])
		}

		num %= vals[i]
	}

    return strings.Join(rom, "")
}