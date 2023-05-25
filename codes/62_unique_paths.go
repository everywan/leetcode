package codes

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？

链接：https://leetcode.cn/problems/unique-paths
start: 0944 done: 0954
*/

// 思路: 每次只能 m-1 or n-1, 直到 m or n 有一个为0. 当两个都为0时, 路径需要返回, 否则返回 nil
func uniquePaths(m int, n int) int {
	if m == 0 && n == 0 {
		return 0
	}
	return uniquePathsUnit(m-1, n-1)
}

func uniquePathsUnit(m int, n int) int {
	if m == 0 && n == 0 {
		return 1
	}
	cnt := 0
	if m > 0 {
		cnt += uniquePathsUnit(m-1, n)
	}
	if n > 0 {
		cnt += uniquePathsUnit(m, n-1)
	}
	return cnt
}
