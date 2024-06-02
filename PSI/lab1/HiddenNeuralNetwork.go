package lab1

import "fmt"

type LayerNetwork struct {
	Net []NeuronN
}

func (layer LayerNetwork) Calculate(input []float64) []float64 {
	output := make([]float64, len(layer.Net))
	for i := 0; i < len(layer.Net); i++ {
		output[i] += layer.Net[i].Scale(input)
	}
	return output
}


type HiddenNeuralNetwork struct {
	Layers []LayerNetwork
}

func (neuralNet HiddenNeuralNetwork) Calculate(input []float64) []float64 {
	hiddenInput := input
	for _, layer := range neuralNet.Layers {
		fmt.Println(layer)
		hiddenInput = layer.Calculate(hiddenInput)
	}
	return hiddenInput
}