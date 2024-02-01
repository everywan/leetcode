package codes

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		current := strs[i]
		if len(prefix) < len(current) {
			current = current[:len(prefix)]
		} else if len(prefix) > len(current) {
			prefix = prefix[:len(current)]
		}
		for j := 0; j < len(prefix); j++ {
			if prefix[j] != current[j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}
