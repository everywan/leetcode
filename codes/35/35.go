package searchinsert

// 给定一个排序数组和一个目标值, 在数组中找到目标值, 并返回其索引. 如果目标值不存在于数组中, 返回它将会被按顺序插入的位置.
// 你可以假设数组中无重复元素.

func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] < target {
			continue
		}
		return i
	}
	if len(nums) == 0 {
		return 0
	}
	return len(nums)
}

// 中文版的 leetcode 性能计算貌似有问题...
// 同样的代码, 跑几次结果就不一样..
