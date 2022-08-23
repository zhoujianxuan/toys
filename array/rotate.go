package array

func rotate(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l/2; i++ {
		matrix[i], matrix[l-i-1] = matrix[l-i-1], matrix[i]
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
