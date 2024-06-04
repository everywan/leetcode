package codes

/*
560. 和为 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。
子数组是数组中元素的连续非空序列。

示例 1：
输入：nums = [1,1,1], k = 2
输出：2
*/

// tag 偏科选手 非最佳解法
// 2024-06-03 start: 18:58 end:19:05
func subarraySum(nums []int, k int) int {
	cnt := 0
	for i, start := range nums {
		sub := k - start
		if sub == 0 {
			cnt += 1
		}
		for j := i + 1; j < len(nums); j++ {
			sub = sub - nums[j]
			if sub == 0 {
				cnt += 1
			}
		}
	}
	return cnt
}

// tag: 前缀和正确, 计算方式有待优化
func subarraySum2(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	cnt := 0
	prefixSum := map[int]int{
		0: nums[0],
	}
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = nums[i] + prefixSum[i-1]
	}
	// map[sum][]int{pos,pos}
	sumPosSet := map[int][]int{}
	for k, v := range prefixSum {
		sumPosSet[v] = append(sumPosSet[v], k)
	}
	for i, sum := range prefixSum {
		if sum == k {
			cnt += 1
		}
		for _, pos := range sumPosSet[sum+k] {
			if i < pos {
				cnt += 1
			}
		}
	}

	return cnt
}
