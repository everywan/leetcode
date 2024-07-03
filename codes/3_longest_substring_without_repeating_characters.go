package codes

/*
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

1. start:10:12 end_code:10:20 end_all:10:34
1. start:18:06 end:18:24 debug:18:29 全部优化: 18:54
*/

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	set := map[byte]int{}
	left := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if pos, ok := set[c]; ok && pos >= left {
			_maxLen := i - left
			if maxLen < _maxLen {
				maxLen = _maxLen
			}
			left = set[c] + 1
		}
		set[c] = i
	}
	if left != len(s) {
		if maxLen < len(s)-left {
			maxLen = len(s) - left
		}
	}
	return maxLen
}

// 滑动窗口解法
func lengthOfLongestSubstringV2(s string) int {
	maxLen := 0
	set := [256]*struct{}{}
	l, r := 0, 0 // 滑动窗口左右位置
	for r < len(s) {
		for ; r < len(s); r++ {
			if set[s[r]] == nil {
				set[s[r]] = &struct{}{}
				continue
			}
			// 有重复值时, 移动左指针到重复地点
			for ; l < r; l++ {
				if set[s[l]] == set[s[r]] {
					break
				}
				set[s[l]] = nil
			}
			break
		}
		if maxLen < r-l {
			maxLen = r - l
		}
		set[s[l]] = nil
		l++
	}
	return maxLen
}
