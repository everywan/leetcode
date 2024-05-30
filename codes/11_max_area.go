package codes

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
*/

// tag: 非最佳解法
// 备注: 有一个很奇妙的贪心算法, 需要你想到规律
// 2024-05-30 start: 20:53 end: 21:09
func maxArea(height []int) int {
	maxValue := 0
	for posl, hl := range height {
		maxHr := 0
		for posr := len(height) - 1; posr > posl; posr-- {
			if height[posr] <= maxHr {
				continue
			}
			maxHr = height[posr]
			if maxHr >= hl {
				_maxValue := hl * (posr - posl)
				if _maxValue > maxValue {
					maxValue = _maxValue
				}
				break
			}
			_maxValue := maxHr * (posr - posl)
			if _maxValue > maxValue {
				maxValue = _maxValue
			}
		}
	}
	return maxValue
}

// func maxArea(height []int) int {
// 	area := 0
// 	for i := 0; i < len(height)-1; i++ {
// 		var crisis int
// 		for j := len(height) - 1; j > i; j-- {
// 			if height[j] > crisis {
// 				if height[j] >= height[i] {
// 					_area := (j - i) * height[i]
// 					if _area > area {
// 						area = _area
// 					}
// 					break
// 				}
// 				crisis = height[j]
// 				_area := (j - i) * crisis
// 				if _area > area {
// 					area = _area
// 				}
// 			}
// 		}
// 	}
// 	return area
// }
