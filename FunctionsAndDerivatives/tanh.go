package FunctionsAndDerivatives

import "math"

func Tanh(x float64) float64 {
	return math.Tanh(x)
}

func TanhDeriv(x float64) float64 {
	return 1 - x * x
}