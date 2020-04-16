package SIFullyConnected

import (
	"fmt"
	"strings"
)

func (network DeepNeuralNet) Predict(input []float64) []float64 {
	layersIO := input
	for i := 0; i < len(network.Layers) - 1; i++ {
		layersIO = network.Layers[i].predict(input)
	}
	return network.Layers[len(network.Layers) - 1].scale(layersIO)
}

func (network *DeepNeuralNet) Fit(alpha float64, goals []float64, input []float64) float64 {
	// outputs
	outputs := network.outputs(input)

	deltas  := make([][][]float64, len(network.Layers))
	//deltas
	deltas[len(network.Layers) - 1] = network.outDelta(outputs[len(network.Layers)], goals)
	for i := len(network.Layers) - 1; i > 0; i-- {
		deltas[i-1] = hidDelta(network.Layers[i], deltas[i])

		if network.Layers[i - 1].DerivFunc != nil {
			deltasDeriv :=  network.Layers[i - 1].DerivFunc(outputs[i])
			for each := range deltas[i - 1][0] {
				deltas[i - 1][0][each] *= deltasDeriv[each]
			}
		}
	}

	// fitting
	for i := range network.Layers {
		network.Layers[i].fit(alpha, outputs[i], deltas[i])
	}

	errors  := .0
	for neuron := range deltas[len(network.Layers) - 1][0] {
		errors += deltas[len(network.Layers) - 1][0][neuron] * deltas[len(network.Layers) - 1][0][neuron]
	}
	return errors
}

func (network *DeepNeuralNet) outputs(input []float64) [][]float64 {
	layersCount := len(network.Layers)
	outputs := make([][]float64, layersCount + 1)
	outputs[0] = input
	for i := 0; i < layersCount - 1; i++ {
		outputs[i+1] = network.Layers[i].predict(outputs[i])
	}
	outputs[layersCount] = network.Layers[layersCount - 1].scale(outputs[layersCount - 1])
	return outputs
}

func (network *DeepNeuralNet) outDelta(netOut, netExp []float64) [][]float64{
	layersCount := len(network.Layers)
	deltas := make([][]float64, 1)
	deltas[0] = make([]float64, len(netOut))

	for i := range network.Layers[layersCount - 1].Neurons {
		deltas[0][i] = netOut[i] - netExp[i]
	}
	return deltas
}

func hidDelta(nextLayer NeuralLayer, nextLayerDelta [][]float64) [][]float64{
	return layerDeltaMul(nextLayer, nextLayerDelta)
}

func layerDeltaMul(layer NeuralLayer, delta [][]float64) [][]float64 {
	result := make([][]float64, 1)
	result[0] = make([]float64, len(layer.Neurons[0].Weights))

	for i := range result[0] {
		for k := range layer.Neurons {
			result[0][i] += delta[0][k] * layer.Neurons[k].Weights[i]
		}
	}

	return result
}

func (network DeepNeuralNet) Description() string {
	var b strings.Builder
	for iL, layer := range network.Layers  {
		fmt.Fprintf(&b, "{ Layer %d\n", iL)
		for iN, neuron := range layer.Neurons {
			fmt.Fprintf(&b, "\tN%d", iN)
			fmt.Fprintln(&b, neuron)
		}
		fmt.Fprintln(&b, "}")
	}
	return b.String()
}

func (network *DeepNeuralNet) DisplayNet() {
	for iL, layer := range network.Layers  {
		fmt.Printf("{ Layer %d\n", iL)
		for iN, neuron := range layer.Neurons {
			fmt.Printf("\tN%d", iN)
			fmt.Println(neuron)
		}
		fmt.Println("}")
	}
}

func (network DeepNeuralNet) dropout() {
	panic("NotImplemented")
}

func PredictBatch(input [][]float64) {
	panic("NotImplemented")
}
func FitBatch(goals [][]float64, inputs [][]float64) []float64 {
	panic("NotImplemented")
}

func (network DeepNeuralNet) CheckError(expectedOutputs, input []float64) {
	panic("NotImplemented")
}