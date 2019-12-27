package codes

// 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
// 注意, int 包括负数

func divide(dividend int, divisor int) int {
	if divisor == -1 && dividend == (-1*(1<<31)) {
		return 1<<31 - 1
	}
	isMinus := false
	if dividend < 0 {
		isMinus = !isMinus
		dividend = -dividend
	}
	if divisor < 0 {
		isMinus = !isMinus
		divisor = -divisor
	}
	i := 0
	for ; dividend >= divisor; dividend -= divisor {
		i += 1
	}
	if isMinus {
		return -i
	}
	return i
}

// 方法一: 使用加法, 循环叠加, 对于差别特别大的数, 可以采用倍增法. 具体看实际数据情况决定, 这里就不写了.
// 方法二: 考虑位运算: 经提示, 解法如下
/*
	1. 首先, 转换为二进制
	2. 通过列竖式除法思想, 移位计算: 每次移动 divisor 的位数
*/
