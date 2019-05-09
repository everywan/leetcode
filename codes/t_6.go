package codes

import (
	"math"
)

// Convert is  将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
func Convert(s string, numRows int) string {
	return Convert(s, numRows)
}

// 借助方向变量比自己去判断省很多事, 代码也更清晰
func convert(s string, numRows int) string {
	sLen := len(s)
	if sLen <= numRows || numRows == 1 {
		return s
	}
	var arr = make([]string, numRows)
	goDowning := false
	pos := 0
	for _, v := range s {
		arr[pos] = arr[pos] + string(v)
		if pos == 0 || pos == numRows-1 {
			goDowning = !goDowning
		}
		if goDowning {
			pos++
		} else {
			pos--
		}
	}
	result := ""
	for _, v := range arr {
		result += v
	}
	return result
}

func convert2(s string, numRows int) string {
	sLen := len(s)
	if sLen <= numRows || numRows == 1 {
		return s
	}
	var result string
	var arr = make([]string, numRows)
	n := numRows
	for i := 0; i < sLen; i = i + 2*n - 2 {
		for j := 0; j < 2*n-2; j++ {
			if (i + j) >= sLen {
				break
			}
			if j < n {
				arr[j] += string(s[i+j])
			} else {
				arr[2*n-j-2] += string(s[i+j])
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		for _, v := range arr[i] {
			result += string(v)
		}
	}
	return result
}

func convert3(s string, numRows int) string {
	sLen := len(s)
	if sLen <= numRows || numRows == 1 {
		return s
	}
	n := numRows
	sessionLen := int(math.Ceil(float64(sLen) / float64(2*n-2)))
	var rs = make([]byte, sessionLen*2*(n-1))
	for i := 0; i < sLen; i = i + 2*n - 2 {
		t := i / (2*n - 2)
		rs[t] = s[i]
		for j := 1; j < 2*n-2; j++ {
			if (i + j) >= sLen {
				break
			}
			if j < n-1 {
				rs[2*t+sessionLen*(j*2-1)] = s[j+i]
			} else if j == n-1 {
				rs[t+sessionLen*(j*2-1)] = s[j+i]
			} else {
				rs[2*t+sessionLen*(4*n-5-2*j)+1] = s[j+i]
			}
		}
	}
	var result string
	for _, v := range rs {
		if v == 0 {
			continue
		}
		result += string(v)
	}
	return result
}
