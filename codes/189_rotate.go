package codes

/*
189. 轮转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

// 使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗
*/

// tag 有个更神奇的做法, 可以让 k 次变1次
// 2024-07-03 start: 19:53 end:19:09
func rotate(nums []int, k int) {
	// 1: 创建长度为 k 的数组作为中间数组. 等同于 add 然后截取
	// 2. 循环 k 次, 每一次替换前后值
	k = k % len(nums)
	for i := 0; i < k; i++ {
		temp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
	}
}
