package SI

import "math/rand"

func GibLayers(weights [][][]float64) DeepNeuralNet {
	born := DeepNeuralNet{}
	for i := range weights {
		born.Layers = append(born.Layers, GibNeurons(weights[i]))
	}
	return born
}

// zad 4
func ConstructRandomNetwork(inputsCount, outputsCount int) DeepNeuralNet {
	firstLayer := make([]Neuron, outputsCount)
	for i := range firstLayer {
		firstLayer[i] = Neuron{}
		dendrites := make([]float64, inputsCount)
		for j := range dendrites {
			dendrites[j] = randStrategy1()
		}
		firstLayer[i].Weights = dendrites
	}
	return DeepNeuralNet{[]NeuralLayer{NeuralLayer{Neurons: firstLayer}}}
}

func (network *DeepNeuralNet) appendRandomLayer(outputNeuronsNumber int) {
	neurons := make([]Neuron, outputNeuronsNumber)
	newLayer := NeuralLayer{Neurons: neurons}
	connectionsNumber := len(network.Layers[len(network.Layers)-1].Neurons)

	for i := range neurons {
		dendrites := make([]float64, connectionsNumber)
		for j := range dendrites {
			dendrites[j] = randStrategy1()
		}
		neurons[i] = Neuron{dendrites}
	}

	network.Layers = append(network.Layers, newLayer)
}

func (network *DeepNeuralNet) AppendRandomLayerWithActiveFunc(outputNeuronsNumber int, activeFunc ActiveFunc, derivFunc ActiveFunc) {
	if activeFunc == nil || derivFunc == nil {
		network.appendRandomLayer(outputNeuronsNumber)
		return
	}

	neurons := make([]Neuron, outputNeuronsNumber)
	newLayer := NeuralLayer{
		Neurons: neurons,
		ActiveFunc: activeFunc,
		DerivFunc: derivFunc,
	}
	connectionsNumber := len(network.Layers[len(network.Layers)-1].Neurons)

	for i := range neurons {
		dendrites := make([]float64, connectionsNumber)
		for j := range dendrites {
			dendrites[j] = randStrategy1()
		}
		neurons[i] = Neuron{dendrites}
	}

	network.Layers = append(network.Layers, newLayer)
}

func randStrategy1() float64 {
	return rand.Float64() * 0.2 - 0.1
}