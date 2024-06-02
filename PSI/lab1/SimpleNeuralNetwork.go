package lab1

//Simple network
type NeuralNetwork struct {
	Net []NeuronN
}

func (neuralNet NeuralNetwork) Calculate(input []float64) []float64 {
	output := make([]float64, len(neuralNet.Net))
	for i := 0; i < len(neuralNet.Net); i++ {
		output[i] += neuralNet.Net[i].Scale(input)
	}
	return output
}
