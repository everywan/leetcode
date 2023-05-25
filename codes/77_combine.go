package codes

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。


示例 1：
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
示例 2：
输入：n = 1, k = 1
输出：[[1]]
*/

func combine(n int, k int) [][]int {
	//return combineRun(1, n+1, k, []int{})
	return combineRun(1, n+1, k)
}

func combineForLeetcode(n, k int) [][]int {
	n = n + 1
	var dfs func(int, int) [][]int
	dfs = func(start, k int) [][]int {
		if k == 0 {
			return [][]int{{}}
		}
		res := [][]int{}
		for ; start < n; start++ {
			combines := dfs(start+1, k-1)
			for _, combine := range combines {
				res = append(res, append(combine, start))
			}
		}
		return res
	}
	return dfs(1, k)
}

// 递归时, 构成前缀链
//func combineRun(start, n, k int, preItems []int) [][]int {
//	if k == 0 {
//		return [][]int{preItems}
//	}
//	res := [][]int{}
//	for ; start < n; start++ {
//		// 这里存在 slice 覆盖的问题
//		item := append(preItems, start)
//		_res := combineRun(start+1, n, k-1, item)
//		res = append(res, _res...)
//	}
//	return res
//}

// 递归, 由最下层向上依次构建链子
func combineRun(start, n, k int) [][]int {
	if k == 0 {
		return [][]int{{}}
	}
	res := [][]int{}
	for ; start < n; start++ {
		combines := combineRun(start+1, n, k-1)
		for _, combine := range combines {
			res = append(res, append(combine, start))
		}
	}
	return res
}
