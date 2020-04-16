package FunctionsAndDerivatives

import "math"

func ReLuFunc(x float64) float64 {
	return math.Max(0, x)
}

func ReLuFuncDeriv(x float64) float64 {
	if x <= 0 {
		return 0
	}
	return 1
}