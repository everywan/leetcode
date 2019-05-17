package codes

import (
	"github.com/everywan/leetcode/codes/utils"
)

// MyAtoi 实现一个 atoi 函数，使其能将字符串转换成整数
func MyAtoi() {
	utils.ShouldEqual(myAtoi("1"), 1)
	utils.ShouldEqual(myAtoi("-1"), -1)
	utils.ShouldEqual(myAtoi("   -1"), -1)
	utils.ShouldEqual(myAtoi("123 12a"), 123)
	utils.ShouldEqual(myAtoi("123 a"), 123)
	utils.ShouldEqual(myAtoi("a 123"), 0)
	utils.ShouldEqual(myAtoi("-2147483649"), -1*(1<<31))
	utils.ShouldEqual(myAtoi("-91283472332"), -2147483648)
	utils.ShouldEqual(myAtoi("2147483800"), 2147483647)
}

// 陷于特殊情况的判断, 无法自拔
// 忙于特殊情况的判断, 一点也不好玩

func myAtoi(str string) int {
	var result int32
	var minusTag bool
	for i, c := range str {
		switch c {
		case ' ':
			if i == 0 || str[i-1] == ' ' {
				continue
			}
			break
		case '-':
			if i == 0 || str[i-1] == ' ' {
				minusTag = true
				continue
			}
			break
		case '+':
			if i == 0 || str[i-1] == ' ' {
				minusTag = false
				continue
			}
			break
		default:
			if c < '0' || c > '9' {
				break
			}
			// 保证此次循环中不会溢出
			if result < (1<<31-1)/10-9 {
				result = result*10 + c - '0'
				continue
			}
			// 当此次循环可能溢出时
			var temp int32
			if minusTag {
				temp = (1 << 31) * -1
			} else {
				temp = 1<<31 - 1
			}
			// 当result比 (1<<31)/10 还大时, 自动忽略个位数, 取最大值. 或者 如果后续还有数字
			if result > (1<<31)/10 || (i < len(str)-1 && (str[i+1] >= '0' || str[i+1] <= '9')) {
				result = temp
				break
			}
			result = temp
			// 如果后续没有数字
			if minusTag && c < '8' {
				result = result + 8 - (c - '0')
			} else if !minusTag && c < '7' {
				result = result - 7 + c - '0'
			}
		}
		break
	}
	if minusTag && result > 0 {
		result = result * -1
	}
	return int(result)
}
