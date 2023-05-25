package codes

func xorOperation(n int, start int) int {
	result := start
	for i := 1; i < n; i++ {
		result = result ^ (start + 2*i)
	}
	return result
}
