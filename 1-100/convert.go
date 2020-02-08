// approach 1 - visit by Row
// import "bytes"
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	var buffer bytes.Buffer
	cycleLen := 2*numRows - 2
// 	for i := 0; i < numRows; i++ {
// 		if i == 0 || i == numRows-1 {
// 			for j := i; j < n; j += 2*numRows - 2 {
// 				buffer.WriteByte(s[j])
// 			}
// 		} else {
// 			j := i
// 			for j < n {
// 				buffer.WriteByte(s[j])
// 				j += 2*(numRows-i) - 2
// 				if j < n {
// 					buffer.WriteByte(s[j])
// 					j += 2 * i
// 				}
// 			}
// 		}
// 	}
	for i := 0; i < numRows; i++ {
		for j := 0; i+j < n; j += cycleLen {
			buffer.WriteByte(s[i+j])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				buffer.WriteByte(s[j+cycleLen-i])
			}
		}
	}
	return buffer.String()
}

// approach 2 - sort by row *
// import (
// 	"bytes"
// 	"strings"
// )
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]strings.Builder, numRows)
	curRow := 0
	goDown := false
	ans := ""
	for i := 0; i < len(s); i++ {
		rows[curRow].WriteString(s[i:i+1])
		if curRow == 0 || curRow == numRows-1 {
			goDown = !goDown
		}
		switch goDown {
		case true:
			curRow++
		default:
			curRow--
		}
	}
	for _, row := range rows {
		ans += row.String()
	}
	return ans
}