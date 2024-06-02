package LegacySIFullyConnected

func GibWeights(weights []float64) Neuron {
	born := Neuron{}
	for _, weight := range weights {
		born.Weights = append(born.Weights, weight)
	}
	return born
}

func GibNeurons(weights [][]float64) NeuralLayer {
	born := NeuralLayer{}
	for i := range weights {
		born.Neurons = append(born.Neurons, GibWeights(weights[i]))
	}
	return born
}