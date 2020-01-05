package codes

import (
	"strings"
)

// 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

// 方法一: 找出所有 words 的排列组合, 然后在 s 中查找. 存在浪费的情况. 如 asdf, [a,s,f], asd会遍历两次
// 方法二: 动态规划, 逐词遍历

/* PS: leetcode 的判定算法太傻逼了...
第一次没过, 是因为 没看到 words 长度相等, 写了个任意长度匹配的, 加上没用 map 匹配而是使用
HasPrefix 导致leetcode报内存超时+时间超时
第二次, 改为map匹配, 却忘记了words有重复字符串
第三次, 全都该好了, 还超时(就那个全是a的字符串)... 崩了...
看了看 leetcode 答案, 才发现是限制了 handle 时传入s的长度, 只要传入时判断 s 剩余len< words 总len 即推出,
然后就过了... 我本来想的是这些小细节先不管, 没想到坑到这了..
所以说, 不要太相信 leetcode 的判定算法, leetcode 的判定标准太格式化了.
*/

func findSubstring(s string, words []string) []int {
	f := &FindSubstring{
		s:      s,
		words:  words,
		Result: []int{},
	}
	f.Do()
	return f.Result
}

type FindSubstring struct {
	s      string
	length int
	words  []string
	Result []int
}

func (f *FindSubstring) Do() {
	if len(f.words) < 1 {
		return
	}
	f.length = len(f.words[0])
	totalLen := f.length * len(f.words)
	result := []int{}
	wordMap := make(map[string]int)
	for _, word := range f.words {
		wordMap[word]++
	}
	for i := 0; i < len(f.s)-totalLen+1; i++ {
		if match := f.handle(f.s[i:i+totalLen], "", wordMap, len(f.words)); match != "" {
			result = append(result, i)
		}
	}
	f.Result = result
}

func (f *FindSubstring) handle(s, current string, words map[string]int, rest int) string {
	if rest == 0 {
		return current
	}
	if len(s) < f.length {
		return ""
	}
	word := s[:f.length]
	if exist, ok := words[word]; ok && exist > 0 {
		words[word]--
		result := f.handle(s[f.length:], current+word, words, rest-1)
		words[word]++
		if result != "" {
			return result
		}
	}
	return ""
}

// 无重复word时
func (f *FindSubstring) noRepeatWord(s, current string, words map[string]bool, rest int) string {
	if rest == 0 {
		return current
	}
	if len(s) < f.length {
		return ""
	}
	word := s[:f.length]
	if exist, ok := words[word]; ok && exist {
		words[word] = false
		result := f.noRepeatWord(s[f.length:], current+word, words, rest-1)
		words[word] = true
		if result != "" {
			return result
		}
	}
	return ""
}

// 长度不固定时
func (f *FindSubstring) reorder(words []string, i int) []string {
	// 调换顺序, 将 i 元素放到首位
	tmp := words[i]
	words[i] = words[0]
	words[0] = tmp
	return words
}

func (f *FindSubstring) wordsNotFixLength(s, current string, words []string) string {
	if len(words) == 0 {
		return current
	}
	for j, word := range words {
		if strings.HasPrefix(s, word) {
			_words := f.reorder(words, j)
			if result := f.wordsNotFixLength(s[len(word):], current+word, _words); result != "" {
				return result
			}
		}
	}
	return ""
}
