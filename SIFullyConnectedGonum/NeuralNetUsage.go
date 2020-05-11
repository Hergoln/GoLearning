package SIFullyConnectedGonum

import "gonum.org/v1/gonum/mat"

func (network DeepNeuralNet) Predict(input []float64) []float64 {
	denseInput := mat.NewDense(len(input), 1, input)
	for i := range network.Layers {
		denseInput = network.Layers[i].predict(denseInput)
	}
	return denseInput.RawRowView(0)
}

func (network DeepNeuralNet) outputs(input []float64) []*mat.Dense {
	layersCount := len(network.Layers)
	outputs := make([]*mat.Dense, layersCount+1)
	outputs[0] = mat.NewDense(1, len(input), input)
	for i := 0; i < layersCount; i++ {
		outputs[i+1] = network.Layers[i].predict(outputs[i])
	}
	return outputs
}

func (network *DeepNeuralNet) outDelta(netOut, netExp *mat.Dense) *mat.Dense{
	delta := &mat.Dense{}
	delta.Sub(netOut, netExp)
	return delta
}

func hidDelta(nextLayer NeuralLayer, nextLayerDelta *mat.Dense) *mat.Dense{
	result := &mat.Dense{}
	result.Mul(nextLayerDelta, nextLayer.Neurons)
	return result
}

func (network *DeepNeuralNet) Fit(alpha float64, goals []float64, input []float64) float64 {
	layersCount := len(network.Layers)

	outputs := network.outputs(input)

	deltas := make([]*mat.Dense, layersCount)
	deltasDeriv := make([]*mat.Dense, layersCount)

	// deltas
	deltas[layersCount - 1] = network.outDelta(outputs[layersCount], mat.NewDense(1, len(goals), goals))
	for i := layersCount - 1; i > 0; i-- {
		deltas[i-1] = hidDelta(network.Layers[i], deltas[i])
	}

	// error
	errors := 0.0
	x, _ := deltas[layersCount - 1].Dims()
	for i := 0; i < x; i++ {
		rawLastDelta := deltas[layersCount - 1].RawRowView(i)
		for neuron := range rawLastDelta {
			errors += rawLastDelta[neuron]
		}
	}

	// deltas derivatives
	for layer := 0; layer < layersCount; layer ++ {
		deltasDeriv[layer] = mat.DenseCopyOf(deltas[layer])
		x, _ := deltasDeriv[layer].Dims()
		for i := 0; i < x; i++ {
			// here might be problem with dims as well
			deltasDeriv[layer].SetRow(i, network.Layers[i].ActiveFunc(mat.VecDenseCopyOf(deltasDeriv[layer].RowView(i)).RawVector().Data))
		}
	}

	// fitting
	for i := range network.Layers {
		network.Layers[i].fit(alpha, outputs[i], deltas[i])
	}
	return errors
}
