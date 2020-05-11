package SIFullyConnectedGonum

import (
	"gonum.org/v1/gonum/mat"
)

func CreateNetwork(layersIO []int, activeFuncs []ActiveFunc, derivFuncs []ActiveFunc, randStrategy func() float64) *DeepNeuralNet {
	layersCount := len(layersIO) - 1
	network := &DeepNeuralNet{
		Layers: make([]NeuralLayer, layersCount),
	}
	for layerI := 0; layerI < layersCount; layerI++ {
		weights := make([]float64, layersIO[layerI + 1]*layersIO[layerI])
		// layersIO[layerI + 1] number of neurons of current layer and at the same time number of inputs of next layer
		for neuronI := range weights {
			weights[neuronI] = randStrategy()
		}

		newbornLayer := NeuralLayer{
			Neurons:    mat.NewDense(layersIO[layerI + 1], layersIO[layerI], weights),
			ActiveFunc: nil,
			DerivFunc:  nil,
		}
		network.Layers[layerI] = newbornLayer
	}

	for fun := range activeFuncs {
		network.Layers[fun].ActiveFunc = activeFuncs[fun]
		network.Layers[fun].DerivFunc = derivFuncs[fun]
	}
	return network
}

func (network *DeepNeuralNet) AppendLayerWithActiveFunc(outputNeuronsNumber int, activeFunc ActiveFunc, derivFunc ActiveFunc,  randStrategy func() float64) {
	x, _ := network.Layers[len(network.Layers)-1].Neurons.Dims()
	neurons := make([]float64, outputNeuronsNumber * x)

	for i := range neurons {
		neurons[i] = randStrategy()
	}

	newLayer := NeuralLayer{
		Neurons: mat.NewDense(outputNeuronsNumber, x, neurons),
		ActiveFunc: activeFunc,
		DerivFunc: derivFunc,
	}

	network.Layers = append(network.Layers, newLayer)
}