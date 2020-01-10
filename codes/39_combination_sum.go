package codes

import (
	"sort"
)

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，
找出 candidates 中所有可以使数字和为 target 的组合。

1. candidates 中的数字可以无限制重复被选取。
2. 所有数字（包括 target）都是正整数。
3. 解集不能包含重复的组合
*/

/*
有个简单的思路
1. 过滤掉比target大的值
2. 背包问题, 迭代寻找

做出来简单, 但是如果要做好, 用Go的话, 难点在于如何高效利用切片
(不深拷贝的情况下)
*/

func combinationSum(candidates []int, target int) [][]int {
	c := &CombinationSum{
		Ans: make([][]int, 0),
	}
	return c.Do(candidates, target)
}

type CombinationSum struct {
	Ans [][]int
}

func (c *CombinationSum) Do(candidates []int, target int) [][]int {
	candidates = c.filter(candidates, target)
	sort.Sort(sort.IntSlice(candidates))
	c.do(candidates, target, []int{})
	return c.Ans
}

func (c *CombinationSum) do(candidates []int, target int, current []int) {
	for i, x := range candidates {
		if x < target {
			// append 时可能更改之前 append 到 c.Ans 中的值
			c.do(candidates[i:], target-x, append(current, x))
			continue
		} else if x == target {
			current = append(current, target)
			newCurrent := make([]int, len(current))
			copy(newCurrent, current)
			c.Ans = append(c.Ans, newCurrent)
		}
		break
	}
}

// 去除比 target 大的值
func (c *CombinationSum) filter(candidates []int, target int) []int {
	pos := 0
	for i, x := range candidates {
		if x > target {
			continue
		}
		if pos != i {
			candidates[pos] = candidates[i]
		}
		pos++
	}
	return candidates[:pos]
}
