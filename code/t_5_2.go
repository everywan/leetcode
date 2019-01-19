package main

import "fmt"

func main() {
	s := "cna"
	fmt.Println(longestPalindrome(s))
}

var maxLen int
var re string

// 思路: 从中间开始, 依次向两边扩展查询, 如果失败则递归查询两边的值. 与上一个不同的是, 1 是从最大向最小查询, 这个是从最小递归想最大查询
// 没有特别优化, 其中有些地方可以提取出公共函数, 暂时这样吧, 小修小该就不做了.
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	re = s[0:1]
	maxLen = 1
	re = findAll(s)
	if re == s {
		return s
	}
	for i := 1; i < len(s)-1; i++ {
		_re := findAll(s[i:len(s)])
		if len(_re) == len(s)-i {
			return _re
		}
		_re = findAll(s[0 : len(s)-i])
		if len(_re) == len(s)-i {
			return _re
		}
	}
	return re
}

func findAll(s string) string {
	length := len(s)
	if length%2 == 0 {
		return findX(s, length/2-1, length/2)
	}
	return findX(s, length/2-1, length/2+1)
}

func findX(s string, l, r int) string {
	length := len(s)
	if s[l] == s[r] {
		if l > 0 && r < length-1 {
			return findX(s, l-1, r+1)
		}
		_maxLen := r - l + 1
		if _maxLen > maxLen {
			maxLen = _maxLen
			re = s[l : r+1]
		}
		return re
	}
	_maxLen := r - l - 1
	if _maxLen > maxLen {
		maxLen = _maxLen
		re = s[l+1 : r]
	}
	return re
}
