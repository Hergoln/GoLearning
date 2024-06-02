package LegacySIFullyConnected

import (
	"math/rand"
)

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
			dendrites[j] = SomewhatOKRandStrategy()
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
			dendrites[j] = SomewhatOKRandStrategy()
		}
		neurons[i] = Neuron{dendrites}
	}

	network.Layers = append(network.Layers, newLayer)
}

func (network *DeepNeuralNet) AppendRandomLayerWithActiveFunc(outputNeuronsNumber int, activeFunc ActiveFunc, derivFunc ActiveFunc) {
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
			dendrites[j] = SomewhatOKRandStrategy()
		}
		neurons[i] = Neuron{dendrites}
	}

	network.Layers = append(network.Layers, newLayer)
}

func SomewhatOKRandStrategy() float64 {
	return rand.Float64() * 2.0 - 0.5
}
