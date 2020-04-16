package GoMatrixUtils

func OnesVector(size int) []float64 {
	return make([]float64, size)
}

func OnesMatrix(width, height int) [][]float64 {
	mat := make([][]float64, height)
	for i := range mat {
		mat[i] = make([]float64, width)
		for j := range mat[i] {
			mat[i][j] = 1
		}
	}
	return mat
}
