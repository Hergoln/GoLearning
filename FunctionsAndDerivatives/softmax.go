package FunctionsAndDerivatives

import "math"

func Softmax(x []float64) []float64 {
	result := make([]float64, len(x))
	sum := 0.0
	for i := range x {
		sum += x[i]
	}
	for i := range result {
		result[i] = math.Exp(result[i]) / sum
	}
	return result
}

func softmaxDeriv(delta []float64) []float64 {
	result := make([]float64, len(delta))
	for i := range result {
		result[i] = (delta[i]) / float64(len(delta))
	}
	return result
}