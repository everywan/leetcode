package codes

// LongestCommonPrefix is 编写一个函数来查找字符串数组中的最长公共前缀。
func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(s) > len(strs[i]) {
			s = s[0:len(strs[i])]
		}
		for j := 0; j < len(s); j++ {
			if s[j] != strs[i][j] {
				s = s[0:j]
				break
			}
		}
	}
	return s
}
