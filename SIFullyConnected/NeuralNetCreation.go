package SIFullyConnected

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

/*
	layersIO -> lengths of layers input vectors, last number is length of network output vector
	activeFuncs -> activation functions for layers, shorter than layersIO by 2
	derivFuncs -> activation functions derivatives for layers, shorter than layersIO by 2
*/
func CreateNetwork(layersIO []int, activeFuncs []ActiveFunc, derivFuncs []ActiveFunc) *DeepNeuralNet {
	layersCount := len(layersIO) - 1
	network := &DeepNeuralNet{
		Layers: make([]NeuralLayer, layersCount),
	}
	for layerI := 0; layerI < layersCount; layerI++ {
		// functions will be added later
		newbornLayer := NeuralLayer{
			Neurons:    make([]Neuron, layersIO[layerI + 1]),
			ActiveFunc: nil,
			DerivFunc:  nil,
		}

		// layersIO[layerI + 1] number of neurons of current layer and at the same time number of inputs of next layer
		for neuronI := range newbornLayer.Neurons {
			dendrites := make([]float64, layersIO[layerI])
			for connectionI := range dendrites {
				dendrites[connectionI] = doctorsOkRandStrategy()
			}
			newbornLayer.Neurons[neuronI] = Neuron{dendrites}
		}

		network.Layers[layerI] = newbornLayer
	}

	for fun := range activeFuncs {
		network.Layers[fun].ActiveFunc = activeFuncs[fun]
		network.Layers[fun].DerivFunc = derivFuncs[fun]
	}


	return network
}

func constructRandomNetwork(inputsCount, outputsCount int) *DeepNeuralNet {
	firstLayer := make([]Neuron, outputsCount)
	for i := range firstLayer {
		firstLayer[i] = Neuron{}
		dendrites := make([]float64, inputsCount)
		for j := range dendrites {
			dendrites[j] = somewhatOKRandStrategy()
		}
		firstLayer[i].Weights = dendrites
	}
	return &DeepNeuralNet{Layers: []NeuralLayer{NeuralLayer{Neurons: firstLayer}}}
}

func (network *DeepNeuralNet) AppendLayerWithActiveFunc(outputNeuronsNumber int, activeFunc ActiveFunc, derivFunc ActiveFunc) {
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
			dendrites[j] = somewhatOKRandStrategy()
		}
		neurons[i] = Neuron{dendrites}
	}

	network.Layers = append(network.Layers, newLayer)
}

func somewhatOKRandStrategy() float64 {
	return rand.Float64() * 2.0 - 0.5
}

func doctorsOkRandStrategy() float64 {
	return rand.Float64() * 0.2 - 0.1
}