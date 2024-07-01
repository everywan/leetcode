package codes

/*
438. 找到字符串中所有字母异位词
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
*/

// 2024-07-01 start: 18:52 end:20:10
func findAnagrams(s string, p string) []int {
	slen, plen := len(s), len(p)
	if slen < plen {
		return []int{}
	}

	sCount, pCount := [26]int{}, [26]int{}
	for i, v := range p {
		sCount[s[i]-'a']++
		pCount[v-'a']++
	}

	result := []int{}
	if sCount == pCount {
		result = append(result, 0)
	}

	for i, v := range s[:slen-plen] {
		sCount[v-'a']--
		sCount[s[i+plen]-'a']++
		if sCount == pCount {
			result = append(result, i+1)
		}
	}
	return result
}
