package codes

import (
	"sort"
)

/*
四数之和: 找出 nums 中相加等于 target 的四个数, 不可重复.
*/

// tags: 可改进

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for x1 := 0; x1 < len(nums)-3; x1++ {
		// 这么做是错误的, 因为 nums[x1]+nums[x2] 可能比 nums[x1] 小(nums[x2]<0).
		// 只有 sum+x4 的结果是一定会变小的, sum+ x1/x2/x3 的结果即可能变大, 也可能变小.
		//if nums[x1] > target {
		//	break
		//}
		for x2 := x1 + 1; x2 < len(nums)-2; x2++ {
			sum2 := nums[x1] + nums[x2]
			for x3 := x2 + 1; x3 < len(nums)-1; x3++ {
				sum3 := sum2 + nums[x3]
				// 这一步在 leetcode 中评测中至关重要, 十分影响执行用时.
				// 用于筛选极端情况
				if sum3+nums[x3+1] > target {
					break
				}
				for x4 := len(nums) - 1; x4 > x3; x4-- {
					// 其实可以考虑将 x4 设置为游标式, 根据 sum3 与 target 的比较决定 x4 从哪一端开始遍历
					// 如果 sum3>target, x4 从左端开始. 否则从右端开始
					sum4 := sum3 + nums[x4]
					if sum4 == target {
						result = append(result, []int{nums[x1], nums[x2], nums[x3], nums[x4]})
						x1, x2, x3, x4 = distinctFourSum(x1, x2, x3, x4, nums)
						// result = distinctFourSum(x1, x2, x3, x4, nums, result)
						break
					}
					if sum4 < target {
						break
					}
				}
			}
		}
	}
	return result
}

// 去除重复: 只有当相邻元素相等, 且符合的数组仅含有其中一个时, 会出现.
// 所以, 只要去掉 在 x2 之前, 相同的 x1 就可以了, 注意, 要从最后位开始, 否则清洗不完全
func distinctFourSum(x1, x2, x3, x4 int, nums []int) (int, int, int, int) {
	for x3 < len(nums)-1 && nums[x3] == nums[x3+1] && x4 != x3+1 {
		x3++
	}
	for x2 < len(nums)-2 && nums[x2] == nums[x2+1] && x3 != x2+1 {
		x2++
	}
	for x1 < len(nums)-3 && nums[x1] == nums[x1+1] && x2 != x1+1 {
		x1++
	}
	return x1, x2, x3, x4
}

// 简单暴力清除重复
//func distinctFourSum(x1, x2, x3, x4 int, nums []int, result [][]int) [][]int {
//	for _, item := range result {
//		if nums[x1] == item[0] && nums[x2] == item[1] && nums[x3] == item[2] && nums[x4] == item[3] {
//			return result
//		}
//	}
//	result = append(result, []int{nums[x1], nums[x2], nums[x3], nums[x4]})
//	return result
//}
