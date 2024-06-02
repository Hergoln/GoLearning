package lab1

type NeuronN struct {
	Weights []float64
}

func (neuron NeuronN) Scale(input []float64) float64 {
	output := 0.0
	for i := 0; i < len(neuron.Weights); i++ {
		output += input[i] * neuron.Weights[i]
	}
	return output
}