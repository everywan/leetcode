package codes

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

// 示例 1:
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

// 方案一: 使用map加速两两比较
// 方案二: 所有字符串排序, 排序后借助 hash 计算即可.
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	uset := map[int]map[rune]int{}
	for _, str := range strs {
		m := map[rune]int{}
		for _, c := range str {
			m[c] += 1
		}
		isSame := false
		for i, m2 := range uset {
			if len(m) != len(m2) {
				continue
			}
			s2 := true
			for c, cnt := range m {
				if m2[c] != cnt {
					s2 = false
					break
				}
			}
			if !s2 {
				continue
			}
			isSame = true
			if isSame {
				result[i] = append(result[i], str)
				break
			}
			break
		}
		if !isSame {
			result = append(result, []string{str})
			uset[len(result)-1] = m
		}
	}
	return result
}
