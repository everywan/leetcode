package codes

/*
49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
*/

// 示例 1:
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

// 2024-07-02 start: 19:19 end:19:23
func groupAnagrams(strs []string) [][]string {
	sets := map[[26]int][]string{}
	for _, str := range strs {
		_set := [26]int{}
		for _, c := range str {
			_set[c-'a']++
		}
		if _, ok := sets[_set]; ok {
			sets[_set] = append(sets[_set], str)
			continue
		}
		sets[_set] = []string{str}
	}
	ans := [][]string{}
	for _, strs := range sets {
		ans = append(ans, strs)
	}
	return ans
}
