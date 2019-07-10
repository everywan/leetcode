package strStr

//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

func strStr(haystack string, needle string) int {
	if haystack == "" && needle == "" {
		return 0
	}
	lenNeedle := len(needle)
	for i := 0; i <= len(haystack)-lenNeedle; i++ {
		if haystack[i:i+lenNeedle] == needle {
			return i
		}
	}
	return -1
}
