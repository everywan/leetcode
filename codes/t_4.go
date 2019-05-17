package codes

import (
	"github.com/everywan/leetcode/codes/utils"
)

// FindMedianSortedArrays is 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
// 请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
func FindMedianSortedArrays() {
	utils.ShouldEqual(findMedianSortedArrays([]int{1, 2}, []int{3, 4}), 2.5)
	utils.ShouldEqual(findMedianSortedArrays([]int{1}, []int{3, 4}), 3)
	utils.ShouldEqual(findMedianSortedArrays([]int{1, 3}, []int{2}), 2)
	utils.ShouldEqual(findMedianSortedArrays([]int{}, []int{1}), 1)
	utils.ShouldEqual(findMedianSortedArrays([]int{1}, []int{1}), 1)
	utils.ShouldEqual(findMedianSortedArrays([]int{2, 3, 4}, []int{1}), 2.5)
	utils.ShouldEqual(findMedianSortedArrays([]int{3}, []int{-2, -1}), -1)
}

// 两种思路:
// 思路一: 循环找出前一半较少的数, 然后在这个位置找出中位数. 不过我用这种方法经常被坐标/位置搞混...
// 思路二: 从两个数组的中位数出发, 依次向两边扩展

/*
思路一: 循环, 剔除前50小的数.
循环比较 nums1[i] nums2[j], 保留数据最大和次大的数据(方法findMedianSortedArrays)

对于按照 index 查找到方法而言, i/j 中只有一个是有用的, 另一个时没有用处的(因为存在nums1全部小于nums2某一部分的情况, 导致记录i没有意义).
正是这个问题+索引越界, 导致 arch1FindMedianSortedArrays 的思路越写越恶心. 没有删除, 留下来以后作为警醒吧.

特殊情况:
1. m == 0 || n==0
2. m==n==1
3. 某一队列存在值大于/小于另一队列所有值
*/

/*
教训如下
1. 先完成主要的逻辑, 特殊情况加注释, 留于后续完善. 第一次写就是疲于完善而忘了主要逻辑.
2. 过早地优化是万恶之源.
3. 人脑判断索引界限真是恶心之源.
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	var i, j int
	x, y := 1<<31, 1<<31
	m := len(nums1)
	n := len(nums2)

	// 要求必须循环2次以上, 1th给y赋值, 2th 给x赋值, 所以 m+n>=2
	if m == 0 && n == 0 {
		return result
	}
	// 可以加一些特殊的判断, 如 nums1 全部小于 nums2, 或 m==0||n==0
	if m+n < 2 {
		if m == 0 {
			return float64(nums2[0])
		}
		return float64(nums1[0])
	}
	// 因为要循环到中位数的位置, 所以要 <=
	for z := 0; z <= (m+n)/2; z++ {
		x = y
		if i == m {
			y = nums2[j]
			j++
			continue
		}
		if j == n {
			y = nums1[i]
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			y = nums1[i]
			i++
		} else {
			y = nums2[j]
			j++
		}
	}
	if (m+n)%2 == 0 {
		result = float64(x+y) / 2
	} else {
		result = float64(y)
	}
	return result
}

// 被索引越界恶心的不想写了
func arch1FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	m := len(nums1)
	n := len(nums2)
	// -1 是因为从两个队列取值, 会比从单队列多一个值.
	mid := (m+n)/2 - 1
	p := (m+n)%2 == 0
	if m == 0 || n == 0 {
		_nums := append(nums1, nums2...)
		if len(_nums) == 0 {
			return 0
		}
		if p {
			return float64(_nums[mid]+_nums[mid+1]) / 2
		}
		return float64(_nums[mid+1])
	}
	if m == 1 && n == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}
	var i, j int
	for {
		if i+j == mid {
			if p {
				var _min int
				if i == m-1 {
					_min = min(nums1[i], nums2[j+1])
				} else if j == n-1 {
					_min = min(nums1[i+1], nums2[j])
				} else {
					_min = min(nums1[i+1], nums2[j+1])
				}
				result = float64(max(nums1[i], nums2[j])+_min) / 2
			} else {
				if nums1[i] > nums2[j] {
					if j == n-1 {
						result = float64(nums1[i])
						break
					}
					if nums1[i] < nums2[j+1] {
						result = float64(nums1[i])
					} else {
						result = float64(nums2[j+1])
					}
				} else {
					if i == m-1 {
						result = float64(nums2[j])
						break
					}
					if nums2[j] < nums1[i+1] {
						result = float64(nums2[j])
					} else {
						result = float64(nums1[i+1])
					}
				}
			}
			break
		}
		if nums1[i] < nums2[j] {
			if i == m-1 {
				j = mid - i
				continue
			}
			i++
			continue
		} else {
			if j == n-1 {
				i = mid - j
				continue
			}
			j++
		}
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
