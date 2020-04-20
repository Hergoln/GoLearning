package FunctionsAndDerivatives

import "math"

func Softmax(x []float64) []float64 {
	result := make([]float64, len(x))
	sum := 0.0
	for i := range x {
		result[i] = math.Exp(x[i])
		sum += result[i]
	}
	for i := range result {
		result[i] = result[i] / sum
	}
	return result
}

// this softmax derivative implementation is based on assumption that
// what is passed into this function is in reality delta -> (x - expected)
// rather than x itself
func SoftmaxDeriv(pred []float64) []float64 {
	result := make([]float64, len(pred))
	for i := range result {
		result[i] = (pred[i]) / float64(len(result))
	}
	return result
}