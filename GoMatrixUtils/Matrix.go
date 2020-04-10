package GoMatrixUtils


func MatrixMatrixMul(first, other [][]float64) [][]float64 {
	result := make([][]float64, len(first))
	for i := range result {
		result[i] = make([]float64, len(other[0]))
	}

	for i := range first {
		for j := range other[0] {
			for k := range other {
				result[i][j] += first[i][k] * other[k][j]
			}
		}
	}

	return result
}

func MatrixScalarMul(matrix [][]float64, scalar float64) [][]float64{
	result := make([][]float64, len(matrix))

	for i := range result {
		result[i] = make([]float64, len(matrix[i]))
		for j := range result[0] {
			result[i][j] = matrix[i][j] * scalar
		}
	}

	return result
}

func ScaleMatrix(matrix [][]float64, scalingFact func(float64) float64) [][]float64 {
	for i := range matrix {
		for j := range matrix[0] {
			matrix[i][j] = scalingFact(matrix[i][j])
		}
	}
	return matrix
}