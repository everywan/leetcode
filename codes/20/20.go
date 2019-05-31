package isvalid

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
*/

/*
思路其实很简单: 将括号的左边入栈, 匹配到非左边(即右边)括号时, 从栈中取出一个元素比较, 是否符合要求.
有两个特殊情况需要考虑:
1. 当 s 为 "[" 这种只有左边的括号时
2. 当 s 为 "]" 这种只有右边的括号时
*/
func isValid(s string) bool {
	var stack []rune
	stack = make([]rune, len(s))
	var i int
	i = -1
	push := func(stack []rune, c rune) {
		i++
		stack[i] = c
	}
	pop := func(stack []rune) rune {
		if i < 0 {
			return '0'
		}
		i--
		return stack[i+1]
	}
	isEmpty := func(stack []rune) bool {
		if i < 0 {
			return true
		}
		return false
	}
	for _, c := range s {
		switch c {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			push(stack, c)
		case ')':
			if pop(stack) != '(' {
				return false
			}
		case ']':
			if pop(stack) != '[' {
				return false
			}
		case '}':
			if pop(stack) != '{' {
				return false
			}
		}
	}
	return isEmpty(stack)
}
