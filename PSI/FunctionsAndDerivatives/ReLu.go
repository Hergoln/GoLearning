package FunctionsAndDerivatives

import "math"

func ReLu(x []float64) []float64 {
	newOne := make([]float64, len(x))
	for i := range x {
		newOne[i] = math.Max(0, x[i])
	}
	return newOne
}

func ReLuDeriv(x []float64) []float64 {
	newOne := make([]float64, len(x))
	for i := range x {
		if x[i] <= 0 {
			newOne[i] = 0
		} else {
			newOne[i] = 1
		}
	}
	return newOne
}