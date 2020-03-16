package main

import (
	"fmt"
	"math/rand"
)

type NeuronN struct {
	Weights []float64
}

func (this NeuronN) Scale(input []float64) float64 {
	output := 0.0
	for i := 0; i < len(this.Weights); i++ {
		output += input[i] * this.Weights[i]
	}
	return output
}

type LayerNetwork struct {
	Net []NeuronN
}

func (this LayerNetwork) Calculate(input []float64) []float64 {
	output := make([]float64, len(this.Net))
	for i := 0; i < len(this.Net); i++ {
		output[i] += this.Net[i].Scale(input)
	}
	return output
}

type HiddenNeuralNetwork struct {
	Layers []LayerNetwork
}

func (this HiddenNeuralNetwork) Calculate(input []float64) []float64 {
	hiddenInput := input
	for _, layer := range this.Layers {
		hiddenInput = layer.Calculate(hiddenInput)
	}
	return hiddenInput
}

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

func main() {
	input := []float64{8.5, 0.65, 1.2}

	network := CreateNeuralNetWithRandomWeights([]int{2, 3, 4})

	fmt.Println(network.Calculate(input))
}
