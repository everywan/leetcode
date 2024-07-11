package codes

/*
118. 杨辉三角
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:
输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
*/

// 2024-07-11 start: 20:07 end:19:09
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	ans := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		ans = append(ans, growGenerate(ans[i-1]))
	}
	return ans
}

func growGenerate(row []int) []int {
	newRow := make([]int, 0, len(row)+1)
	for i := 0; i < len(row); i++ {
		value := row[i]
		if i-1 >= 0 {
			value += row[i-1]
		}
		newRow = append(newRow, value)
	}
	newRow = append(newRow, 1)
	return newRow
}
