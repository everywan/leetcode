package codes

// Reverse is 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。如 123 => 321

func reverse(num int) int {
	if num == 0 {
		return 0
	}
	var re int
	re = 0
	// 如果业务中有很多以0结尾的数据, 这个还是很有必要的
	for re == 0 {
		re = num % 10
		num = num / 10
	}
	// 正负同理, 不用做区分
	for num != 0 {
		re = re*10 + num%10
		num = num / 10
	}
	// int32最大值的两种表示方法: 0xffffffff, 以及移位表示
	max32 := 1 << 31
	if re > (max32-1) || re < max32*(-1) {
		return 0
	}
	return re
}
