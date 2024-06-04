package codes

/*
74. 搜索二维矩阵
给你一个满足下述两条属性的 m x n 整数矩阵：
1. 每行中的整数从左到右按非严格递增顺序排列。
2. 每行的第一个整数大于前一行的最后一个整数。
给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
*/

// 2024-06-04 start: 17:45 end:18:20 一遍过
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	// 第一次搜索终止条件: r[0]>target>l[0] 或 == target
	l, r := 0, len(matrix)-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[mid][0] > target {
			r = mid - 1
		} else if matrix[mid][0] < target {
			l = mid + 1
		} else {
			return true
		}
	}

	// 优化点: 二分查找, 最后必然会是 mid==l==r 对比 target.
	// 当 s[mid]>target 时, r--; 当 s[mid]<target 时, l++
	// 这两种情况下, l 是满足 >target 最小的位置. 即 sort.Search() 函数.
	if l == 0 {
		return false
	}
	nums := matrix[l-1]

	l, r = 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}
