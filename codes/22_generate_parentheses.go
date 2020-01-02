package codes

// 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
// 如 n=2, 表示有两对括号, 给出其所有有效组合

// 难点在于, 组织括号最佳的方式是什么, 即无任何计算浪费的方式
// 方案就是 动态规划: 任何时刻, 只要保证 ) 出现的次数小于等于 (, 即可保证括号有效.

func generateParenthesis(n int) []string {
	result := generateParenthesisUtil(0, 0, n, "")
	return result
}

func generateParenthesisUtil(left, right, n int, current string) []string {
	if left < n {
		if left > right {
			// left>right && <n 时, 可以随意添加
			c1 := generateParenthesisUtil(left+1, right, n, current+"(")
			c2 := generateParenthesisUtil(left, right+1, n, current+")")
			result := append(c1, c2...)
			return result
		} else {
			// left==right 时, 只能添加 (
			for ; right < left; right++ {
				current += ")"
			}
			return generateParenthesisUtil(left+1, right, n, current+"(")
		}
	} else {
		// 结束时, 补全 )
		for ; right < n; right++ {
			current += ")"
		}
		return []string{current}
	}
}
