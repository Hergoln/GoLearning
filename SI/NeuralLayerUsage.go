package SI

func (layer NeuralLayer) Predict(input []float64) []float64 {
	output := make([]float64, len(layer.Neurons))
	for i, neuron := range layer.Neurons {
		output[i] = neuron.Scale(input)
	}
	return output
}

func (layer NeuralLayer) PredictActiveFunc(input []float64, activeFunc func(float64) float64) []float64 {
	output := make([]float64, len(layer.Neurons))
	for i, neuron := range layer.Neurons {
		output[i] = activeFunc(neuron.Scale(input))
	}
	return output
}

func (layer *NeuralLayer) Study(alpha float64, goals []float64, inputs []float64) []float64 {
	errors := make([]float64, len(goals))
	for i := range layer.Neurons {
		errors[i] = layer.Neurons[i].Study(alpha, goals[i], inputs)
	}
	return errors
}

// lab 3
func (layer *NeuralLayer) ScaleWithActiveFunc(alpha float64, input []float64,  deltas [][]float64, activeFunc ActiveFunc) {
	for i := range layer.Neurons {
		for j := range layer.Neurons[i].Weights {
			layer.Neurons[i].Weights[j] -= input[j] * deltas[0][i] * alpha
		}
	}
}
