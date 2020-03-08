package others

import (
	"fmt"
	"strconv"
)

// tag: mobile,

/*
给定整数 n, 求符合如下要求的正整数m
1. m是符合要求的最小正整数
2. 整数m的每一位相乘, 结果是 n. 如 m=455, n=4*5*5
3. n>9
*/

func digitProduct(n int) int {
	return new(DigitProduct).Do(n)
}

/*
解法:
	刚开始我陷入了 使用质数的错误路线. 使用质数并没有错, 但是其分析路径太过复杂了.
正确的思路是, 首先, 如果n有约数9, 则9必在m的末尾. 假设现在 n 没有约数9, m 后添加了相应的约数.
	同理, 约去所有的8,7,6,5,4,3,2. 最后, 得到的数要么是1, 即所有约数都被找到,
	或者得到一个不能被 2/3/5/7 整除的质数且大于9, 返回 -1.

其实, 本质上是 n = 2^a + 3^b +5^c +7^d, 然后将各质数相乘, 以最大值的方式呈现,如 [2,2,3]
	应当呈现为 [2,6] 而不是[4,3]. 然后将其乘积按大小排列. 即上述方法, 从9约到2.
*/

type DigitProduct struct{}

func (dp *DigitProduct) Do(n int) int {
	result := dp.findDigit(n, "")
	m, _ := strconv.ParseInt(result, 10, 32)
	return int(m)
}

func (dp *DigitProduct) findDigit(n int, result string) string {
	if n == 1 {
		return result
	}
	for i := 9; i > 1; i-- {
		if n%i == 0 {
			return dp.findDigit(n/i, fmt.Sprintf("%d%s", i, result))
		}
	}
	return "-1"
}
