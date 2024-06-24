package codes

import (
	"fmt"
)

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
// 2024-06-20 start: 19:22 end:
func numIslands(grid [][]byte) int {
	set := &disjointSet_200{
		record: map[string]string{},
		sets:   map[string][]string{},
	}
	for x, line := range grid {
		for y := range line {
			if grid[x][y] == '0' {
				continue
			}
			set.Add(x, y)
		}
	}
	return len(set.sets)
}

type disjointSet_200 struct {
	sets   map[string][]string
	record map[string]string
}

// 找到点所属的集合
func (dset *disjointSet_200) find(x, y int) []string {
	keys := []string{}
	if x > 0 {
		keys = append(keys, fmt.Sprintf("%d,%d", x-1, y))
	}
	if y > 0 {
		keys = append(keys, fmt.Sprintf("%d,%d", x, y-1))
	}
	setIDs := map[string]struct{}{}
	for _, key := range keys {
		if setID, ok := dset.record[key]; ok {
			setIDs[setID] = struct{}{}
		}
	}
	_setIDs := []string{}
	for id := range setIDs {
		_setIDs = append(_setIDs, id)
	}
	return _setIDs
}

func (dset *disjointSet_200) Add(x, y int) {
	setIDs := dset.find(x, y)

	key := fmt.Sprintf("%d,%d", x, y)
	if len(setIDs) == 0 {
		dset.record[key] = key
		dset.sets[key] = []string{key}
		return
	}
	newSetID := setIDs[0]
	dset.record[key] = newSetID
	dset.sets[newSetID] = append(dset.sets[newSetID], key)
	if len(setIDs) > 1 {
		dset.Merge(setIDs)
	}
}

// 合并多个集合
func (dset *disjointSet_200) Merge(setIDs []string) {
	newSetID := setIDs[0]
	for i := 1; i < len(setIDs); i++ {
		setID := setIDs[i]
		dset.sets[newSetID] = append(dset.sets[newSetID], dset.sets[setID]...)
		delete(dset.sets, setID)
	}
	for _, key := range dset.sets[newSetID] {
		dset.record[key] = newSetID
	}
}
