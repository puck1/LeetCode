package main

import (
	"bytes"
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]bytes.Buffer, numRows)
	curRow := 0
	goDown := false
	var ret strings.Builder
	for i := 0; i < len(s); i++ {
		rows[curRow].WriteByte(s[i])
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
		ret.WriteString(row.String())
	}
	return ret.String()
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
	// fmt.Println(longestPalindrome("babab"))
	// fmt.Println(longestPalindrome("b"))
	// fmt.Println(longestPalindrome("ab"))
	// fmt.Println(longestPalindrome("bbb"))
	// fmt.Println(longestPalindrome("bbbb"))
	// fmt.Println(longestPalindrome("cbbd"))
	// fmt.Println(longestPalindrome("abcba"))
}
