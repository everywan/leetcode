// 看了官方 4ms 的解法, 学习了在 re==0 时作区分. 但是最后的结果仍不是4ms... 不知道为什么.. 难道是移位比16进制慢了? 懒得捣鼓了, 就这样吧
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

