package codes

// 全排列, 可能有重复数字负数
func permute(nums []int) [][]int {
	// 方案一: 假设 k 个的排列是 [][], 则 k+1 个是在每个数组的每个位置插入第 k 个元素
	// 		问题: 去重不好做
	// 方案二: 插位置. 第一个位置每个字符放一遍, 第二个位置将剩下的字符放一遍, 以此类推.
	// 如果某个位置两次放的字符一致, 则跳过.

	if len(nums) < 1 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums}
	}
	result := [][]int{}
	prev := -100000
	for i, m := range nums {
		if prev == m {
			continue
		}
		prev = m
		inNew := permuteDeepcopy(nums, i)
		for _, arrItem := range permute(inNew) {
			result = append(result, append([]int{m}, arrItem...))
		}
	}
	return result
}

func permuteDeepcopy(in []int, n int) []int {
	out := make([]int, len(in)-1)
	copy(out[:n], in[:n])
	if n < len(in)-1 {
		copy(out[n:len(in)-1], in[n+1:])
	}
	return out
}
