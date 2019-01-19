// package main

// import "fmt"

// func main() {
// 	longestPalindrome("asdasd")
// }

// // 思路: 查找到每个字符所在的位置, 从最长距离依次查找
// // 思路二应该比这个好, 至少表现上更直观. 详细见 t_5_2.go
// func longestPalindrome(s string) string {
// 	if len(s) < 2 {
// 		return s
// 	}
// 	maxLen := 1
// 	maxS := s[0]
// 	all := make(map[rune][]int)
// 	for i, c := range s {
// 		all[c] = append(all[rune(c)], i)
// 	}
// 	fmt.Println(all)
// 	for _, list := range all {
// 		llen := len(list)
// 		if llen < 2 {
// 			continue
// 		}
// 		if list[llen-1]-llen[0] < maxLen {
// 			continue
// 		}

// 		for

// 	}
// 	return s
// }
