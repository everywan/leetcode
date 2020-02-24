package codes

import (
	"sort"
)

/*
给定一个数组 candidates 和一个目标数 target,
找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合
*/

func combinationSum2(candidates []int, target int) [][]int {
	c := &CombinationSum2{
		Ans: make([][]int, 0),
	}
	return c.Do(candidates, target)
}

type CombinationSum2 struct {
	Ans    [][]int
	repeat map[int]int
}

func (c *CombinationSum2) Do(candidates []int, target int) [][]int {
	candidates = c.filter(candidates, target)
	sort.Sort(sort.IntSlice(candidates))
	c.calRepeat(candidates)
	c.do(candidates, target, []int{})
	return c.Ans
}

func (c *CombinationSum2) do(candidates []int, target int, current []int) {
	for i, x := range candidates {
		if x < target {
			if !c.isRepeat(x, current, candidates[i:]) {
				c.do(candidates[i+1:], target-x, append(current, x))
			}
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

func (c *CombinationSum2) isRepeat(x int, current, candidates []int) bool {
	// 当有n个重复元素时,
	// 当扫描到第 k 个重复元素时, 必须保证 current 中出现了 k-1 个重复元素
	if _, ok := c.repeat[x]; !ok {
		return false
	}
	sum := 0
	for _, e := range candidates {
		if x == e {
			sum++
			continue
		}
		break
	}
	for i := len(current) - 1; i > -1; i-- {
		if x == current[i] {
			sum++
			continue
		}
		break
	}
	if sum == c.repeat[x] {
		return false
	}
	return true
}

// 去除比 target 大的值
func (c *CombinationSum2) filter(candidates []int, target int) []int {
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

func (c *CombinationSum2) calRepeat(candidates []int) {
	c.repeat = make(map[int]int, 0)
	for _, x := range candidates {
		if _, ok := c.repeat[x]; ok {
			c.repeat[x] += 1
			continue
		}
		c.repeat[x] = 1
	}
}
