package SI

type INeuron interface {
	Scale(input []float64) float64
	Study(alpha float64, goal float64, inputs []float64) // for series data
}

type Neuron struct {
	Weights []float64
}

// zad 1, 2
/*
	error = (prediction - goal)^2 = delta^2
	delta = prediction - goal
	weight_delta = input * delta
	weight = weight - weight_delta * alpha = weight - input * (prediction - goal) * alpha
*/
func (neuron Neuron) Scale(input []float64) float64 {
	output := 0.0
	for i := range neuron.Weights {
		output += input[i] * neuron.Weights[i]
	}
	return output
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