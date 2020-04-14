package SI

func (network DeepNeuralNet) Predict(input []float64) []float64 {
	layersIO := input
	for _, layer := range network.Layers {
		layersIO = layer.Predict(layersIO)
	}
	return layersIO
}

func (network *DeepNeuralNet) PredictActiveFunc(input []float64, activeFunc ActiveFunc) []float64 {
	layersIO := input
	for i := 0; i < len(network.Layers) - 1; i++ {
		layersIO = network.Layers[i].PredictActiveFunc(layersIO, activeFunc)
	}
	layersIO = network.Layers[len(network.Layers)-1].Predict(layersIO)
	return layersIO
}

func (network *DeepNeuralNet) StudyActiveFunc(alpha float64, goals []float64, input []float64, activeFunc ActiveFunc) float64 {
	outputs := make([][]float64, len(network.Layers) + 1)
	deltas  := make([][][]float64, len(network.Layers))

	// outputs
	outputs[0] = input
	for i := 0; i < len(network.Layers) - 1; i++ {
		outputs[i+1] = network.Layers[i].PredictActiveFunc(outputs[i], activeFunc)
	}
	outputs[len(network.Layers)] = network.Layers[len(network.Layers) - 1].Predict(outputs[len(network.Layers) - 1])

	//deltas
	deltas[len(network.Layers) - 1] = network.outDelta(outputs[len(network.Layers)], goals)
	for i := len(network.Layers) - 1; i > 0; i-- {
		deltas[i-1] = hidDelta(network.Layers[i], deltas[i])
		for each := range deltas[i - 1][0] {
			deltas[i - 1][0][each] *= tempDerivReLu(outputs[i][each])
		}
	}

	for i := range network.Layers {
		network.Layers[i].ScaleWithActiveFunc(alpha, outputs[i], deltas[i], activeFunc)
	}

	errors  := .0
	for neuron := range deltas[len(network.Layers) - 1][0] {
		errors += deltas[len(network.Layers) - 1][0][neuron] * deltas[len(network.Layers) - 1][0][neuron]
	}

	return errors
}


func (network *DeepNeuralNet) outDelta(netOut, netExp []float64) [][]float64{
	layerLen := len(network.Layers)
	deltas := make([][]float64, 1)
	deltas[0] = make([]float64, len(netOut))

	for i := range network.Layers[layerLen - 1].Neurons {
		deltas[0][i] = netOut[i] - netExp[i]
	}
	return deltas
}

func hidDelta(nextLayer NeuralLayer, nextLayerDelta [][]float64) [][]float64{
	return layerDeltaMul(nextLayer, nextLayerDelta)
}

func layerDeltaMul(layer NeuralLayer, delta [][]float64) [][]float64 {
	result := make([][]float64, 1)
	result[0] = make([]float64, len(layer.Neurons))

	for i := range delta[0] {
		for k := range layer.Neurons {
			result[0][i] += delta[0][k] * layer.Neurons[k].Weights[i]
		}
	}

	return result
}

func (network DeepNeuralNet) Dropout() {
	panic("NotImplemented")
}

func (network DeepNeuralNet) PredictBatch(packages interface{}) {
	panic("NotImplemented")
}

func (network DeepNeuralNet) CheckError(expectedOutputs, input []float64) {
	panic("NotImplemented")
}

func tempDerivReLu(f float64) float64 {
	if f <= 0 {
		return 0
	}
	return 1
}