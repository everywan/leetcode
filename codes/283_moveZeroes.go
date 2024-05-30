package codes

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
*/

// 2024-05-30 start: 20:25 end:20:51
func moveZeroes(nums []int) {
	lastPos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastPos] = nums[i]
			lastPos += 1
		}
	}
	for i := lastPos; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 2024-05-30 start: 20:17 end: 20:24
// 读错题了, 不是保持 0 的顺序, 而是非 0 的顺序
// func moveZeroes(nums []int) {
// 	swap := func(pos1, pos2 int, nums []int) {
// 		temp := nums[pos1]
// 		nums[pos1] = nums[pos2]
// 		nums[pos2] = temp
// 	}
// 	rightPos := len(nums) - 1
// 	for i := 0; i < rightPos; i++ {
// 		if nums[i] == 0 {
// 			for nums[rightPos] == 0 && rightPos > i {
// 				rightPos -= 1
// 			}
// 			swap(i, rightPos, nums)
// 			rightPos -= 1
// 		}
// 	}
// }
