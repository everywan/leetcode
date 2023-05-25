package codes

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
1. 字符串不以 0 开头
2. 长度小于 110
3. 不可使用标准库转换为 int
*/

func multiply(num1 string, num2 string) string {
	result := []int{}
	for i := len(num1) - 1; i > -1; i-- {
		index := 0
		for j := len(num2) - 1; j > -1; j-- {
			x := int((num1[i] - '0') * (num2[j] - '0'))
			result[index] = result[index] + x%10
			result[index+1] = result[index+1] + x/10
			index++
		}
	}

	return ""
}

// index*index, 结果在 index&index+1位上, 然后把每一位的想加

// i[0]*j[0] => r[1]r[0]
// i[1]*j[0] => r[2]r[1]
// i[2]*j[0] => r[3]r[2]
//
// i[0]*j[1] => r[2]r[1]
// i[0]*j[2] => r[3]r[2]
