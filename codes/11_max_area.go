package codes

// MaxArea is 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器，且 n 的值至少为 2。

func maxArea(height []int) int {
	area := 0
	for i := 0; i < len(height)-1; i++ {
		var crisis int
		for j := len(height) - 1; j > i; j-- {
			if height[j] > crisis {
				if height[j] >= height[i] {
					_area := (j - i) * height[i]
					if _area > area {
						area = _area
					}
					break
				}
				crisis = height[j]
				_area := (j - i) * crisis
				if _area > area {
					area = _area
				}
			}
		}
	}
	return area
}
