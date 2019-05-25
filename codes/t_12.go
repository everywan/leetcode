package codes

import (
	"strings"

	"github.com/everywan/leetcode/codes/utils"
)

const (
	RomanI = "I"
	// 5
	RomanV = "V"
	// 10
	RomanX = "X"
	// 50
	RomanL = "L"
	// 100
	RomanC = "C"
	// 500
	RomanD = "D"
	// 1000
	RomanM = "M"
)

// 这题太无聊了, 了解转换规则就需要很多时间... 不写了
func IntToRoman() {
	utils.ShouldEqual(intToRoman(1), RomanI)
	utils.ShouldEqual(intToRoman(9), "IX")
	utils.ShouldEqual(intToRoman(18), "XVIII")
	utils.ShouldEqual(intToRoman(19), "IXX")
	utils.ShouldEqual(intToRoman(1994), "MCMXCIV")
}

// 思路:
// 1. 从最大到最小依次合并所有值, 如 9 == IIIIIIIII == VIIII
// 2. VLD 连续出现要少于两次, 否则合并为 XCM
// 3. UXCM 连续出现要少于四次, 否则合并为 N_VLD 形式
func intToRoman(num int) string {
	var result string
	for num > 999 {
		num -= 1000
		result += RomanM
	}
	for num > 499 {
		num -= 500
		result += RomanD
	}
	for num > 99 {
		num -= 100
		result += RomanC
	}
	for num > 49 {
		num -= 50
		result += RomanL
	}
	for num > 9 {
		num -= 10
		result += RomanX
	}
	for num > 4 {
		num -= 5
		result += RomanV
	}
	for num > 0 {
		num--
		result += "I"
	}
	var mergeIXCM = func(s, old, _new, prefix string) string {
		if strings.Contains(s, old) {
			result := strings.Replace(s, old, _new, 1)
			result = prefix + result
			return result
		}
		return s
	}
	result = mergeIXCM(result, "IIII", "V", "I")
	result = strings.Replace(result, "VV", "X", 1)
	result = mergeIXCM(result, "XXXX", "L", "X")
	result = strings.Replace(result, "LL", "C", 1)
	result = mergeIXCM(result, "CCCC", "D", "C")
	result = strings.Replace(result, "DD", "M", 1)
	return result
}
