package SIFullyConnected

func (layer NeuralLayer) predict(input []float64) []float64 {
	output := make([]float64, len(layer.Neurons))
	if layer.ActiveFunc != nil {
		for i, neuron := range layer.Neurons {
			output[i] = layer.ActiveFunc(neuron.Scale(input))
		}
	} else {
		for i, neuron := range layer.Neurons {
			output[i] = neuron.Scale(input)
		}
	}
	return output
}

func (layer NeuralLayer) scale(input []float64) []float64 {
	output := make([]float64, len(layer.Neurons))
	for i, neuron := range layer.Neurons {
		output[i] = neuron.Scale(input)
	}
	return output
}

func (layer *NeuralLayer) fit(alpha float64, input []float64,  deltas [][]float64) {
	for i := range layer.Neurons {
		layer.Neurons[i].fit(alpha, input, deltas[0][i])
	}
}