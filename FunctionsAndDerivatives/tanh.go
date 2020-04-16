package FunctionsAndDerivatives

import "math"

func Tanh(x []float64) []float64 {
	newOne := make([]float64, len(x))
	for i := range x {
		newOne[i] = math.Tanh(x[i])
	}
	return newOne
}

func TanhDeriv(x []float64) []float64 {
	newOne := make([]float64, len(x))
	for i := range x {
		newOne[i] = 1 - x[i] * x[i]
	}
	return newOne
}