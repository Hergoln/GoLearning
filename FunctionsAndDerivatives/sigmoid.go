package FunctionsAndDerivatives

import "math"

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func SigmoidDeriv(x float64) float64 {
	return x * (1 - x)
}