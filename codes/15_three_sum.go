package codes

import (
	"fmt"
	"sort"
)

// ThreeSum is 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

// 排序+双指针
// 这个题蛋疼的是去除重复元素.
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	_nums := sort.IntSlice(nums)
	_nums.Sort()
	res := [][]int{}
	for i := range _nums {
		if i != 0 && _nums[i] == _nums[i-1] {
			continue
		}
		left, right := i+1, len(_nums)-1
		for left < right {
			sum := _nums[right] + _nums[i] + _nums[left]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				res = append(res, []int{_nums[i], _nums[left], _nums[right]})
				for left < right {
					if _nums[left] == _nums[left+1] {
						left++
						continue
					}
					if _nums[right] == _nums[right-1] {
						right--
						continue
					}
					left++
					right--
					break
				}
			}
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
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
