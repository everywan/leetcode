package codes

/*
54. 螺旋矩阵
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/

// 2024-06-17 start: 19:38 end: 19:58 开始调试 20:14 完全通过
func spiralOrder(matrix [][]int) []int {
	nums := []int{}
	loopCnt := 0
	for {
		_nums := _spiralOrder(matrix, loopCnt)
		if _nums == nil {
			return nums
		}
		loopCnt += 1
		nums = append(nums, _nums...)
	}
}

func _spiralOrder(matrix [][]int, loopCnt int) []int {
	ystart, yend := loopCnt, len(matrix)-loopCnt
	if ystart >= yend || yend <= 0 {
		return nil
	}
	xstart, xend := loopCnt, len(matrix[0])-loopCnt
	if xstart >= xend || xend <= 0 {
		return nil
	}
	nums := []int{}
	for i := xstart; i < xend; i++ {
		nums = append(nums, matrix[loopCnt][i])
	}
	for j := ystart + 1; j < yend; j++ {
		nums = append(nums, matrix[j][xend-1])
	}
	if ystart+1 >= yend {
		return nums
	}
	for i := xend - 2; i > xstart; i-- {
		nums = append(nums, matrix[yend-1][i])
	}
	if xend-2 < xstart {
		return nums
	}
	for j := yend - 1; j > ystart; j-- {
		nums = append(nums, matrix[j][xstart])
	}
	return nums
}
