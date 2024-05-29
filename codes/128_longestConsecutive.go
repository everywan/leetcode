package codes

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
*/

// 2024-05-29 start: 20:30 end: 21:09
func longestConsecutive(nums []int) int {
	m := map[int]int{}
	max := 0
	for _, num := range nums {
		if _, ok := m[num]; ok {
			continue
		}
		m[num] = 1
		_max := 1
		cntAdd1, ok1 := m[num+1]
		cndSub1, ok2 := m[num-1]
		if ok1 && ok2 {
			_max = cndSub1 + cntAdd1 + 1
			m[num+cntAdd1] = _max
			m[num-cndSub1] = _max
		} else if ok1 {
			m[num] = cntAdd1 + 1
			m[num+cntAdd1] += 1
			_max = m[num]
		} else if ok2 {
			m[num] = cndSub1 + 1
			m[num-cndSub1] += 1
			_max = m[num]
		}
		if _max > max {
			max = _max
		}
	}
	return max
}
