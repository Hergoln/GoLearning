package SI

import "math/rand"

func GibLayers(weights [][][]float64) DeepNeuralNet {
	born := DeepNeuralNet{}
	for i := range weights {
		born.Layers = append(born.Layers, GibNeurons(weights[i]))
	}
	return born
}

func ConstructRandomNetwork(inputsCount, outputsCount int) DeepNeuralNet {
	firstLayer := make([]Neuron, outputsCount)
	for i := range firstLayer {
		firstLayer[i] = Neuron{}
		dendrites := make([]float64, inputsCount)
		for j := range dendrites {
			dendrites[j] = rand.Float64()
		}
		firstLayer[i].Weights = dendrites
	}
	return DeepNeuralNet{[]NeuralLayer{NeuralLayer{firstLayer}}}
}

func (network *DeepNeuralNet) AppendRandomLayer(outputNeuronsNumber int) {
	neurons := make([]Neuron, outputNeuronsNumber)
	newLayer := NeuralLayer{neurons}
	connectionsNumber := len(network.Layers[len(network.Layers)-1].Neurons)

	for i := range neurons {
		dendrites := make([]float64, connectionsNumber)
		for j := range dendrites {
			dendrites[j] = rand.Float64()
		}
		neurons[i] = Neuron{dendrites}
	}

	network.Layers = append(network.Layers, newLayer)
}

func (network *DeepNeuralNet) AppendRandomLayerWithActiveFunc(outputNeuronsNumber int, activeFunc ActiveFunc) {
	neurons := make([]Neuron, outputNeuronsNumber)
	newLayer := NeuralLayer{neurons}
	connectionsNumber := len(network.Layers[len(network.Layers)-1].Neurons)

	for i := range neurons {
		dendrites := make([]float64, connectionsNumber)
		for j := range dendrites {
			dendrites[j] = rand.Float64()
		}
		neurons[i] = Neuron{dendrites}
	}

	network.Layers = append(network.Layers, newLayer)
}