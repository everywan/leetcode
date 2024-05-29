package codes

// 实现 pow(x, n) ，即计算 x 的整数 n 次幂函数（即，xn ）。
func myPow(x float64, n int) float64 {
	// 方法: x*x => (x*x)*(x*x) => ...
	// 第一次 x^2, 第二次 x^4, 第三次 x^8, ...
	if x == 0 || x == 1 {
		return x
	}
	if n == 0 {
		return 1
	}
	for ; n >= 1; n = n >> 1 {
		x = x * x
	}
	return x
}
