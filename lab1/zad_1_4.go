package lab1

import (
	"fmt"
)

// Hidden network
func Main() {

	// 1.1 neuron
	input := []float64 {1, 1, 1}

	neuron := NeuronN{[]float64{.5}}
	fmt.Printf("zad 1.1\n%f\n", neuron.Scale(input))

	// 1.2 layer
	weights := []float64{.5, .1, 1}
	neuronN := NeuronN{weights}

	fmt.Printf("zad 1.2\n%f\n", neuronN.Scale(input))

	// 1.3 simple network
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

	fmt.Println("1.3 simple network")
	for _, serie := range inputs {
		fmt.Println(network.Calculate(serie))
	}

	// 1.4 hidden network
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

	fmt.Println("zad 1.4 hidden network")
	for _, serie := range inputs {
		fmt.Println(fullNetwork.Calculate(serie))
		fmt.Println()
	}

	// 1.5 fully connected
	hiddenInputs := []float64{8.5, 0.65, 1.2}

	randomNetwork := CreateNeuralNetWithRandomWeights([]int{2, 3, 4})

	fmt.Println(randomNetwork.Calculate(hiddenInputs))
}