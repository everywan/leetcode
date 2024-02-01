package codes

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标.
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
// in: nums = [3,2,1,0,4] out: false

// 思路: 只要不跳到0即可. 按0对数组进行分段, 其逆转顺序必然为 0-1-1|2-1|2|3, 也就是说必须再x步内有0
func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	maxDis := 0
	for i, currentDis := range nums {
		if currentDis > maxDis {
			maxDis = currentDis - 1
			continue
		}
		if currentDis == 0 && maxDis == 0 {
			return i == len(nums)-1
		}
		maxDis = maxDis - 1
	}
	return true
}
