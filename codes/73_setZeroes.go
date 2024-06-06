package codes

/*
73. 矩阵置零
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。
请使用 原地 算法。
*/

// 2024-06-05 start: 19:12 end:19:24
func setZeroes(matrix [][]int) {
	// 按行, 找到所有为 0 的元素
	// 将元素所在行变为 0

	var xset, yset []int
	for x, line := range matrix {
		for y, num := range line {
			if num == 0 {
				xset = append(xset, x)
				yset = append(yset, y)
			}
		}
	}

	for _, x := range xset {
		for i := range matrix[x] {
			matrix[x][i] = 0
		}
	}
	for _, y := range yset {
		for i := range matrix {
			matrix[i][y] = 0
		}
	}
}
