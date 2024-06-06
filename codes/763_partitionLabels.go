package codes

/*
763. 划分字母区间
给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
返回一个表示每个字符串片段的长度的列表。


示例 1：
输入：s = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca"、"defegde"、"hijhklij" 。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 这样的划分是错误的，因为划分的片段数较少。

示例 2：
输入：s = "eccbbbbdec"
输出：[10]
*/

// tag: 不简洁-已调整
// 2024-06-06 start: 18:48 end:19:14 一遍过
func partitionLabels(s string) []int {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	result := []int{}
	start, end := 0, 0
	for i, c := range s {
		if end < lastPos[c-'a'] {
			end = lastPos[c-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	return result
}
