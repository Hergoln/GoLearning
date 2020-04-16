package FunctionsAndDerivatives

import "math"

func Sigmoid(x []float64) []float64 {
	newOne := make([]float64, len(x))
	for i := range x {
		newOne[i] = 1 / (1 + math.Exp(-x[i]))
	}
	return newOne
}

func SigmoidDeriv(x []float64) []float64 {
	newOne := make([]float64, len(x))
	for i := range x {
		newOne[i] = x[i] * (1 - x[i])
	}
	return newOne
}