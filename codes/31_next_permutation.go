package codes

/*
 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

 如果不存在下一个更大的排列，则将数字重新排列成最小的排列

即: 从数组末尾开始, 将每一个遍历的数的集合设为 ends, 如果下一个遍历的数大于ends中的最大值,
那么将ends重新排序后, 替换原位置, 既是结果.

*/

func nextPermutation(nums []int) {
	swap := func(nums []int, i, j int) {
		nums[i] = nums[i] ^ nums[j]
		nums[j] = nums[i] ^ nums[j]
		nums[i] = nums[i] ^ nums[j]
	}

	max := len(nums) - 1
	for i := max - 1; i > -1; i-- {
		if nums[i] >= nums[max] {
			// 在 ends 中, 需要移位, 将 i 移动到末尾, 其他值前移一位腾出空间
			// 如何寻找更好的方法
			tmp := nums[i]
			for j := i; j < max; j++ {
				nums[j] = nums[j+1]
			}
			nums[max] = tmp
		} else if nums[i] < nums[max] {
			for j := i + 1; j < max+1; j++ {
				if nums[j] > nums[i] {
					swap(nums, i, j)
					return
				}
			}
		}
	}
}
