package codes

import (
	"sort"
)

// ThreeSum is 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

// 排序+双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	for x1 := range nums {
		if x1 > 0 && nums[x1] == nums[x1-1] {
			continue
		}
		for x2 := x1 + 1; x2 < len(nums); x2++ {
			if x2-1 > x1 && nums[x2] == nums[x2-1] {
				continue
			}
			target := 0 - nums[x1] - nums[x2]
			x3 := sort.SearchInts(nums[x2+1:], target)
			if x3+x2+1 < len(nums) && nums[x3+x2+1] == target {
				result = append(result, []int{nums[x1], nums[x2], nums[x3]})
				break
			}
		}
	}
	return result
}
