package codes

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
*/

// tag 需要回顾
// 2024-06-24 start: 19:33 end:54
func numIslands2(grid [][]byte) int {
	cnt := 0
	for i := range grid {
		for j, value := range grid[i] {
			if value == '0' {
				continue
			}
			cnt++
			numIslands2_dfs(i, j, grid)
		}
	}
	return cnt
}

func numIslands2_dfs(x, y int, grid [][]byte) {
	if grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	if x < len(grid)-1 {
		numIslands2_dfs(x+1, y, grid)
	}
	if y < len(grid[0])-1 {
		numIslands2_dfs(x, y+1, grid)
	}
	if x > 0 {
		numIslands2_dfs(x-1, y, grid)
	}
	if y > 0 {
		numIslands2_dfs(x, y-1, grid)
	}
}
