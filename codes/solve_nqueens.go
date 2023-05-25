package codes

import "math"

/*
设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。
链接：https://leetcode.cn/problems/eight-queens-lcci

start: 0957 begin_write: 1010 test_1:1040 done_1:1053
思路一: 当 (a,b) 被选择后, x=a|y=b|(abs(x-a)=abs(y-b)) 都不符合要求. 存储这样的一系列 ab, 直到能走完全程即可
*/
func solveNQueens(n int) [][][2]int {
	points := [][2]int{}
	return sloveNQueensSubFunc(0, n, points)
}

func sloveNQueensSubFunc(i, n int, points [][2]int) [][][2]int {
	checkQueenRule := func(i, j int, points [][2]int) bool {
		for _, point := range points {
			if i == point[0] || j == point[1] || math.Abs(float64(i-point[0])) == math.Abs(float64(j-point[1])) {
				return true
			}
		}
		return false
	}
	arrayCopy := func(points [][2]int) [][2]int {
		out := make([][2]int, len(points))
		for i := range points {
			out[i] = points[i]
		}
		return out
	}
	if i == n {
		return [][][2]int{points}
	}
	result := [][][2]int{}
	for j := 0; j < n; j++ {
		if !checkQueenRule(i, j, points) {
			println(i, j)
			newPoints := append(arrayCopy(points), [2]int{i, j})
			result = append(result, sloveNQueensSubFunc(i+1, n, newPoints)...)
		}
	}
	return result
}
