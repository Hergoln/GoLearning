package GoMatrixUtils

func OnesVector(size int) []float64 {
	result := make([]float64, size)
	for i := range result {
		result[i] = 1
	}
	return result
}

func OnesMatrix(width, height int) [][]float64 {
	mat := make([][]float64, height)
	for i := range mat {
		mat[i] = OnesVector(width)
	}
	return mat
}
