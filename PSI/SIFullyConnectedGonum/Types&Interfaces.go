package SIFullyConnectedGonum

import "gonum.org/v1/gonum/mat"

type NeuralLayer struct {
	Neurons *mat.Dense
	ActiveFunc ActiveFunc
	DerivFunc ActiveFunc
}

type ActiveFunc func([]float64) []float64 // get neuron([]float64) & return neuron([]float64)

type INeuralNet interface {
	Predict(input []float64) []float64
	Fit(goals []float64, inputs []float64) float64
	PredictBatch(input [][]float64) [][]float64
	FitBatch(goals [][]float64, inputs [][]float64) []float64
}

type DeepNeuralNet struct {
	Layers []NeuralLayer
	DropoutStrategy func() float64
}