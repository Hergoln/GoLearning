package FunctionsAndDerivatives

import "math"

func ReLu(x float64) float64 {
	return math.Max(0, x)
}

func ReLuDeriv(x float64) float64 {
	if x <= 0 {
		return 0
	}
	return 1
}