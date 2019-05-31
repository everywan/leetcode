package threesumclosest

import (
	"sort"
)

// ThreeSumClosest is 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
// 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
// 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

func threeSumClosest(nums []int, target int) int {
	sort.Sort(sort.IntSlice(nums))
	if len(nums) < 3 {
		return target
	}
	var result = 1 << 31

	for i := 0; i < len(nums)-2; i++ {
		target2 := target - nums[i]
		j, z := i+1, len(nums)-1
		for j < z {
			temp := target2 - nums[j] - nums[z]
			if temp == 0 {
				return target
			}
			if abs(temp) < abs(result) {
				result = temp
			}
			if temp > 0 {
				j++
			} else {
				z--
			}
		}
	}
	return target - result
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return ^(a - 1)
}
