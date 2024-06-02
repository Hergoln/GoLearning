package SIFullyConnected

func (neuron Neuron) Scale(input []float64) float64 {
	output := 0.0
	for i := range neuron.Weights {
		output += input[i] * neuron.Weights[i]
	}
	return output
}

func (neuron Neuron) fit(alpha float64, input []float64, delta float64) {
	for i := range neuron.Weights {
		neuron.Weights[i] -= alpha * input[i] * delta
	}
}

func (neuron *Neuron) Study(alpha float64, goal float64, inputs []float64) float64 {
	prediction := neuron.Scale(inputs)
	for i := range neuron.Weights {
		neuron.Weights[i] -= inputs[i] * (prediction - goal) * alpha
	}
	return (prediction - goal) * (prediction - goal)
}

func (neuron Neuron) checkError(goal float64, inputs []float64) float64 {
	prediction := neuron.Scale(inputs)
	return (prediction - goal) * (prediction - goal)
}

func (neuron Neuron) SumErrorForSeries(goals []float64, inputs [][]float64) float64 {
	errorsSlice := .0
	for i := range goals {
		errorsSlice += neuron.checkError(goals[i], inputs[i])
	}
	return errorsSlice
}