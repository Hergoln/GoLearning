package lab1

import (
	"math/rand"
)

func CreateNeuralNetWithRandomWeights(layersVector []int) HiddenNeuralNetwork {
	network := HiddenNeuralNetwork{}
	network.Layers = make([]LayerNetwork, len(layersVector))
	for i := 0; i < len(network.Layers); i++ {
		net := make([]NeuronN, layersVector[i])
		for netI := 0; netI < len(net); netI++ {
			net[netI] = NeuronN{make([]float64, 1)}
			net[netI].Weights[0] = rand.Float64()
		}
		network.Layers[i].Net = net
	}
	return network
}
