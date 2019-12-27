package codes

import "strconv"

/*
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数
示例
1.     1
2.     11
3.     21
4.     1211
5.     111221
*/

func countAndSay(n int) string {
	result := "1"
	for i := 1; i < n; i++ {
		result = count(result)
	}
	return result
}

func count(s string) string {
	prevPos := 0
	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != s[prevPos] {
			ans = append(ans, []byte(strconv.Itoa(i-prevPos))...)
			ans = append(ans, s[prevPos])
			prevPos = i
		}
	}
	ans = append(ans, []byte(strconv.Itoa(len(s)-prevPos))...)
	ans = append(ans, s[prevPos])
	return string(ans)
}

// 事实证明多次拼接字符串时, fmt.Sprintf 性能不如 byte 等方式
//func count(s string) string {
//	result := ""
//	prevPos := 0
//	for i := 0; i < len(s); i++ {
//		if s[i] != s[prevPos] {
//			result = fmt.Sprintf("%s%d%c", result, i-prevPos, s[prevPos])
//			prevPos = i
//		}
//	}
//	result = fmt.Sprintf("%s%d%c", result, len(s)-prevPos, s[prevPos])
//	return result
//}
