package codes

/*
53. 最大子数组和
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组中的一个连续部分。

示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/

// 2024-07-03 start: 19:28 end:19:33
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	// 贪心思想
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// 2024-06-24 start: 19:11 end:20:32
func maxSubArray2(nums []int) int {
	return maxSubArray_dfs(nums[1:], nums[0])
}

func maxSubArray_dfs(nums []int, currMax int) int {
	if len(nums) == 0 {
		return currMax
	}
	if len(nums) == 1 {
		if currMax > 0 && nums[0] > 0 {
			return currMax + nums[0]
		}
		if currMax > nums[0] {
			return currMax
		}
		return nums[0]
	}

	// 包含 nums[0] 且包含之前数据的最大值
	max1 := maxSubArray_dfs(nums[1:], currMax+nums[0])
	// 包含 nums[0] 且不包含之前数据的最大值
	max2 := maxSubArray_dfs(nums[1:], nums[0])
	// 不包含当前元素时的最大值
	max3 := maxSubArray_dfs(nums[2:], nums[1])
	if max1 > currMax {
		currMax = max1
	}
	if max2 > currMax {
		currMax = max2
	}
	if max3 > currMax {
		currMax = max3
	}
	return currMax
}
