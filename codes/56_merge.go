package codes

import "sort"

/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

// 2024-06-05 start: 18:47 end:19:05 一遍过
func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if prev[1] >= curr[0] {
			if curr[1] <= prev[1] {
				continue
			}
			prev = []int{prev[0], curr[1]}
			continue
		}
		result = append(result, prev)
		prev = curr
	}
	result = append(result, prev)
	return result
}
