package main

import (
	"fmt"
)

type Neuron struct {
	Weight float64
}

func (this Neuron) Scale(input float64) float64  {
	return this.Weight * input
}

type NeuronN struct {
	Weights []float64
}

func (this NeuronN) Scale(input []float64) float64 {
	if len(input) != len(this.Weights) {
		return float64(.0)
	}

	output := 0.0
	for i := 0; i < len(this.Weights); i++ {
		output += input[i] * this.Weights[i]
	}
	return output
}

//Simple network
type NeuralNetwork struct {
	Net []NeuronN
}

func (this NeuralNetwork) Calculate(input []float64) []float64 {
	output := make([]float64, len(this.Net))
	for i := 0; i < len(this.Net); i++ {
		output[i] += this.Net[i].Scale(input)
	}
	return output
}


// Hidden network
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
		fmt.Println(layer)
		hiddenInput = layer.Calculate(hiddenInput)
	}
	return hiddenInput
}


func main() {

	input := []float64 {1, 1, 1}

	neuron := Neuron{.5}
	fmt.Println(neuron.Scale(input[0]))

	weights := []float64{.5, .1, 1}
	neuronN := NeuronN{weights}

	fmt.Println(neuronN.Scale(input))

	// simple network
	network := NeuralNetwork{ 
		[]NeuronN {
			{[]float64{.1, .1, -0.3}}, 
			{[]float64{.1, .2, .0}}, 
			{[]float64{.0, 1.3, .1}}}}

	inputs := [][]float64{
		{8.5, 0.65, 1.2}, // seria 1
		{9.5, 0.8, 1.3},  // seria 2
		{9.9, 0.8, 1.3},  // seria 3
		{9.0, 0.9, 1.0}}  // seria 4
	//   in1, in2, in3

	for _, serie := range inputs {
		fmt.Println(network.Calculate(serie))
	}

	// hidden network
	hiddenNetwork := LayerNetwork{ 
		[]NeuronN {
			{[]float64{.1, .2, -0.1}}, 
			{[]float64{-0.1, .1, .9}}, 
			{[]float64{.1, .4, .1}}}}

	outputNetwork := LayerNetwork{ 
		[]NeuronN {
			{[]float64{.3, 1.1, -0.3}}, 
			{[]float64{.1, .2, .0}}, 
			{[]float64{.0, 1.3, .1}}}}

	fullNetwork := HiddenNeuralNetwork{
		[]LayerNetwork{
			hiddenNetwork,
			outputNetwork}}

	fmt.Println("Deep net")
	for _, serie := range inputs {
		fmt.Println(fullNetwork.Calculate(serie))
		fmt.Println()
	}
}