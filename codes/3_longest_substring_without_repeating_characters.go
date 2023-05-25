package codes

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
start:10:12 end_code:10:20 end_all:10:34
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	hasAppeal := map[rune]struct{}{}
	maxLength := 0
	for start := range s {
		length := len(s) - start
		hasAppeal[rune(s[start])] = struct{}{}
		for i := start + 1; i < len(s); i++ {
			e := rune(s[i])
			if _, ok := hasAppeal[e]; ok {
				length = i - start
				break
			}
			hasAppeal[e] = struct{}{}
		}
		if maxLength < length {
			maxLength = length
		}
		hasAppeal = map[rune]struct{}{}
	}
	return maxLength
}
