package threesum

import (
	"fmt"
	"sort"
)

// ThreeSum is 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

// 主要思路:
// 1. 排序
// 2. 从一端开始遍历
// 3. 从两端各取一个值, 比较 三者之和与0, 从而确定下次移动哪一个游标

// 两个方法中, 都对某些特殊情况做了处理, 如 [0,0,0], [1,2,3]等

// 不同数据集下步骤三采取不同的策略会有不同的影响.
// 当分布比较均匀时, 采取两端依次渐进较好(见后续的优化方法)
// 当 n[i]+n[z]<0 时, 采用第一种方法, 固定iz, j 从大端依次递减更好

func threeSum(nums []int) [][]int {
	_len := len(nums)
	if _len < 3 {
		return nil
	}
	_nums := sort.IntSlice(nums)
	_nums.Sort()
	if _nums[0] > -1 || _nums[_len-1] < 1 {
		if _nums[2] == 0 {
			return [][]int{[]int{0, 0, 0}}
		}
		return nil
	}

	result := [][]int{}
	zeroIdx := _nums.Search(0)
	hasZero := false
	fmt.Println(zeroIdx, _nums)
	if _nums[zeroIdx] == 0 {
		if zeroIdx+3 < _len && _nums[zeroIdx+2] == 0 {
			result = append(result, []int{0, 0, 0})
		}
		hasZero = true
	}
	for i := 0; i < zeroIdx; i++ {
		if i > 0 && _nums[i] == _nums[i-1] {
			continue
		}
		for j := _len - 1; j > zeroIdx-1; j-- {
			if j < _len-1 && _nums[j] == _nums[j+1] {
				continue
			}
			sum := _nums[i] + _nums[j]
			if sum > 0 {
				for z := i + 1; z < zeroIdx; z++ {
					_sum := sum + _nums[z]
					if _sum > 0 {
						break
					} else if _sum == 0 {
						result = append(result, []int{_nums[i], _nums[j], _nums[z]})
						break
					}
				}
			} else if sum < 0 {
				for z := j - 1; z > zeroIdx-1; z-- {
					_sum := sum + _nums[z]
					if _sum < 0 {
						break
					} else if _sum == 0 {
						result = append(result, []int{_nums[i], _nums[j], _nums[z]})
						break
					}
				}
			} else {
				if hasZero {
					result = append(result, []int{_nums[i], _nums[j], 0})
				}
			}
		}
	}
	return result
}

// 主要思路和刚开始的相同, 不同的是 采用 双端快排 的思想, 固定 i 之后, 从两端依次变化 j/z

func threeSumOptimize(nums []int) [][]int {
	_len := len(nums)
	_nums := sort.IntSlice(nums)
	_nums.Sort()
	result := [][]int{}
	for i := 0; i < _len-2; i++ {
		if nums[i] > -1 {
			if nums[i+2] == 0 {
				result = append(result, []int{0, 0, 0})
			}
			return result
		}
		if i == 0 || _nums[i] != _nums[i-1] {
			z := _len - 1
			if _nums[i]+_nums[z]+_nums[z-1] < 0 {
				continue
			}
			sumShouldBe := 0 - _nums[i]
			j := i + 1
			for j < z {
				sum := _nums[j] + _nums[z]
				if sum > sumShouldBe {
					z--
				} else if sum < sumShouldBe {
					j++
				} else {
					result = append(result, []int{_nums[i], _nums[j], _nums[z]})
					j++
					z--
					for j < z {
						if _nums[j] == _nums[j-1] {
							j++
							continue
						}
						if _nums[z] == _nums[z+1] {
							z--
							continue
						}
						break
					}
				}
			}
		}
	}
	return result
}
